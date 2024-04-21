package main

import (
	"assigment3/another"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", another.Login)
	http.HandleFunc("/home", another.Home)
	http.HandleFunc("/refresh", another.Refresh)

	//log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
