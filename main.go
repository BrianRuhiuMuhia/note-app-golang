package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var (
	PORT     string = ":5000"
	notes    []string
	dirname  string = "./notes"
	filename string = "data.json"
	filepath string = dirname + "/" + filename
	file     *os.File
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/addnote", addNote)
	http.HandleFunc("/getallnotes", getAllNotes)
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
	if _, err := os.Stat("notes"); os.IsNotExist(err) {
		fmt.Println("Directory does not exist")
		err := os.Mkdir(dirname, 0755)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Directory Created Successfully")
	}
	file, fileError := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE, 0644)
	if fileError != nil {
		fmt.Println(fileError)
	}
	noteJson, err := json.Marshal(note)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(noteJson)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	defer file.Close()
}
func getAllNotes(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat("notes"); os.IsNotExist(err) {
		fmt.Println("Directory does not exist")
		fmt.Fprintf(w, "server error directory does not exist")
	}
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "There was an error")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		note, err := json.Marshal(scanner.Text())
		if err != nil {
			fmt.Println("there was an error line 79")
		}
		notes = append(notes, string(note))
		for _, note := range notes {
			fmt.Fprintf(w, note)
		}
	}

}
