package router

import "net/http"

type Router interface {
	Logger(use bool)
	Cors(use bool)

	GROUP(uri string, rg Group)
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	PUT(uri string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(uri string, f func(w http.ResponseWriter, r *http.Request))

	SERVE(port string)
}

type Group struct {
	GET    func(w http.ResponseWriter, r *http.Request)
	POST   func(w http.ResponseWriter, r *http.Request)
	PUT    func(w http.ResponseWriter, r *http.Request)
	DELETE func(w http.ResponseWriter, r *http.Request)
}
