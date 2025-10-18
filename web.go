package main

import (
	"net/http"
	"fmt"
	"log"
	"satinfo/data_handler"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[1:])
	log.Println("ID:", r.URL.Path[1:])
	fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
	
	var data []data_handler.Response
	data = data_handler.ReadData()

	var sat data_handler.Member

	sat = data_handler.GetByID(id, data)

	var ret data_handler.Member
	if sat == ret {
		fmt.Fprintf(w, "No souch id as %d in database:(\n", id)
	} else {
		fmt.Fprintf(w, "Name: %s\n", sat.Name)
		fmt.Fprintf(w, "\tID: %d\n", sat.Satid)
		fmt.Fprintf(w, "\tLine 1: %s\n", sat.Line1)
		fmt.Fprintf(w, "\tLine 2: %s\n", sat.Line2)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":4321", nil))
}