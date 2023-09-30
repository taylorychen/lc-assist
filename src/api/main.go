package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/questions", GetQuestions).Methods("GET")
	r.HandleFunc("/", Default)
	r.Use(mux.CORSMethodMiddleware(r))

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}

func Default(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK\n")
}

type question struct {
	Number     int    `json:"number"`
	Name       string `json:"name"`
	Difficulty string `json:"difficulty"`
	Type       string `json:"type"`
	Solved     bool   `json:"solved"`
}

type output struct {
	Questions []question `json:"questions"`
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	printLog("GetQuestions")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	list := []question{
		{1, "BST", "E", "Tree", true},
		{2, "Trie", "E", "Trie", true},
		{3, "Dict", "M", "Hashmap", false},
		{4, "Linked List", "M", "Two pointer", true},
		{5, "Enumerate", "M", "Python", true},
		{6, "Queue", "M", "Queue", true},
		{7, "Breadth-First-Search", "M", "Graph", false},
		{8, "Kruskal's Algorithm", "H", "Graph", false},
		{9, "Binary Tree", "E", "Tree", true},
		{10, "Prim's Algorithm", "M", "Graph", false},
		{11, "Dijkstra", "M", "Graph", true},
		{12, "Hello World", "E", "String", true},
	}
	out := output{
		Questions: []question{},
	}
	for _, q := range list {
		out.Questions = append(out.Questions, q)
	}
	res, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func printLog(fn string) {
	fmt.Printf("Ran: %s\n", fn)
}
