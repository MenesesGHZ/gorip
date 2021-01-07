package fbreach

import(
	"fmt"
	"strconv"
	"strings"
	"golang.org/x/net/html"
    	"net/url"
	"net/http"
	"net/http/cookiejar"
)


type UserBreach struct{
	parameters map[string]string
	cookies []*http.Cookie
}

func CreateUser(email string, pass string) UserBreach{
	parameters := make(map[string]string)
	parameters["email"] = email
	parameters["pass"] = pass
	userBreach := UserBreach{parameters:parameters}
	return userBreach
}


func (u *UserBreach) Sense(URL_struct *url.URL)  {
	// Making GET request for https://mbasic.facebook.com/
	response, err := http.Get(URL_struct.String())
	if err!=nil{
		fmt.Println(err)
	}
	
	//Getting cookies & saving them to user
	var cookies []*http.Cookie
	for _,cookie := range response.Cookies() {
		if includes(CookieNames,cookie.Name){
			cookies = append(cookies,cookie)
		}
	}
	u.cookies = cookies
	
	//Parsing html returning an *html.Node. Searching params and adding them to user.
	doc,_ := html.Parse(response.Body)
	searchParameters(doc,u)
}


func (u *UserBreach) Rip(URL_struct *url.URL){
		// FACEBOOK LOGIN //
	
	//Adding cookies to URL
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(URL_struct, u.cookies)

	// Setting parameters and econding them
	parameters := url.Values{}
	for _,param := range ParameterNames{
		parameters.Set(param, u.parameters[param])
	}
	fmt.Println("Encoded Parameters:",parameters.Encode())
	
	// Making an HTTP Client and a New Request 
	client :=  &http.Client{
			CheckRedirect: func(request *http.Request, via []*http.Request) error {
				fmt.Println("REDIRECT:",request,via)
				return http.ErrUseLastResponse
			},
			Jar:jar,
		}
	request,_:= http.NewRequest("POST",URL_struct.String(),strings.NewReader(parameters.Encode()))
	
	//Adding cookies
	for _,cookie := range u.cookies{
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
}


// Extra Variables
var ParameterNames = []string{
	"lsd",
	"jazoest",
	"m_ts",
	"li",
	"try_number",
	"unrecognized_tries",
	"email",
	"pass",
	"login",
}

var CookieNames = []string{
	"datr",
	"sb",
	"c_user",
	"xs",
	"fr",
}


// Extra functions
func searchParameters(node *html.Node, u *UserBreach){
	// Declaration of functions
	var engine func(*html.Node)
	
	// Defining functions
	engine = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "input" {
			for _,attr := range n.Attr{
				if includes(ParameterNames,attr.Val){
					for _,attr2 := range n.Attr{
						if attr2.Key == "value"{
							u.parameters[attr.Val] = attr2.Val
							break
						}
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			engine(c)
		}
	}
	
	// Running engine
	engine(node)
}



func includes(slice []string,v string) bool{
	for _,value := range slice{
		if value == v{
			return true
		}
	}
	return false
}


