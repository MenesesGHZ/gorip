package fbrip

import(
	"fmt"
	"strings"
	"net/url"
	"net/http"
	"net/http/cookiejar"
	"golang.org/x/net/html"
)

type UserRip struct{
	Parameters map[string]string
	Cookies []*http.Cookie
	Info info
}

func CreateUser(email string, pass string) UserRip{
	parameters := make(map[string]string)
	parameters["email"] = email
	parameters["pass"] = pass
	userRip := UserRip{Parameters:parameters}
	return userRip
}

func (u *UserRip) Sense()  {
	// Making GET request for https://mbasic.facebook.com/
	URL_struct,_ := url.Parse("https://mbasic.facebook.com/")
	response := u.GET(URL_struct)
	
	//Getting cookies & saving them to user
	u.MergeCookies(response.Cookies())
	
	//Parsing html returning an *html.Node. Searching params and adding them to user.
	defer response.Body.Close()
	doc,_ := html.Parse(response.Body)
	searchParameters(doc,u)
	
	fmt.Println("Sense Completed.\n")
}

func (u *UserRip) Rip() {
	URL_struct,_ := url.Parse("https://mbasic.facebook.com/login/device-based/regular/login/")
	
	loginRequest,status := u.ripPhase1(URL_struct)
	if status == 302{
		fmt.Println("*Rip [1/3] Completed.")
	}
	_,status = u.ripPhase2(loginRequest)
	if status == 200{
		fmt.Println("*Rip [2/3] Completed.")
	}
	_,status = u.ripPhase3()
	if status == 200{
		fmt.Println("*Rip [3/3] Completed.")
	}
}

func (u *UserRip) Do(config *ActionConfig) bool{
	success := false
	if(config.GetBasicInfo){
		success = u.getBasicInfo()
	}
	if(config.React.Url!=nil && config.React.Id != ""){
		success = u.makeReaction(config.React.Url, config.React.Id)
	}
	if(config.Post.Url!=nil && config.Post.Content != ""){
		//TO DEVELOP
		fmt.Println("`fbrip` for the moment does not contain logic for posting :( ")
		fmt.Println("comming soon...")
	}
	if(config.Comment.Url!=nil && config.Comment.Content != ""){
		//TO DEVELOP
		fmt.Println("`fbrip` for the moment does not contain logic for comment :( ")
		fmt.Println("comming soon...")
	}
	return success
}

func (u *UserRip) ripPhase1(URL_struct *url.URL) (*http.Request,int){
	//Get user's parameters as url.Values type
	parameters := u.GetParameters()

	//Making request to URL with respective parameters & setting its headers
	request,_:= http.NewRequest("POST",URL_struct.String(),strings.NewReader(parameters.Encode()))
	setHeaders(request, "application/x-www-form-urlencoded;", len(parameters.Encode()))

	//Injecting cookies and getting Jar to be passed to client
	jar := u.GetAndInjectCookies(request)

	// Making an HTTP Client and a New Request  &  Saving cookies from  response with [StatusCode = 302]
	var loginRequest *http.Request
	client :=  &http.Client{
		CheckRedirect: func(request *http.Request, via []*http.Request) error {
			loginRequest = request
			return http.ErrUseLastResponse
		},
		Jar:jar,
	}
	//Doing POST request & getting a response with [StatusCode = 302]
	response,_ := client.Do(request)
	response.Body.Close()

	//Merging response cookies to user
	u.MergeCookies(response.Cookies())

	return loginRequest,response.StatusCode
}

func (u *UserRip) ripPhase2(loginRequest *http.Request) (*http.Response,int){
	//Injecting cookies
	jar := u.GetAndInjectCookies(loginRequest)

	//Making http client
	client :=  &http.Client{Jar:jar}

	//Doing POST request & getting a response with [StatusCode = 200]
	response,_ := client.Do(loginRequest)
	response.Body.Close()

	return response,response.StatusCode
}

func (u *UserRip) ripPhase3() (*http.Response,int){
	//URL To submit the cancelation of `sign in with a touch`
	URL_struct,_ := url.Parse("https://mbasic.facebook.com/login/save-device/cancel/?flow=interstitial_nux&nux_source=regular_login")
	//Making GET Request and Closing Body Response
	response := u.GET(URL_struct)
	response.Body.Close()

	return response,response.StatusCode
}

func(u *UserRip) GET(URL_struct *url.URL) *http.Response{
	//Making new http GET request
	request,_:= http.NewRequest("GET",URL_struct.String(),nil)
	setHeaders(request, "", -1)
	//Injecting cookies
	jar := u.GetAndInjectCookies(request)
	//Making http client
	client :=  &http.Client{Jar:jar}
	//Doing GET request
	response,_ := client.Do(request)
	return response
}

func(u *UserRip) GetParameters() url.Values{
	// Setting user's parameters 
	parameters := url.Values{}
	for _,param := range ParameterNames{
		parameters.Set(param, u.Parameters[param])
	}
	return parameters
}

func (u *UserRip) GetAndInjectCookies(request *http.Request) *cookiejar.Jar{
	//Adding cookies to Jar
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(request.URL, u.Cookies)

	//Adding cookies to Request
	for _,cookie := range u.Cookies{
		request.AddCookie(cookie)
	}
	return jar
}

func (u *UserRip) MergeCookies(c1 []*http.Cookie){
	for _,cookie := range c1{
		if !includesCookie(u.Cookies,cookie){
			u.Cookies = append(u.Cookies,cookie)
		}
	}
}
