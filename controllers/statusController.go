package controllers

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Time string
	Wind int
	Water int
	WindStatus string
	WaterStatus string
}

func CreateStatus(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./views/index.html")

	if err != nil {
		panic(err)
	}
	jsonFile, err := os.Open("./status.json")

	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(15 * time.Second)

		ioJson, _ := ioutil.ReadAll(jsonFile)
		var statusValue Status

		json.Unmarshal(ioJson, &statusValue)

		h, m, s := time.Now().Clock()
		statusValue.Time = fmt.Sprintf("%d:%d:%d", h, m, s)
		statusValue.Wind = rand.Intn(100)
		statusValue.Water = rand.Intn(100)

		if statusValue.Water < 5 {
			statusValue.WaterStatus = "Aman"
		} else if statusValue.Water >= 6 && statusValue.Water <= 8 {
			statusValue.WaterStatus = "Siaga"
		} else if statusValue.Water > 8 {
			statusValue.WaterStatus = "Bahaya"
		}

		if statusValue.Wind < 6 {
			statusValue.WindStatus = "Aman"
		} else if statusValue.Wind >= 7 && statusValue.Wind <= 15 {
			statusValue.WindStatus = "Siaga"
		} else if statusValue.Wind > 15 {
			statusValue.WindStatus = "Bahaya"
		}

		errx := tmpl.Execute(w, statusValue)
		if errx != nil {
			panic(errx)
		}
	}
}