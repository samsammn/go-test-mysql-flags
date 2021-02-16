package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/samsammn/project7-test/response"
)

type httpRouter struct {
}

func NewHttpRouter() Router {
	return &httpRouter{}
}

func errorMethodNotAllowed(w http.ResponseWriter) {
	response.Json(w, response.Map{
		"status":  http.StatusMethodNotAllowed,
		"message": "i'm so sorry, you are access endpoint with a method not allowed!",
	})
}

func (h *httpRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(uri, func(resp http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			f(resp, req)
		} else {
			errorMethodNotAllowed(resp)
		}
	})
}

func (h *httpRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(uri, func(resp http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			f(resp, req)
		} else {
			errorMethodNotAllowed(resp)
		}
	})
}

func (h *httpRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(uri, func(resp http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPut {
			f(resp, req)
		} else {
			errorMethodNotAllowed(resp)
		}
	})
}

func (h *httpRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(uri, func(resp http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodDelete {
			f(resp, req)
		} else {
			errorMethodNotAllowed(resp)
		}
	})
}

func (h *httpRouter) SERVE(port string) {
	fmt.Println("Server running on port :" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
