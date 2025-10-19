package web

import (
	"fmt"
	"log"
	"net/http"
	"satinfo/data_handler"
	"strconv"
)

var data []data_handler.Response

func handler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[1:])
	log.Println("ID:", r.URL.Path[1:])

	sat := data_handler.GetByID(id, data)

	var ret data_handler.Member
	if sat == ret {
		log.Println("No such id as", id, "(", r.URL.Path[1:], ")")
		fmt.Fprintf(w, "No souch id as %d in database:(\n", id)
	} else {
		fmt.Fprintf(w, "Name: %s\n", sat.Name)
		fmt.Fprintf(w, "\tID: %d\n", sat.Satid)
		fmt.Fprintf(w, "\tLine 1: %s\n", sat.Line1)
		fmt.Fprintf(w, "\tLine 2: %s\n", sat.Line2)
	}
}

func Listen(d []data_handler.Response) {
	data = d
	log.Println("Starting listener on :4321")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":4321", nil))
}
