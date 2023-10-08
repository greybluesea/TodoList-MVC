package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

func main() {

	http.HandleFunc("/", todolistHandler)
	http.HandleFunc("/todolist", todolistHandler)

	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/delete", deleteHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func helloHandler(writer http.ResponseWriter,
	request *http.Request) {
	_, err := writer.Write([]byte(`
	Hello, this is a todo list web app,
	written in Go lang,
	based on Model–view–controller (MVC) pattern,
	Learned from Derek Banas,
	powered by greybluesea`))
	errorCheck(err)
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
	//todo += "\n"
	//_, err = fmt.Fprint(file, todo)
	_, err = fmt.Fprintln(file, todo)
	errorCheck(err)
	// Close file
	err = file.Close()
	errorCheck(err)
	// Redirect to defined page while passing
	// ResponseWriter, original request,
	// and a successful request message
	http.Redirect(writer, request, "/todolist", http.StatusFound)
}

func deleteHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the index of the To-Do item to delete from the form
	indexStr := request.FormValue("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		http.Error(writer, "Invalid index", http.StatusBadRequest)
		return
	}

	// Get the current list of To-Do items
	todoStrs := getTodos("todos.txt")

	// Check if the index is valid
	if index >= 0 && index < len(todoStrs) {
		// Remove the item at the specified index
		todoStrs = append(todoStrs[:index], todoStrs[index+1:]...)

		// Write the updated list of To-Do items back to the file
		err := writeTodos("todos.txt", todoStrs)
		if err != nil {
			http.Error(writer, "Failed to update To-Do list", http.StatusInternalServerError)
			return
		}
	}

	// Redirect to the To-Do list page
	http.Redirect(writer, request, "/todolist", http.StatusFound)
}

// Add a function to write the updated To-Do list to the file
func writeTodos(fileName string, todos []string) error {
	// Open file with options and permissions
	options := os.O_WRONLY | os.O_TRUNC | os.O_CREATE
	file, err := os.OpenFile(fileName, options, os.FileMode(0600))
	if err != nil {
		return err
	}
	defer file.Close()

	// Write each To-Do item to the file
	for _, todo := range todos {
		_, err := fmt.Fprintln(file, todo)
		if err != nil {
			return err
		}
	}
	return nil
}

/* func castAndWrite(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	errorCheck(err)

} */
