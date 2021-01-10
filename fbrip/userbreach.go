package fbrip

import(
	"fmt"
	"strconv"
	"golang.org/x/net/html"
	"net/url"
	"net/http"
	"net/http/cookiejar"
)


type UserBreach struct{
	name string
	gender string
	birthdate string
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

func (u *UserBreach) Do(content ActionContent, config ActionConfig) bool{
	success := false
	if(config.GetBasicInfo){
		success = u.getBasicInfo()
	}
	if(config.MakeReaction){
	//	URL_struct,_ := url.Parse(content["url"])
	//	success = u.makeReaction(URL_struct,content["reaction"])
	}
	if(config.MakePost){
		//TO DEVELOP
		fmt.Println("`fbreach` for the moment does not contain logic for posting :( ")
		fmt.Println("comming soon...")
	}
	return success
}

func (u *UserBreach) Sense()  {
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


func (u *UserBreach) Rip(){
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


func setHeaders(request *http.Request, contentType string, paramsLength int){
	//Setting default headers
	request.Header.Set("Host",request.URL.Host)
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	request.Header.Set("Accept-Language", "en-US,en;q=1.0")
	request.Header.Set("Connection", "close")
	request.Header.Set("Upgrade-Insecure-Requests", "1")

	//Setting parameters if POST request
	if request.Method == "POST"{
		request.Header.Set("Content-Type",contentType)
		request.Header.Set("Content-Length", strconv.Itoa(paramsLength))
		request.Header.Set("Origin", request.URL.String())
	}
}



func(u *UserBreach) GET(URL_struct *url.URL) *http.Response{
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

// Search for input parameters. * It must be improved
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


