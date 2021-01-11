package fbrip

import (
	"io"
	"fmt"
	"bytes"
	"net/http"
	"golang.org/x/net/html"
)



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

func showBody(body io.Reader){
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	fmt.Println(buf.String())
}

// Search for input parameters. * It must be improved
func searchParameters(node *html.Node, u *UserRip){
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
