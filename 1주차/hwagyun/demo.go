// demo.go 파일 생성
// 단순한 웹 서버입니다.
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, SLSA Level 3!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server running on :8080")
    http.ListenAndServe(":8080", nil)
}