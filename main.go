package main

import (
	"log"
	"satinfo/data_handler"
	"satinfo/web"
)

func main() {
	data := data_handler.GetData()

	var ids []int64
	for _, d := range data {
		for _, sat := range d.Member {
			ids = append(ids, sat.Satid)
		}
	}
	log.Println(len(ids))
	web.Listen(data)

}
