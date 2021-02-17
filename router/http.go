package router

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/samsammn/project7-test/response"
)

type httpRouter struct{}

func NewHttpRouter() Router {
	return &httpRouter{}
}

func errorMethodNotAllowed(w http.ResponseWriter) {
	response.Json(w, response.Map{
		"status":  http.StatusMethodNotAllowed,
		"message": "i'm so sorry, you are access endpoint with a method not allowed!",
	})
}

var cors bool = false
var logger bool = false

func useLogger(req *http.Request) {
	if logger {
		RFC850 := "Monday, 02-Jan-06 15:04:05 MST"
		now := time.Now().Format(RFC850)

		fmt.Printf("%v -> %s: %s\n", now, req.Method, req.URL.Path)
	}
}

func (h *httpRouter) Logger(use bool) {
	logger = use
}

func (h *httpRouter) Cors(use bool) {
	cors = use
}

func useCors(w http.ResponseWriter) {
	if cors {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}
}

func (h *httpRouter) GROUP(uri string, rg Group) {
	regex := regexp.MustCompile(`{[^{}]*}`)
	res := regex.FindString(uri)
	newUri := strings.Replace(uri, res, "", -1)

	http.HandleFunc(newUri, func(resp http.ResponseWriter, req *http.Request) {
		useLogger(req)
		useCors(resp)

		if req.Method == http.MethodGet {
			rg.GET(resp, req)
		} else if req.Method == http.MethodPost {
			rg.POST(resp, req)
		} else if req.Method == http.MethodPut {
			rg.PUT(resp, req)
		} else if req.Method == http.MethodDelete {
			rg.DELETE(resp, req)
		}
	})
}

func (h *httpRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(uri, func(resp http.ResponseWriter, req *http.Request) {
		useLogger(req)
		useCors(resp)

		if req.Method == http.MethodGet {
			f(resp, req)
		} else {
			errorMethodNotAllowed(resp)
		}
	})
}

func (h *httpRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(uri, func(resp http.ResponseWriter, req *http.Request) {
		useLogger(req)
		useCors(resp)

		if req.Method == http.MethodPost {
			f(resp, req)
		} else {
			errorMethodNotAllowed(resp)
		}
	})
}

func (h *httpRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(uri, func(resp http.ResponseWriter, req *http.Request) {
		useLogger(req)
		useCors(resp)

		if req.Method == http.MethodPut {
			f(resp, req)
		} else {
			errorMethodNotAllowed(resp)
		}
	})
}

func (h *httpRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(uri, func(resp http.ResponseWriter, req *http.Request) {
		useLogger(req)

		if req.Method == http.MethodDelete {
			f(resp, req)
		} else {
			errorMethodNotAllowed(resp)
		}
	})
}

func (h *httpRouter) SERVE(port string) {
	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func Arg(req *http.Request) string {
	sliceArgs := strings.Split(req.URL.Path, "/")
	lenArg := len(sliceArgs)

	if lenArg > 2 {
		return sliceArgs[lenArg-1]
	} else {
		return ""
	}
}
