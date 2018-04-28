package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type MyHandler map[string]dollars

func (self MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range self {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}

	fmt.Println(req.Header)
	switch req.URL.Path {
	case "/list":
		for item, price := range self {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := self[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		w.Header().Add("IFRAME", "SAMEORIGIN")
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

/////////
func (self MyHandler) rootPath(w http.ResponseWriter, req *http.Request) {
	for item, price := range self {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
	w.WriteHeader(http.StatusNotFound) // 404
	w.Header().Add("IFRAME", "SAMEORIGIN")
	fmt.Fprintf(w, "no such page: %s\n", req.URL)
}
func (self MyHandler) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range self {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
	w.Header().Add("key1", "value1")
	w.Header().Set("key2", "value1")
	w.Header().Set("Content-Type", "text/html")

	html := `<doctype html>
        <html>
        <head>
          <title>Hello World</title>
        </head>
        <body>
        <p>
          <a href="/welcome">Welcome</a> |  <a href="/message">Message</a>
        </p>
        </body>
</html>`
	fmt.Fprintln(w, html)

}
func (self MyHandler) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := self[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
func text(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Comleted %s in %v", r.URL.Path, time.Since(start))
	})
}

func hook(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("before hook")
		next.ServeHTTP(w, r)
		log.Println("after hook")

	})
}

///////

func main() {
	handler := MyHandler{"shoes": 50, "socks": 5}
	///
	///log.Fatal(http.ListenAndServe("localhost:8000", handler))
	////
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(handler.list))
	mux.Handle("/price", http.HandlerFunc(handler.price))
	mux.Handle("/text", hook(loggingHandler(http.HandlerFunc(text))))
	mux.Handle("/", http.HandlerFunc(handler.rootPath))
	log.Fatal(http.ListenAndServe(":8000", mux))

}
