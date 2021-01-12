package fbrip

import(
	"io"
	"fmt"
	"net/url"
	"strings" // just to make the first letter uppercase ;d
	"github.com/PuerkitoBio/goquery"
)

// Declaring core function for the search engine function
type coreFunc func(int,*goquery.Selection)
// Search engine
func searchEngine(body io.Reader, query string, f coreFunc){
	// Creating doc by body
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		panic("Error while reading utf-8 enconded HTML")
	}
	doc.Find(query).Each(f)
}


func searchParamsForUser(body io.Reader,u *UserRip){
	// Finding user's parameters that are in input tags
	searchEngine(body,"input",func(i int, s *goquery.Selection){
		name,nOk := s.Attr("name")
		value,vOk := s.Attr("value")
		if nOk && vOk {
			if includes(u.GetParameterKeys(),name){
				u.Parameters[name] = value
			}
		}
	})
}

func searchBasicInfo(body io.Reader) map[string]string{
	//Searching attributes base on `searchList`
	searchList := []string{"birthday","gender",}
	// Making map for basic info
	bi := make(map[string]string)
	// Searching path: 1*<div id="basic-info"> -> 6*<a>
	//(<a> contains href which helps to determine what type of info attribute we are dealing)
	searchEngine(body,"div#basic-info a",func(i int,s *goquery.Selection){
		// Parsing url from href to then get the values to determine attr
		hrefValue,hOk := s.Attr("href")
		hUrl,_ := url.Parse(hrefValue)
		v := hUrl.Query()
		if includes(searchList,v.Get("edit")) && hOk{
			key := strings.Title(v.Get("edit"))
			// a < span < div < td - td > div > InfoAttribute 
			bi[key] = s.Parent().Parent().Parent().Next().Children().Text()
		}
	})
	//Getting user's name
	searchEngine(body,"title",func(i int,s *goquery.Selection){
		bi["Name"] = s.Text()
	})
	return bi
}


//func searchQueryParameters(body io.Reader){
//	
//}

