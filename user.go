package main

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Students struct {
	Students []Student `json:"users"`
}

var students []Student

func remove[T comparable](arr *[]T, target T) {
	var index int
	v := *arr

check:
	for i := 0; i < len(v); i++ {
		if v[i] == target {
			index = i
			break check
		}
	}
	*arr = append(v[:(index)], v[index+1:]...)
}
