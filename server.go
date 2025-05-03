package configs

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ты молодец у тебя получилось сделать это!")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8080", nil)

}
