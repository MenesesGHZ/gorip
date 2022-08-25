package fbrip

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/url"
	"strconv"
	"strings" // just to make the first letter uppercase ;d
)

// Finding user's parameters that are in input tags
func searchParamsForUser(body io.Reader, u *UserRip) {
	document, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		panic("Error while reading utf-8 enconded HTML")
	}
	document.Find("input").Each(func(i int, s *goquery.Selection) {
		name, nOk := s.Attr("name")
		value, vOk := s.Attr("value")
		if nOk && vOk {
			for _, key := range u.GetParameterKeys() {
				if key == name {
					u.Parameters[name] = value
					break
				}
			}
		}
	})
}

// Searching path: 1*<div id="basic-info"> -> 6*<a>
// (<a> contains href which helps to determine what type of info attribute we are dealing)
func searchBasicInfo(body io.Reader) map[string]string {
	document, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		panic("Error while reading utf-8 enconded HTML")
	}
	searchList := []string{"birthday", "gender"}
	basicInfoMap := make(map[string]string)
	document.Find("div#basic-info a").Each(func(i int, a *goquery.Selection) {
		hrefValue, hOk := a.Attr("href")
		hUrl, _ := url.Parse(hrefValue)
		if hOk {
			v := hUrl.Query()
			for _, element := range searchList {
				if element == v.Get("edit") {
					key := strings.Title(v.Get("edit"))
					basicInfoMap[key] = a.Parent().Parent().Parent().Next().Children().Text()
				}
			}
		}
	})
	document.Find("title").Each(func(i int, t *goquery.Selection) {
		basicInfoMap["Name"] = t.Text()
	})
	return basicInfoMap
}

// Looking for ActionBar where its patern path is: tbody > tr > td > a
func getReactionsPickerUrl(body io.Reader) *url.URL {
	document, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		panic("Error while reading utf-8 enconded HTML")
	}
	var reactionsPickerUrl *url.URL
	document.Find("tbody tr td a").Each(func(i int, a *goquery.Selection) {
		hrefValue, hOk := a.Attr("href")
		if hOk {
			Url, _ := url.Parse(hrefValue)
			if Url.Path == "/reactions/picker/" {
				transformUrlToBasicFacebook(Url)
				reactionsPickerUrl = Url
			}
		}
	})
	return reactionsPickerUrl
}

// Declaring Url & Converting `reactId` to integer
// Looking for ActionBar where its patern path is: tbody > tr > td > a
func getReactionUrl(body io.Reader, reactId string) *url.URL {
	document, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		panic("Error while reading utf-8 enconded HTML")
	}
	var reactionUrl *url.URL
	id, _ := strconv.Atoi(reactId)
	document.Find("tbody tr td a").Each(func(i int, a *goquery.Selection) {
		hrefValue, hOk := a.Attr("href")
		if hOk && i == id {
			reactionUrl, _ = url.Parse(hrefValue)
			transformUrlToBasicFacebook(reactionUrl)
		}
	})
	return reactionUrl
}
