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




var CookieNames = []string{
	"datr",
	"sb",
	"c_user",
	"xs",
	"fr",
}

func CreateUser(email string, pass string) UserBreach{
	parameters := make(map[string]string)
	parameters["email"] = email
	parameters["pass"] = pass
	userBreach := UserBreach{parameters:parameters}
	return userBreach
}


func (u *UserBreach) Sense(URL_struct *url.URL)  {
	response, err := http.Get(URL_struct.String())
	if err!=nil{
		fmt.Println(err)
	}
	
	//Getting cookies
	var cookies []*http.Cookie
	for _,cookie := range response.Cookies() {
		if includes(CookieNames,cookie.Name){
			cookies = append(cookies,cookie)
		}
	}
	u.cookies = cookies

	//Getting parameters
//	tkz := html.NewTokenizer(response.Body)
//	for{
//		tkd:=tkz.Next()
//		if tkd == html.ErrorToken {
//			return
//		}
//		fmt.Println("Token:",tkz.Token())
//	}
	doc,err := html.Parse(response.Body)

}


func (u *UserBreach) rip(URL_struct *url.URL){
		// FACEBOOK LOGIN //

	//Adding cookies to URL
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(URL_struct, u.cookies)

	// Setting parameters and econding them
	parameters := url.Values{}
	parameters.Set("email", u.parameters["email"])
	parameters.Set("pass", u.parameters["pass"])
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
	defer response.Body.Close()
}




// extra functions
func searchParameters(node *html.Node){
	var engine func(*html.Node)
	engine = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "input" {
			for _,attr := range 
			fmt.Println("ATTR:",n.Attr)
			fmt.Println("KEY:",n.Attr[1].Key)
			fmt.Println("VALUE:",n.Attr[1].Val)
			fmt.Println("NAMESPACE:",n.Attr[1].Namespace)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			engine(c)
		}
	}
	engine(doc)
}



func includes(slice []string,v string) bool{
	for _,value := range slice{
		if value == v{
			return true
		}
	}
	return false
}


