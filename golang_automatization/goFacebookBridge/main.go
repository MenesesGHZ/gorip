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

type NetData struct{
	hostname string
	protocol string
	port int
}

func sense(URL_string string, bridgeUser *BridgeUser)  {
	fmt.Println("URL:",URL_string)
	response, err := http.Get(URL_string)
	if err!=nil{
		fmt.Println(err)
	}
	var cookies []*http.Cookie
	for _,cookie := range response.Cookies(){
		if includes(cookieFilter,cookie.Name) {
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
	fmt.Println(jar)
	
	// Setting parameters and econding them
	parameters := url.Values{}
	parameters.Set("email", bridgeUser.email)
	parameters.Set("pass", bridgeUser.pass)
	parameters.Set("lsd", "AVpJfDk-quE")
	fmt.Println("Encoded Parameters:",parameters.Encode())
	fmt.Println("URL:",URL_struct.String())

	// Making an HTTP Client and a New Request 
	client :=  &http.Client{
			Jar:jar,
			CheckRedirect: func(request *http.Request, via []*http.Request) error {
				fmt.Println("REDIRECT:",request,via)
				return http.ErrUseLastResponse
			}}

	request,_ := http.NewRequest(http.MethodPost,URL_struct.String(),strings.NewReader(parameters.Encode()))
	//Adding Headers to Request
	request.Header.Add("Host",URL_struct.Host)
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Accept-Language", "en-US,en;q=0.5")
	request.Header.Add("Accept-Encoding", "gzip, deflate")
	request.Header.Add( "Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add( "Content-Length", strconv.Itoa(len(parameters.Encode())))
	request.Header.Add( "Origin", URL_struct.String())
	request.Header.Add( "Connection", "close")
	request.Header.Add( "Upgrade-Insecure-Requests", "1")

	//Making POST request
	response,err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("StatusCode:", response.StatusCode)
	fmt.Println(response.Request.URL)
	fmt.Println(response.Cookies())

}

