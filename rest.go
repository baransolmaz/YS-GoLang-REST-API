package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	Key   string "json:\"Key\""
	Value string "json:\"Value\""
}
type Datas map[string]string

func newDatas() *Datas {
	return &Datas{
		"Test_Key0": "Test_Value0",
		"Test_Key1": "Test_Value1",
	}
}
func (d Datas) get(w http.ResponseWriter, r *http.Request) {
	bodybytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var data Data

	err = json.Unmarshal(bodybytes, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("No Key Found"))
		return
	}
	if _, ok := d[data.Key]; !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Key Not Found\n"))
		return
	}
	fmt.Println(data)
	fmt.Println("EndPoint Hit: GET EndPoint")
	json.NewEncoder(w).Encode(d[data.Key])

}
func (d Datas) put(w http.ResponseWriter, r *http.Request) {
	bodybytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //500
		w.Write([]byte(err.Error()))
		return
	}
	var data Data
	err = json.Unmarshal(bodybytes, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway) //502
		w.Write([]byte(err.Error()))
		return
	}
	//Key yoksa eklemez
	/* if _, ok := d[data.Key]; !ok {
		w.WriteHeader(http.StatusNotFound)//404
		w.Write([]byte("Key Not Found\n"))
		return
	} */
	fmt.Println(data)
	d[data.Key] = data.Value
	fmt.Println("EndPoint Hit: SET EndPoint")
	w.WriteHeader(http.StatusNoContent) //204
}
func (d Datas) delete(w http.ResponseWriter, r *http.Request) {

}
func (d *Datas) datas(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		d.get(w, r)
		return
	case "PUT":
		d.put(w, r)
		return
	case "DELETE":
		d.delete(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed\n"))
		return
	}
}
func (d Datas) homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage EndPoint Hit")
}

func handleRequests() {
	datas := newDatas()
	http.HandleFunc("/", datas.homepage)
	http.HandleFunc("/datas", datas.datas)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
