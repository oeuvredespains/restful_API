package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// POST
func createStudent(w http.ResponseWriter, r *http.Request) {
	// get data from the body
	data := r.Body
	// data to struct
	s := student{}
	// decode the data
	err := json.NewDecoder(data).Decode(&s)
	if err != nil {
		fmt.Println(err)
		// return error code
		w.WriteHeader(http.StatusBadRequest)
	}
	// TODO add student to DB
	w.WriteHeader(http.StatusCreated)
}

// GET
func getStudent(w http.ResponseWriter, r *http.Request) {
	// get ID from the url
	vars := mux.Vars(r)
	id := vars["id"]
	// TODO get student from DB
}

// GET ALL
func getAllStudent(w http.ResponseWriter, r *http.Request) {
	// TODO get all students from DB
}

// PUT
func updateStudent(w http.ResponseWriter, r *http.Request) {
	// TODO update something
}

// DELETE
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	// get ID from the url
	vars := mux.Vars(r)
	id := vars["id"]
	// TODO delete student from DB
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
