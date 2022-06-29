package main

import (
	"io"
	"net/http"
)

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("merhaba mars"))
	// })
	// http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/index ", indexHandler)
	// http.HandleFunc("/about", aboutHandler)

	//	http.ListenAndServe(":8080", nil)

	var i ironman
	var w wolwerine
	mux := http.NewServeMux()
	mux.Handle("/ironman", i)
	mux.Handle("/wolwerine", w)
	http.ListenAndServe(":8080", mux)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Zort"))
	x := r.URL.Path[1:]
	if !(len(x) > 0) {

	}
	w.Write([]byte(x))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

type ironman int

func (x ironman) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr. Iron!")
}

type wolwerine int

func (x wolwerine) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Wolwerine!")
}

type green struct {
	Name string
}

func (x *green) TheyAreSoStrong() int {
	return 5
}
