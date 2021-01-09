package fbreach

import(
	"fmt"
	"strconv"
	"strings"
	"golang.org/x/net/html"
    	"net/url"
	"net/http"
	"net/http/cookiejar"
	"io/ioutil"
)


type UserBreach struct{
	Parameters map[string]string
	Cookies []*http.Cookie
}

func CreateUser(email string, pass string) UserBreach{
	parameters := make(map[string]string)
	parameters["email"] = email
	parameters["pass"] = pass
	userBreach := UserBreach{Parameters:parameters}
	return userBreach
}


func (u *UserBreach) Sense()  {
	// Making GET request for https://mbasic.facebook.com/
	URL_struct,_ := url.Parse("https://mbasic.facebook.com/")
	response, err := http.Get(URL_struct.String())
	if err!=nil{
		fmt.Println(err)
	}
	
	//Getting cookies & saving them to user
	u.MergeCookies(response.Cookies())
	
	//Parsing html returning an *html.Node. Searching params and adding them to user.
	defer response.Body.Close()
	doc,_ := html.Parse(response.Body)
	searchParameters(doc,u)
}


func (u *UserBreach) Rip(){
	URL_struct,_ := url.Parse("https://mbasic.facebook.com/login/device-based/regular/login/")
	
	//Ripping 	
	loginRequest := u.ripPhase1(URL_struct)
	response := u.ripPhase2(loginRequest)
	
	//printing stuff
	fmt.Println("\n\n------------------------------------\n\n")
	fmt.Println("HEADERS:\n",response.Header)
	fmt.Println("\n\n------------------------------------\n\n")
	fmt.Println("NATURAL RESPONSE:\n",response)
}


func setHeaders(request *http.Request, contentType string, paramsLength int){
	//Setting default headers
	request.Header.Set("Host",request.URL.Host)
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	request.Header.Set("Accept-Language", "en-US,en;q=0.5")
	request.Header.Set("Accept-Encoding", "gzip, deflate")
	request.Header.Set("Connection", "close")
	request.Header.Set("Upgrade-Insecure-Requests", "1")

	//Setting parameters if POST request
	if request.Method == "POST"{
		request.Header.Set("Content-Type",contentType)
		request.Header.Set("Content-Length", strconv.Itoa(paramsLength))
		request.Header.Set("Origin", request.URL.String())
	}
}

func (u *UserBreach) ripPhase1(URL_struct *url.URL) *http.Request{
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
	defer response.Body.Close()
	
	//Merging response cookies to user
	u.MergeCookies(response.Cookies())
	
	//Printing status code
	fmt.Printf("[ StatusCode = %d ]\n",response.StatusCode)
	
	return loginRequest
}

func (u *UserBreach) ripPhase2(loginRequest *http.Request) *http.Response{
	//Injecting cookies
	jar := u.GetAndInjectCookies(loginRequest)
	
	//Making http client
	client :=  &http.Client{Jar:jar}

	//Doing POST request & getting a response with [StatusCode = 200]
	response,_ := client.Do(loginRequest)
	defer response.Body.Close()
	
	//Printing status code [ StatusCode = 200 ]
	fmt.Printf("[ StatusCode = %d ]\n",response.StatusCode)

	body,err := ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Println("BODY ERROR:",err)
	}
	fmt.Println("BODY:",string(body))


	return response
}

func(u *UserBreach) GetParameters() url.Values{
	// Setting user's parameters 
	parameters := url.Values{}
	for _,param := range ParameterNames{
		parameters.Set(param, u.Parameters[param])
	}
	return parameters
}

func (u *UserBreach) GetAndInjectCookies(request *http.Request) *cookiejar.Jar{
	//Adding cookies to Jar
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(request.URL, u.Cookies)
	
	//Adding cookies to Request
	for _,cookie := range u.Cookies{
		request.AddCookie(cookie)
	}
	return jar
}

func (u *UserBreach) MergeCookies(c1 []*http.Cookie){
	for _,cookie := range c1{
		if !includesCookie(u.Cookies,cookie){
			u.Cookies = append(u.Cookies,cookie)
		}
	}
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
							u.Parameters[attr.Val] = attr2.Val
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

func includesCookie(cookies []*http.Cookie, cookie *http.Cookie) bool{
	for _,c := range cookies{
		if c.Name == cookie.Name{
			return true
		}
	}
	return false
}


