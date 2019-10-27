package main

import "fmt"
import "http"

func hanlder(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome")
}