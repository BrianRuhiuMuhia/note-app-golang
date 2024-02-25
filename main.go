package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var (
	PORT string = ":5000"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/addnote", addNote)
	fmt.Println("server running on port, ", PORT)
	http.ListenAndServe(PORT, nil)
}
func addNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	reqBody, IoErr := io.ReadAll(r.Body)
	if IoErr != nil {
		fmt.Fprintf(w, "err: %+v\n", IoErr)
	}
	err := json.Unmarshal(reqBody, &note)
	if err != nil {
		fmt.Fprintf(w, "err: %+v\n", err)
	}

}
