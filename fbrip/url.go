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

func parseUrls(rawUrls []string) []*url.URL {
	var parsedUrls []*url.URL
	for i, rawUrl := range rawUrls {
		parsedUrl, err := url.Parse(rawUrl)
		if err != nil {
			fmt.Printf("Error while parsing #%d -> %s\nIt was omitted.\n", i, rawUrl)
			continue
		}
		parsedUrls = append(parsedUrls, parsedUrl)
	}
	return parsedUrls
}
