package server

import (
	"log"
	"net/http"
	"os"

	"robaertshi.xyz/easyfs/config"
)

func StartServer(config *config.Config, logger *log.Logger) error {
	root := http.NewServeMux()
	fileServerMux := http.NewServeMux()

	// router.Handle("/static/", )
	fs := os.DirFS(config.ServeDirectory)
	fileserver := http.FileServerFS(fs)
	fileServerMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("called %+v", fileserver)
		fileserver.ServeHTTP(w, r)
	})

	root.Handle("/static/", http.StripPrefix("/static", fileServerMux))

	logger.Print("Started server on http://localhost:3000")
	logger.Printf("Serving all files in %s on http://localhost:3000/static", config.ServeDirectory)
	return http.ListenAndServe(":3000", root)
}
