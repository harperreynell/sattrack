package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"satinfo/data_handler"
)

func checkForUpdate() bool {
	date := time.Now()
	lastUpdate, _ := time.Parse(time.RFC3339, data_handler.ReadData()[0].Member[0].Date)
	diff := date.Sub(lastUpdate)

	return int(diff.Hours()/24) > 5
}

func main() {
	var data []data_handler.Response
	if checkForUpdate() {
		log.Println("Updating data...")
		data = data_handler.UpdateData()
	} else {
		log.Println("Reading data...")
		data = data_handler.ReadData()
	}

	for _, response := range data {
		fmt.Println(response)
		for _, sat := range response.Member {
			fmt.Println(sat.Name)
			fmt.Println("\tID    : " + strconv.Itoa(int(sat.Satid)))
			fmt.Println("\tLine1 : " + sat.Line1)
			fmt.Println("\tLine2 : " + sat.Line2)
		}
	}
}

