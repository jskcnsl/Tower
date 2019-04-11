package handler

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

func TransformURL(urlToDo *url.URL, target string) (string, error) {
	targetURL, err := url.Parse(target)
	if err != nil {
		return "", err
	}
	targetURL.Path = urlToDo.Path
	targetURL.RawPath = urlToDo.RawPath
	if strings.HasPrefix(targetURL.Path, "/data/ctwaf") {
		targetURL.Path = strings.Replace(targetURL.Path, "/data/ctwaf", "", 1)
	}
	targetURL.ForceQuery = urlToDo.ForceQuery
	targetURL.RawQuery = urlToDo.RawQuery
	targetURL.Fragment = urlToDo.Fragment
	return targetURL.String(), nil
}

func ErrorResponse(w *http.ResponseWriter, err error) error {
	(*w).Header().Add("Content-Type", "application/json")
	io.WriteString(*w, `{"err": "something error: `+err.Error()+`"}`)
	return nil
}

func DoRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NormalResponse(w *http.ResponseWriter, r *http.Response) error {
	for k, _ := range r.Trailer {
		(*w).Header().Add("Trailer", k)
	}
	for k, v := range r.Header {
		for _, every_v := range v {
			(*w).Header().Add(k, every_v)
		}
	}
	(*w).WriteHeader(r.StatusCode)
	_, err := io.Copy(*w, r.Body)
	if err != nil {
		return err
	}
	return nil
}

func PrintRequest(req *http.Request) error {
	fmt.Printf("\n=========================REQUEST======================================\n")

	fmt.Printf("\n[%s]\t%s\nHost: %s\nProto: %s\nRemoteAddr: %s\nRequestURI: %s\nContentLength: %d\n", req.Method, req.URL.String(), req.Host, req.Proto, req.RemoteAddr, req.RequestURI, req.ContentLength)

	if len(req.TransferEncoding) > 0 {
		fmt.Printf("----------------------------------------------------------------------\n")
		fmt.Printf("TransferEncoding: %s\n", req.TransferEncoding)
	}

	fmt.Printf("----------------------------------------------------------------------\n")

	for k, v := range req.Header {
		fmt.Printf("%s:\t%s\n", k, v)
	}

	fmt.Printf("----payload-----------------------------------------------------------\n")

	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	req.Body = ioutil.NopCloser(buf)
	fmt.Printf("%s\n", buf.String())

	return nil
}

func PrintResponse(resp *http.Response) error {
	fmt.Printf("\n=========================RESPONSE=====================================\n")

	fmt.Printf("\n%s\nProto: %s\nContentLength: %d\n", resp.Status, resp.Proto, resp.ContentLength)

	fmt.Printf("----------------------------------------------------------------------\n")

	fmt.Printf("[%s]\t%s\n", resp.Request.Method, resp.Request.URL.String())

	if len(resp.TransferEncoding) > 0 {
		fmt.Printf("----------------------------------------------------------------------\n")
		fmt.Printf("TransferEncoding: %s\n", resp.TransferEncoding)
	}

	fmt.Printf("----------------------------------------------------------------------\n")

	for k, v := range resp.Header {
		fmt.Printf("%s:\t%s\n", k, v)
	}

	fmt.Printf("----payload-----------------------------------------------------------\n")

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	resp.Body = ioutil.NopCloser(buf)
	fmt.Printf("%s\n", buf.String())

	return nil
}

