package main

import (
	"html/template"
	"os"
)

type Dog struct {
	Name string
}

type User struct {
	Name string
	Dog  Dog
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")

	if err != nil {
		panic(err)
	}

	data := User{
		Name: "Stoyan Stoyanov",
		Dog:  Dog{Name: "Rex"},
		Age:  18,
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	data.Name = "John Smith"
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
