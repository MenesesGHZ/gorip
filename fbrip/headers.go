package fbrip


import(
	"strconv"
	"net/http"
)

func setHeaders(request *http.Request, contentType string, paramsLength int){
	//Setting default headers
	request.Header.Set("Host",request.URL.Host)
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	request.Header.Set("Accept-Language", "en-US,en;q=1.0")
	request.Header.Set("Connection", "close")
	request.Header.Set("Upgrade-Insecure-Requests", "1")

	//Setting parameters if POST request
	if request.Method == "POST"{
		request.Header.Set("Content-Type",contentType)
		request.Header.Set("Content-Length", strconv.Itoa(paramsLength))
		request.Header.Set("Origin", request.URL.String())
	}
}
