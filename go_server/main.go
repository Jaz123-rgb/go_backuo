package main

import (
	"fmt"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Addresss = %s\n", address)
}

func HelloHandler() {

}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

}
