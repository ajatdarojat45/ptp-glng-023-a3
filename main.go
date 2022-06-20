package main

import (
	"fmt"
	"net/http"
	"html/template"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	// "time"
)

var PORT = ":8000"

type DataFile struct {
	Status struct {
		Water int `json:"water"`
		Wind int `json:"wind"`
	}
}

func main(){
	http.HandleFunc("/", getWaterAndWind)
	fmt.Println("server is running", PORT)
	http.ListenAndServe(PORT, nil)
}

func getWaterAndWind(w http.ResponseWriter, r *http.Request)  {
	tmplte, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := DataFile{}
	// for {
		randomWater := rand.Intn(15)
		randomWind := rand.Intn(20)
	
		data.Status.Water = randomWater
		data.Status.Wind = randomWind
	
		newJson, _ := json.Marshal(data)
		ioutil.WriteFile("db.json", newJson, 0644)
	
	// 	time.Sleep(time.Second * 15)
	// }
	tmplte.Execute(w, data)
	return
}