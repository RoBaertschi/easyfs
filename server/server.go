package server

import (
	"log"
	"net/http"
	"os"

	"robaertshi.xyz/easyfs/config"
)

type Handler struct {
	HttpHandler http.Handler
	Route       string
}

// Return false, if the next handlers and middlewares should not be called
type MiddlewareHandler = func(w *http.ResponseWriter, r *http.Request) bool

type funcHandler struct {
	WrappedFunc func(w http.ResponseWriter, r *http.Request)
}

func (h funcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.WrappedFunc(w, r)
}

var middlewares []MiddlewareHandler = make([]MiddlewareHandler, 0)
var handlers []Handler = make([]Handler, 0)

func AddMiddleware(middleware MiddlewareHandler) {
	middlewares = append(middlewares, middleware)
}

func Handle(route string, handler http.Handler) {
	handlers = append(handlers, Handler{Route: route, HttpHandler: handler})
}

func HandleFunc(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	handlers = append(handlers, Handler{Route: route, HttpHandler: funcHandler{WrappedFunc: handler}})
}

func StartServer(config *config.Config, logger *log.Logger) error {
	root := http.NewServeMux()
	fileServerMux := http.NewServeMux()

	// router.Handle("/static/", )
	fs := os.DirFS(config.ServeDirectory)
	fileserver := http.FileServerFS(fs)
	fileServerMux.Handle("/", fileserver)

	// root.Handle("/static/", http.StripPrefix("/static", fileServerMux))
	root.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		for middleware := range middlewares {
			if !middlewares[middleware](&w, r) {
				return
			}
		}
		http.StripPrefix("/static", fileServerMux).ServeHTTP(w, r)
	})

	for _, handler := range handlers {
		root.HandleFunc(handler.Route, func(w http.ResponseWriter, r *http.Request) {
			for middleware := range middlewares {
				if !middlewares[middleware](&w, r) {
					return
				}
			}
			handler.HttpHandler.ServeHTTP(w, r)
		})
	}

	logger.Print("Started server on http://localhost:3000")
	logger.Printf("Serving all files in %s on http://localhost:3000/static", config.ServeDirectory)
	return http.ListenAndServe(":3000", root)
}
