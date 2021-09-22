package main

import "github.com/valyala/fasthttp"

func doRequest(url string) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req) // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(url)

	err := fasthttp.Do(req, resp)
	if err != nil {
		return 
	}

	bodyBytes := resp.Body()
	println(string(bodyBytes))
}

func main()  {
	doRequest("https://google.com")
}