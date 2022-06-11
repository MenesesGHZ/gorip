package fbrip

import (
	"net/url"
)

//Ensuring that URL follows to `mbasic.facebook.com`
func fixUrl(Url *url.URL) *url.URL {
	Url.Scheme = "https"
	Url.Host = "mbasic.facebook.com"
	return Url
}

