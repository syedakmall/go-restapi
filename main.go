package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/student/profile/", getStudentById)
	http.HandleFunc("/students", getAllStudents)
	http.HandleFunc("/student/new", createStudent)
	http.HandleFunc("/student/delete", deleteStudent)

	http.ListenAndServe(":8080", nil)
}
