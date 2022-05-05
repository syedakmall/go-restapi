package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func getStudentById(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		errorResponse(w, "Id Invalid", http.StatusBadRequest)
		return
	}
	var u Student

	for _, v := range students {
		if v.Id == id {
			u = v
			data, _ := json.Marshal(u)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(data))
			return
		}
	}

	errorResponse(w, "Id Not Found", http.StatusNotFound)
}

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, "Wrong Method", http.StatusMethodNotAllowed)
		return
	}

	data := Students{Students: students}
	j, err := json.Marshal(data)
	if err != nil {
		errorResponse(w, "Wrong Method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(j))
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, "Wrong Method", http.StatusMethodNotAllowed)
		return
	}

	if h := r.Header.Get("Content-Type"); h != "application/json" {
		errorResponse(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var u Student
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&u)
	if err != nil {
		errorResponse(w, "Error Decoding", http.StatusInternalServerError)
		return
	}
	j, _ := json.Marshal(u)
	students = append(students, u)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(j))
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		errorResponse(w, "Wrong Method", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		errorResponse(w, "Bad Request", http.StatusBadRequest)
		return
	}

	for _, v := range students {
		if v.Id == id {
			remove(&students, v)
			w.Write([]byte("deleted"))
			return
		}

	}

	errorResponse(w, "Id Not Found", http.StatusNotFound)
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
