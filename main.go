package main

import (
	"fmt"

	"github.com/OscarMauricio666/BackAppLife/Connection"
	"github.com/OscarMauricio666/BackAppLife/Test"
)

// func home(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./index.hCml")
// }

func main() {
	fmt.Println(Test.Prueba())
	fmt.Println(Connection.Conectar())
}

// http.HandleFunc("/", home)
// http.ListenAndServe(":3000", nil)
