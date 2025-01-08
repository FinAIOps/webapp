package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed ui/build/*
var embeddedFiles embed.FS

func main() {
	// Create a subdirectory file system to strip the 'ui/build' prefix
	subFS, err := fs.Sub(embeddedFiles, "ui/build")
	if err != nil {
		panic(err)
	}

	fileServer := http.FileServer(http.FS(subFS))

	// Handle root by serving index.html for SPA (Single Page Application)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			// Serve static files normally
			fileServer.ServeHTTP(w, r)
			return
		}
		// Serve index.html for root path
		http.ServeFile(w, r, "ui/build/index.html")
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
