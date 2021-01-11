package fbrip

import "net/url"


type info struct{
	Name string
}

type react struct{
	Url *url.URL
	Id string
}

type post struct{
	Url *url.URL
	Content string
}

type comment struct{
	Url *url.URL
	Content string
}
