package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// POST
func createStudent(w http.ResponseWriter, r *http.Request) {
	// TODO create something
}

// GET
func getStudent(w http.ResponseWriter, r *http.Request) {
	// TODO get something
}

// GET ALL
func getAllStudent(w http.ResponseWriter, r *http.Request) {
	// TODO get all students
}

// PUT
func updateStudent(w http.ResponseWriter, r *http.Request) {
	// TODO update something
}

// DELETE
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	// TODO delete something
}

// Homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Here !")
}

// Handle the requests
func handleRequests() {
	// create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// attach the homePage route and handler to the router
	myRouter.HandleFunc("/", homePage)
	// function per route
	// example of some routes to get or create students
	myRouter.HandleFunc("/students", getAllStudent).Methods("GET")        // GET
	myRouter.HandleFunc("/student", createStudent).Methods("POST")        // CREATE
	myRouter.HandleFunc("/student/{id}", getStudent).Methods("GET")       // READ
	myRouter.HandleFunc("/student/{id}", updateStudent).Methods("PUT")    // UPDATE
	myRouter.HandleFunc("/student/{id}", deleteStudent).Methods("DELETE") // DELETE
	// start the server on port 8080 and log any errors
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	// TODO coonect to DB
	println("Listening on http://localhost:8080")
	handleRequests()
}
