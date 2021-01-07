package main

import(
	"fmt"
	"strconv"
	"strings"
    	"net/url"
	"net/http"
	"net/http/cookiejar"
)

func includes(slice []string,v string) bool{
	for _,value := range slice{
		if value == v{
			return true
		}
	}
	return false
}

var cookieFilter = []string{"datr","sb","c_user","xs","fr"}

type BridgeUser struct{
	email string
	pass string
	cookies []*http.Cookie
}


func sense(URL_string string, bridgeUser *BridgeUser)  {
	response, err := http.Get(URL_string)
	if err!=nil{
		fmt.Println(err)
	}
	var cookies []*http.Cookie
	for _,cookie := range response.Cookies() {
		if includes(cookieFilter,cookie.Name){
			cookies = append(cookies,cookie)
		}
	}
	bridgeUser.cookies = cookies
}

func makeUrl(URL_string string, path string) *url.URL{
	URL_struct,_ := url.Parse(URL_string+path)
	return URL_struct
}


func main(){

	bridgeUser := BridgeUser{
		email:"ghzant.y@gmail.com",
		pass:"password(password)",
	}
	
	// GET for getting part of Cookies
	URL_struct := makeUrl("https://mbasic.facebook.com/","")
	sense(URL_struct.String(),&bridgeUser)
	
		// FACEBOOK LOGIN //
	//Defining URL
	URL_struct = makeUrl(URL_struct.String(),"login/device-based/regular/login/")
	fmt.Println("HOST:",URL_struct.Host)
	
	//Adding cookies to URL
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(URL_struct, bridgeUser.cookies)

	// Setting parameters and econding them
	parameters := url.Values{}
	parameters.Set("email", bridgeUser.email)
	parameters.Set("pass", bridgeUser.pass)
	parameters.Set("lsd", "AVpJfDk-quE")
	fmt.Println("Encoded Parameters:",parameters.Encode())
	
	// Making an HTTP Client and a New Request 
	client :=  &http.Client{
			CheckRedirect: func(request *http.Request, via []*http.Request) error {
				fmt.Println("REDIRECT:",request,via)
				return http.ErrUseLastResponse
			},
			Jar:jar,
		}
	request,_:= http.NewRequest(http.MethodPost,URL_struct.String(),strings.NewReader(parameters.Encode()))
	
	//Adding cookies
	for _,cookie := range bridgeUser.cookies{
		request.AddCookie(cookie)
	}
	fmt.Println("Cookies:",request.Cookies())

	//Adding Headers to Request
	request.Header.Set("Host",URL_struct.Host)
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0")
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Accept-Language", "en-US,en;q=0.5")
	request.Header.Set("Accept-Encoding", "gzip, deflate")
	request.Header.Set( "Content-Type", "application/x-www-form-urlencoded; param=value")
	request.Header.Set( "Content-Length", strconv.Itoa(len(parameters.Encode())))
	request.Header.Set( "Origin", URL_struct.String())
	request.Header.Set( "Connection", "close")
	request.Header.Set( "Upgrade-Insecure-Requests", "1")
	
	//Making POST request
	fmt.Println("POST to:",URL_struct.String())
	response,err := client.Do(request)

	if err != nil {
		fmt.Println("ERROR",err)
	}
	fmt.Println("CookieJar:",jar.Cookies(request.URL))
	fmt.Println("StatusCode:", response.StatusCode)
	fmt.Println("Response Cookies",response.Cookies())
	defer response.Body.Close()
}

