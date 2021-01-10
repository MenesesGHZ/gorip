package fbrip

import (
	"io"
	"fmt"
	"bytes"
	"net/http"
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

