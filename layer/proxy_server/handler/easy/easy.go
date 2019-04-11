package easy

import (
	handlerPkg "local/proxy_server/handler"
	"log"
	"net/http"
)

type Handler struct {
	RequestAddress string
}

func easyHandle(req *http.Request) *http.Request {

	return req
}

func (handler *Handler) Handle(w http.ResponseWriter, req *http.Request) {
	handlerPkg.PrintRequest(req)
	// change request url to do it to another end of the pipe
	url, err := handlerPkg.TransformURL(req.URL, handler.RequestAddress)

	new_req, err := http.NewRequest(req.Method, url, req.Body)
	if err != nil {
		handlerPkg.ErrorResponse(&w, err)
		return
	}

	for k, v := range req.Header {
		new_req.Header[k] = v
	}

	for _, cookie := range req.Cookies() {
		new_req.AddCookie(cookie)
	}

	// change request data
	req = easyHandle(new_req)

	// do the request to another end of the pipe
	resp, err := handlerPkg.DoRequest(req)
	handlerPkg.PrintResponse(resp)
	if err != nil {
		handlerPkg.ErrorResponse(&w, err)
		return
	}
	// rewrite response to pipe
	if err := handlerPkg.NormalResponse(&w, resp); err != nil {
		log.Fatal("[x] response return failed: %s", err)
	}
}
