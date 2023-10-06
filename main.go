package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type TodoList struct {
	Count int
	Todos []string
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func write(resWriter http.ResponseWriter, msg string) {
	_, err := resWriter.Write([]byte(msg))
	errorCheck(err)

}

func englishHandler(resWriter http.ResponseWriter,
	request *http.Request) {
	write(resWriter, `
	Hello, this is a todo list web app,
	written in Go lang,
	based on Model–view–controller (MVC) pattern,
	Learned from Derek Banas,
	powered by greybluesea`)
}

func spanishHandler(writer http.ResponseWriter,
	request *http.Request) {
	write(writer, "Hola Internet")
}

func frenchHandler(writer http.ResponseWriter,
	request *http.Request) {
	write(writer, "Bonjour Internet")
}

func todolistHandler(writer http.ResponseWriter,
	request *http.Request) {

	// Get our text from the file
	todoStrs := getTodos("todos.txt")

	// Print to the terminal
	fmt.Printf("%#v\n", todoStrs)
	// Create a template using the html
	tmpl, err := template.ParseFiles("view.html")
	errorCheck(err)

	// Create a todo list with the number
	todos := TodoList{
		Count: len(todoStrs),
		Todos: todoStrs,
	}

	// Write the template to the ResponseWriter
	// Pass the todo struct data
	err = tmpl.Execute(writer, todos)
	errorCheck(err)
}

func getTodos(fileName string) []string {
	var todoStrs []string

	// Try to open the file (It must exist)
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	errorCheck(err)
	// Close file when the function ends
	defer file.Close()

	// Read lines of text and save to lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		todoStrs = append(todoStrs, scanner.Text())
	}
	errorCheck(scanner.Err())

	// Return the text
	return todoStrs
}

func newHandler(writer http.ResponseWriter,
	request *http.Request) {

	// Create a template using the html
	tmpl, err := template.ParseFiles("new.html")
	errorCheck(err)

	err = tmpl.Execute(writer, nil)
	errorCheck(err)
}

func createHandler(writer http.ResponseWriter,
	request *http.Request) {
	todo := request.FormValue("todo")
	// Define options for working with the file
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	// Open file with options and permissions
	file, err := os.OpenFile("todos.txt", options, os.FileMode(0600))
	errorCheck(err)
	// Append new text to file
	todo += "\n"
	_, err = fmt.Fprint(file, todo)
	errorCheck(err)
	// Close file
	err = file.Close()
	errorCheck(err)
	// Redirect to defined page while passing
	// ResponseWriter, original request,
	// and a successful request message
	http.Redirect(writer, request, "/todolist", http.StatusFound)
}
func main() {
	// Our app is available at directory
	// hello for the localhost port 8080
	// When it receives a request it calls
	// the correct Handler
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/hola", spanishHandler)
	http.HandleFunc("/bonjour", frenchHandler)
	http.HandleFunc("/todolist", todolistHandler)
	http.HandleFunc("/", todolistHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/create", createHandler)

	// Listens for browser requests and responds
	// Only receives a value if there is an error
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
