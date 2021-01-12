package fbrip

import(
	"io"
	"fmt"
	"bytes"
	"net/url"
	"net/http"
)

//Ensuring that URL follows to `mbasic.facebook.com`
func fixUrl(Url *url.URL) *url.URL{
	Url.Scheme = "https"
	Url.Host = "mbasic.facebook.com"
	return Url
}

//It checks if `cookie` is already in `slice of cookies`
func includesCookie(cookies []*http.Cookie, cookie *http.Cookie) bool{
	for _,c := range cookies{
		if c.Name == cookie.Name{
			return true
		}
	}
	return false
}

// io.Reader -> string
func bodyToString(body io.Reader) string{
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	return buf.String()
}
