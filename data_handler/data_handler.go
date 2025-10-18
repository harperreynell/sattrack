package data_handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	View   View     `json:"view"`
	Member []Member `json:"member"`
}

type Member struct {
	Satid int64  `json:"satelliteId"`
	Name  string `json:"name"`
	Date  string `json:"date"`
	Line1 string `json:"line1"`
	Line2 string `json:"line2"`
}

type View struct {
	Last string `json:"last"`
}

func checkForUpdate() bool {
	date := time.Now()
	lastUpdate, _ := time.Parse(time.RFC3339, readData()[0].Member[0].Date)
	diff := date.Sub(lastUpdate)

	return int(diff.Hours()/24) > 5
}

func request(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

func writeData(responseData []byte, filename string) {
	if _, err := os.Stat("data/" + filename); errors.Is(err, os.ErrNotExist) {
		fo, err := os.Create("data/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		defer fo.Close()

		fo.Write(responseData)
	} else {
		fo, err := os.Open("data/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		defer fo.Close()
		fo.Write(responseData)
	}
}

func updateData() []Response {
	APIURL := "https://tle.ivanstanojevic.me/api/tle/?page-size=100&page="
	responseData := request(APIURL + "1")
	writeData(responseData, "1")

	var responseObject Response
	var data []Response

	json.Unmarshal(responseData, &responseObject)
	data = append(data, responseObject)
	lastPage := strings.Split(responseObject.View.Last, "&")
	lastPageNumber := strings.Split(lastPage[len(lastPage)-1], "=")
	num, err := strconv.Atoi(lastPageNumber[len(lastPageNumber)-1])
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= num; i++ {
		url := APIURL + strconv.Itoa(i)
		responseData = request(url)
		writeData(responseData, strconv.Itoa(i))
		json.Unmarshal(responseData, &responseObject)
		data = append(data, responseObject)
	}

	return data
}

func readData() []Response {
	var responseObject Response
	var data []Response
	b, err := ioutil.ReadFile("data/1")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(b, &responseObject)
	data = append(data, responseObject)
	return data
}

func GetByID(id int, data []Response) Member {
	var ret Member
	for _, response := range data {
		for _, sat := range response.Member {
			if id == int(sat.Satid) {
				return sat
			}
		}
	}

	return ret
}

func GetData() []Response {
	var data []Response
	if checkForUpdate() {
		data = updateData()
	} else {
		data = readData()
	}

	return data
}
