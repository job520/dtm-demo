package main

import (
	"fmt"
	"net/http"
)

func transIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("transIn")
}
func transInCompensate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("transInCompensate")
}
func transOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("transOut")
}
func transOutCompensate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("transOutCompensate")
}
func main() {
	http.HandleFunc("/trans_in", transIn)
	http.HandleFunc("/trans_in_compensate", transInCompensate)
	http.HandleFunc("/trans_out", transOut)
	http.HandleFunc("/trans_out_compensate", transOutCompensate)
	_ = http.ListenAndServe(":8080", nil)
}
