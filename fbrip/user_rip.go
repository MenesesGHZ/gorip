package fbrip

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
)

type UserRip struct {
	Email      string
	Password   string
	Parameters map[string]string
	Client     *http.Client
	Info       UserInfo
}

func NewUserRip(email string, password string) *UserRip {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: cookieJar}
	parameters := map[string]string{
		"email":              email,
		"pass":               password,
		"lsd":                "",
		"jazoest":            "",
		"m_ts":               "",
		"li":                 "",
		"try_number":         "",
		"unrecognized_tries": "",
		"login":              "",
	}
	userRip := &UserRip{
		Parameters: parameters,
		Email:      email,
		Password:   password,
		Client:     client,
	}
	return userRip
}

// Getting first set of cookies and parameters needed for make a login request
// Cookies Gathered:
//   - datr				 (e.g. 'vhmkYoqy7RdEbjo_7-CfCB1A')
//
// Parameters Gathered:
//   - jazoest 			 (e.g. 2879)
//   - li 	  			 (e.g 'vhmkYn8H32beqTnQp3ZeUcq3')
//   - login				 (e.g 'Log in')
//   - lsd 	  			 (e.g 'AVqG3uZN6UE')
//   - m_ts    			 (e.g 1654921662)
//   - try_number 		 (e.g 0)
//   - unrecognized_tries (e.g 0)
func (u *UserRip) Sense() error {
	baseUrl, _ := url.Parse("https://mbasic.facebook.com/")
	request, _ := http.NewRequest("GET", baseUrl.String(), nil)
	SetHeaders(request, "", -1)
	response, err := u.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	SearchParamsForUser(response.Body, u)
	return nil
}

// Login workflow; Setting policy for handling redirects by returning `http.ErrUseLastResponse`
// to avoid making next request automatically since is no needed for login.
// Cookies Gathered (minus sense() cookies):
//   - sb	 (e.g. 'mT-kYiYOVgO1REEuVoN3QIkt')
//   - c_user (e.g. 100008137277101)
//   - xs	 (e.g. '3%3AgAfz50LpTd4C6A%3A2%3A1654931354%3A-1%3A2298')
//   - fr	 (e.g. '0z1tHKUHfVz6RQcyW.AWXxkBqxuktzL1QQzfdJ4Z_ZeQ4.BipD-a.pb.AAA.0.0.BipD-a.AWUAUnp-IxI')
func (u *UserRip) Rip() bool {
	err := u.Sense()
	if err != nil {
		return false
	}
	loginUrl, _ := url.Parse("https://mbasic.facebook.com/login/device-based/regular/login/")
	parameters := u.GetParametersAsUrlValues()
	request, _ := http.NewRequest("POST", loginUrl.String(), strings.NewReader(parameters.Encode()))
	SetHeaders(request, "application/x-www-form-urlencoded;", len(parameters.Encode()))
	u.Client.CheckRedirect = func(request *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	response, err := u.Client.Do(request)
	response.Body.Close()
	if err != nil || !u.ValidCookies() {
		return false
	}
	u.Client.CheckRedirect = nil
	return true
}

func (u *UserRip) GetRequest(requestUrl *url.URL) *http.Response {
	request, _ := http.NewRequest("GET", requestUrl.String(), nil)
	SetHeaders(request, "", -1)
	response, _ := u.Client.Do(request)
	return response
}

func (u *UserRip) GetParametersAsUrlValues() url.Values {
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

// Validates if the user has the necessary cookies to login.
// Coockies = "datr", "sb", "c_user", "xs", "fr"
func (u *UserRip) ValidCookies() bool {
	counter := 0
	for _, cookie := range u.Client.Jar.Cookies(FacebookUrl) {
		switch cookie.Name {
		case "datr", "sb", "c_user", "xs", "fr":
			counter += 1
		}
	}
	return counter == 5
}

func SetHeaders(request *http.Request, contentType string, paramsLength int) {
	//Setting default headers
	request.Header.Set("Host", request.URL.Host)
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	request.Header.Set("Accept-Language", "en-US,en;q=1.0")
	request.Header.Set("Connection", "close")
	request.Header.Set("Upgrade-Insecure-Requests", "1")

	//Setting parameters if POST request
	if request.Method == "POST" {
		request.Header.Set("Content-Type", contentType)
		request.Header.Set("Content-Length", strconv.Itoa(paramsLength))
		request.Header.Set("Origin", request.URL.String())
	}
}
