package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	a()
	fmt.Println(c())

	var dividByHalf Operation = func(i int) int {
		return i / 2
	}

	fmt.Println(Map(dividByHalf, []int{2, 4, 6, 8}))

	d := Dog{}

	DoSomething(d)


	f := fooHandler{}

	g := func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}

	http.Handle("/foo", f)
	
	http.HandleFunc("/bar", g)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

type fooHandler struct {}

func(foo fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

type Animal interface {
	Speak()
}

type Dog struct {}

func (d Dog) Speak() {

	fmt.Print("heyy")
}

func DoSomething(a Animal) {

}

func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}

func c() (i int) {
    defer func() { i++ }()
    return 1
}

type Operation func(int) int


func Map(op Operation, arr []int) []int {
	var newarr []int

	n := len(arr)

	for i:=0; i < n; i++ {
		newarr = append(newarr, op(arr[i]))
	}

	return newarr
}