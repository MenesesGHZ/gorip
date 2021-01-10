package fbrip


import(
	"strings"
	"net/url"
	"net/http"
)

func (u *UserBreach) ripPhase1(URL_struct *url.URL) (*http.Request,int){
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

func (u *UserBreach) ripPhase2(loginRequest *http.Request) (*http.Response,int){
	//Injecting cookies
	jar := u.GetAndInjectCookies(loginRequest)

	//Making http client
	client :=  &http.Client{Jar:jar}

	//Doing POST request & getting a response with [StatusCode = 200]
	response,_ := client.Do(loginRequest)
	response.Body.Close()

	return response,response.StatusCode
}

func (u *UserBreach) ripPhase3() (*http.Response,int){

	//URL To submit the cancelation of `sign in with a touch`
	URL_struct,_ := url.Parse("https://mbasic.facebook.com/login/save-device/cancel/?flow=interstitial_nux&nux_source=regular_login")
	//Making GET Request and Closing Body Response
	response := u.GET(URL_struct)
	response.Body.Close()

	return response,response.StatusCode
}
