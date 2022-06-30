package main

import (
	"fmt"
	"io"
	"io/ioutil"
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

	// var i ironman
	// var w wolwerine
	// mux := http.NewServeMux()
	// mux.Handle("/ironman", i)
	// mux.Handle("/wolwerine", w)
	// http.ListenAndServe(":8080", mux)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5555", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	var body, _ = loadFile("html/index.html")
	fmt.Fprint(w, body)
}

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Zort"))
	x := r.URL.Path[1:]
	if !(len(x) > 0) {

	}
	w.Write([]byte(x))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Abox ut Page"))
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
