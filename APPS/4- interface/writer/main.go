/*

type Writer interface {
        Write(p []byte) (n int, err error)
}

*/
package main

import (
    "fmt"
    "net/http"
	"encoding/json"
	"bytes"
)

type user struct {
    Name string `json:"name"`
    Age int `json:"age"`
}

func main() {
    http.HandleFunc("/", HelloServer);
    http.ListenAndServe(":8080", nil);
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:]);
	buf := new(bytes.Buffer);
    customer := user{
        Name: r.URL.Path[1:],
        Age: 20,
    }
	//func NewEncoder(w io.Writer) *Encoder
    err := json.NewEncoder(buf).Encode(customer);
    if err != nil {
        panic(err);
    }
    fmt.Fprintf(w,buf.String());
}