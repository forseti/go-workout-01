package main

import (
	"forseti.github.io/goworkout/c08di"
	"net/http"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	c08di.Greet(w, "world")
}

func main()  {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
