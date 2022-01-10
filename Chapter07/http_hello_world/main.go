package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<html>\n<body>\n"))
	w.Write([]byte("<p>Hello World!</p>\n"))
	w.Write([]byte("</body>\n</html>\n"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting http server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
