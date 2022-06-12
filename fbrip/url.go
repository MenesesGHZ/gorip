package fbrip

import (
	"fmt"
	"net/url"
)


var FacebookUrl, _ = url.Parse("https://www.facebook.com/")
var BasicFacebookUrl, _ = url.Parse("https://mbasic.facebook.com/")


func transformUrlToBasicFacebook(baseUrl *url.URL){
	baseUrl.Scheme = BasicFacebookUrl.Scheme
	baseUrl.Host = BasicFacebookUrl.Host
}

func transformUrlToFacebook(baseUrl *url.URL){
	baseUrl.Scheme = FacebookUrl.Scheme
	baseUrl.Host = FacebookUrl.Host
}

func parseUrls(urls []string) []*url.URL {
	var Urls []*url.URL
	for i, Url := range urls {
		tempUrl, err := url.Parse(Url)
		if err != nil {
			fmt.Printf("Error while parsing #%d -> %s\nIt was omitted.\n", i, Url)
			continue
		}
		Urls = append(Urls, tempUrl)
	}
	return Urls
}
