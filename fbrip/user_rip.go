package fbrip

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type UserRip struct {
	Parameters map[string]string
	Cookies    []*http.Cookie
	Info       InfoStruct
}

func CreateUser(email string, pass string) UserRip {
	parameters := map[string]string{
		"email": email,
		"pass": pass,
		"lsd": "", 
		"jazoest": "",
		"m_ts": "",
		"li": "",
		"try_number": "",
		"unrecognized_tries": "",
		"login": "",
	}
	userRip := UserRip{Parameters: parameters}
	return userRip
}

func (u *UserRip) Sense() {
	// Making GET request for https://mbasic.facebook.com/
	URL_struct, _ := url.Parse("https://mbasic.facebook.com/")
	response := u.GET(URL_struct)

	//Getting cookies & saving them to user
	u.MergeCookies(response.Cookies())

	//Parsing html returning an *html.Node. Searching params and adding them to user.
	defer response.Body.Close()
	searchParamsForUser(response.Body, u)
}

func (u *UserRip) Rip() bool {
	URL_struct, _ := url.Parse("https://mbasic.facebook.com/login/device-based/regular/login/")
	//Starting Login Process
	loginRequest := u.ripPhase1(URL_struct)
	if u.ValidCookies() {
		u.ripPhase2(loginRequest)
		u.ripPhase3()
		return true
	}
	fmt.Printf("** Error while ripping to user: %s\n", u.Parameters["email"])
	return false
}

func (u *UserRip) Do(config *ActionConfig) {
	//Getting Basic Info
	if config.GetBasicInfo {
		u.GetBasicInfo()
	}
	//Make Reaction to acertain post
	if config.React.Checks() {
		u.MakeReactions(config.React.Urls, config.React.Ids)
	}
	//Scrap Urls
	if len(config.Scrap.Urls) > 0 {
		u.Scrap(config.Scrap.Urls, config.Scrap.FolderPath)
	}
	if config.Post.Url != nil && config.Post.Content != "" {
		//TO DEVELOP
		fmt.Println("`fbrip` for the moment does not contain logic for posting :( ")
		fmt.Println("comming soon...")
	}
	if config.Comment.Url != nil && config.Comment.Content != "" {
		//TO DEVELOP
		fmt.Println("`fbrip` for the moment does not contain logic for comment :( ")
		fmt.Println("comming soon...")
	}
}

func (u *UserRip) ripPhase1(URL_struct *url.URL) *http.Request {
	//Get user's parameters as url.Values type
	parameters := u.GetParametersAsUrlValues()
	//Making request to URL with respective parameters & setting its headers
	request, _ := http.NewRequest("POST", URL_struct.String(), strings.NewReader(parameters.Encode()))
	setHeaders(request, "application/x-www-form-urlencoded;", len(parameters.Encode()))

	//Injecting cookies and getting Jar to be passed to client
	jar := u.GetAndInjectCookies(request)

	// Making an HTTP Client and a New Request  &  Saving cookies from  response with [StatusCode = 302]
	var loginRequest *http.Request
	client := &http.Client{
		CheckRedirect: func(request *http.Request, via []*http.Request) error {
			loginRequest = request
			return http.ErrUseLastResponse
		},
		Jar: jar,
	}
	//Doing POST request & getting a response with [StatusCode = 302]
	response, _ := client.Do(request)
	response.Body.Close()
	//Merging response cookies to user
	u.MergeCookies(response.Cookies())

	return loginRequest
}

func (u *UserRip) ripPhase2(loginRequest *http.Request) *http.Response {
	//Injecting cookies
	jar := u.GetAndInjectCookies(loginRequest)
	//Making http client
	client := &http.Client{Jar: jar}
	//Doing POST request & getting a response with [StatusCode = 200]
	response, _ := client.Do(loginRequest)
	response.Body.Close()

	return response
}

func (u *UserRip) ripPhase3() *http.Response {
	//URL To submit the cancelation of `sign in with a touch`
	URL_struct, _ := url.Parse("https://mbasic.facebook.com/login/save-device/cancel/?flow=interstitial_nux&nux_source=regular_login")
	//Making GET Request and Closing Body Response
	response := u.GET(URL_struct)
	response.Body.Close()

	return response
}

func (u *UserRip) GET(URL_struct *url.URL) *http.Response {
	//Making new http GET request
	request, _ := http.NewRequest("GET", URL_struct.String(), nil)
	setHeaders(request, "", -1)
	//Injecting cookies
	jar := u.GetAndInjectCookies(request)
	//Making http client
	client := &http.Client{Jar: jar}
	//Doing GET request
	response, _ := client.Do(request)
	return response
}

func (u *UserRip) GetParametersAsUrlValues() url.Values {
	// Setting user's parameters
	parameters := url.Values{}
	for param := range u.Parameters {
		parameters.Set(param, u.Parameters[param])
	}
	return parameters
}

func (u *UserRip) GetParameterKeys() []string {
	keys := make([]string, 0, len(u.Parameters))
	for k := range u.Parameters {
		keys = append(keys, k)
	}
	return keys
}

func (u *UserRip) GetAndInjectCookies(request *http.Request) *cookiejar.Jar {
	//Adding cookies to Jar
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(request.URL, u.Cookies)

	//Adding cookies to Request
	for _, cookie := range u.Cookies {
		request.AddCookie(cookie)
	}
	return jar
}

func (u *UserRip) MergeCookies(c1 []*http.Cookie) {
	for _, cookie := range c1 {
		if !includesCookie(u.Cookies, cookie) {
			u.Cookies = append(u.Cookies, cookie)
		}
	}
}

//Validates if the user has the necessary cookies to login.
//Coockies = "datr", "sb", "c_user", "xs", "fr"
func (u *UserRip) ValidCookies() bool {
	counter := 0
	for _, cookie := range u.Cookies {
		switch cookie.Name {
		case "datr", "sb", "c_user", "xs", "fr":
			counter += 1
		}
	}
	return counter == 5
}
