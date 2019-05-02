package easy

import (
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
	if err := PrintRequest(req); err != nil {
		log.Printf("request print failed: %s", err)
	}
	// change request url to do it to another end of the pipe
	url, err := TransformURL(req.URL, handler.RequestAddress)
	if err != nil {
		log.Printf("transform url failed: %s", err)
	}

	new_req, err := http.NewRequest(req.Method, url, req.Body)
	if err != nil {
		if e := ErrorResponse(&w, err); e != nil {
			log.Printf("return error response failed: %s", err)
		}
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
	resp, err := DoRequest(req)
	if err != nil {
		log.Printf("do request to backend server failed: %s", err)
		if e := ErrorResponse(&w, err); e != nil {
			log.Printf("return error response failed: %s", err)
		}
		return
	}
	if err = PrintResponse(resp); err != nil {
		log.Printf("print response failed: %s", err)
	}
	// rewrite response to pipe
	if err := NormalResponse(&w, resp); err != nil {
		log.Printf("[x] response return failed: %s", err)
	}
}
