package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Key   string "json:\"Key\""
	Value string "json:\"Value\""
}
type Datas map[string]Data

func allDatas(w http.ResponseWriter, r *http.Request) {
	datas := Datas{
		"Test_Key0": Data{Key: "Test_Key0", Value: "Test_Value0"},
		"Test_Key1": Data{Key: "Test_Key1", Value: "Test_Value1"},
	}
	fmt.Println("EndPoint Hit: All Datas EndPoint")
	json.NewEncoder(w).Encode(datas)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage EndPoint Hit")
}
func handleRequests() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/datas", allDatas)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
