package main

import (
	"fmt"
	"net/http"
	"html/template"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

var PORT = ":8000"

type DataFile struct {
	Status struct {
		Water int `json:"water"`
		Wind int `json:"wind"`
	}
}

func main(){
	go generateWaterAndWind()
	http.HandleFunc("/", index)
	fmt.Println("server is running", PORT)
	http.ListenAndServe(PORT, nil)
}

func index(w http.ResponseWriter, r *http.Request)  {
	tmplte, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := DataFile{}

	dbyte, err := ioutil.ReadFile("db.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	json.Unmarshal(dbyte, &data)

	tmplte.Execute(w, data)
	return
}

func generateWaterAndWind(){
	for {
		data := DataFile{}
		randomWater := rand.Intn(15)
		randomWind := rand.Intn(20)
	
		data.Status.Water = randomWater
		data.Status.Wind = randomWind
	
		newJson, _ := json.Marshal(data)
		ioutil.WriteFile("db.json", newJson, 0644)
	
		time.Sleep(time.Second * 15)
	}
}