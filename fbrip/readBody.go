package fbrip

import(
	"io"
	"fmt"
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
	// Creating doc by body
	return map[string]string{"a":"b"}
}


//func searchQueryParameters(body io.Reader){
//	
//}

