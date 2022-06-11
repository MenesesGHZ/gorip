package fbrip

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/url"
	"strconv"
	"strings" // just to make the first letter uppercase ;d
)

// Declaring core function for the search engine function
type coreFunc func(int, *goquery.Selection)

// Search engine
func searchEngine(body io.Reader, query string, f coreFunc) *goquery.Document {
	// Creating doc by body
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		panic("Error while reading utf-8 enconded HTML")
	}
	doc.Find(query).Each(f)

	//returning document in case if is need
	return doc
}

func searchParamsForUser(body io.Reader, u *UserRip) {
	// Finding user's parameters that are in input tags
	searchEngine(body, "input", func(i int, s *goquery.Selection) {
		name, nOk := s.Attr("name")
		value, vOk := s.Attr("value")
		if nOk && vOk {
			for _, key := range u.GetParameterKeys() {
				if key == value {
					u.Parameters[name] = value
					break
				}
			}
		}
	})
}

func searchBasicInfo(body io.Reader) map[string]string {
	//Searching attributes base on `searchList`
	searchList := []string{"birthday", "gender"}

	// Making map for basic info
	bi := make(map[string]string)

	// Searching path: 1*<div id="basic-info"> -> 6*<a>
	//(<a> contains href which helps to determine what type of info attribute we are dealing)
	doc := searchEngine(body, "div#basic-info a", func(i int, a *goquery.Selection) {

		// Parsing url from href to then get the values to determine attr
		hrefValue, hOk := a.Attr("href")
		hUrl, _ := url.Parse(hrefValue)
		if hOk {
			//Getting Query Parameters from `hUrl`
			v := hUrl.Query()
			for _, element := range searchList {
				if element == v.Get("edit"){
					key := strings.Title(v.Get("edit"))
					bi[key] = a.Parent().Parent().Parent().Next().Children().Text()
				}
			}
		}
	})

	//Getting user's name
	doc.Find("title").Each(func(i int, t *goquery.Selection) {
		bi["Name"] = t.Text()
	})
	return bi
}

func searchReactionPickerUrl(body io.Reader) *url.URL {
	//Declaring URL var
	var Url *url.URL
	// Looking for ActionBar where its patern path is: tbody > tr > td > a
	searchEngine(body, "tbody tr td a", func(i int, a *goquery.Selection) {
		// Finding the url that follows the path == `/reaction/picker/`
		hrefValue, hOk := a.Attr("href")
		if hOk {
			hUrl, _ := url.Parse(hrefValue)
			if hUrl.Path == "/reactions/picker/" {
				Url = fixUrl(hUrl)
			}
		}
	})
	return Url
}

func searchUfiReactionUrl(body io.Reader, reactId string) *url.URL {
	//Declaring Url & Converting `reactId` to integer
	var Url *url.URL
	id, err := strconv.Atoi(reactId)
	if err != nil {
		panic("Reaction ID must be string")
	}
	// Looking for ActionBar where its patern path is: tbody > tr > td > a
	searchEngine(body, "tbody tr td a", func(i int, a *goquery.Selection) {
		// Finding the url that follows the path == `/reaction/picker/`
		hrefValue, hOk := a.Attr("href")
		if hOk && i == id {
			Url, _ = url.Parse(hrefValue)
			Url = fixUrl(Url)
		}
	})
	return Url
}
