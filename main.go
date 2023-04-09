package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Data struct {
	Status      `json:"status"`
	WaterStatus string `json:"water_status"`
	WindStatus  string `json:"wind_status"`
}

func updateData() {
	for {

		var data = Data{Status: Status{}}
		waterMax := 100

		data.Status.Water = rand.Intn(waterMax)

		data.Status.Wind = rand.Intn(waterMax)

		if data.Status.Water < 5 {
			data.WaterStatus = "aman"
		} else if data.Status.Water > 8 {
			data.WaterStatus = "bahaya"
		} else {
			data.WaterStatus = "siaga"
		}

		if data.Status.Wind < 6 {
			data.WindStatus = "aman"
		} else if data.Status.Wind > 15 {
			data.WindStatus = "bahaya"
		} else {
			data.WindStatus = "siaga"
		}

		b, err := json.MarshalIndent(data, "", " ")

		if err != nil {
			log.Fatalln("error while marshalling json data  =>", err.Error())
		}

		err = ioutil.WriteFile("data.json", b, 0644)

		if err != nil {
			log.Fatalln("error while writing value to data.json file  =>", err.Error())
		}
		fmt.Println("menggungu 15 detik")
		time.Sleep(time.Second * 15)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	go updateData()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := template.ParseFiles("index.html")

		var data = Data{Status: Status{}}

		b, err := ioutil.ReadFile("data.json")

		if err != nil {
			fmt.Fprint(w, "error braderku")
			return
		}

		err = json.Unmarshal(b, &data)

		err = tpl.ExecuteTemplate(w, "index.html", data)

	})

	http.ListenAndServe(":8080", nil)
}
