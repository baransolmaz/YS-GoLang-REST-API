package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const sec = 10

type Data struct {
	Key   string "json:\"Key\""
	Value string "json:\"Value\""
}
type Datas map[string]string

func newDatas() *Datas {
	var d Datas
	return d.loadJson("tmp")
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
func (d *Datas) delete(w http.ResponseWriter, r *http.Request) {
	*d = Datas{}
	fmt.Println("EndPoint Hit: DELETE EndPoint")
	w.WriteHeader(http.StatusNoContent) //204
}
func (d *Datas) datas(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": //Get Endpoint
		d.get(w, r)
		return
	case "PUT": //Put/Set Endpoint
		d.put(w, r)
		return
	case "DELETE": //Delete/Flush Endpoint
		d.delete(w, r)
		return
	case "VIEW": //View Endpoint for listing all datas
		d.view(w, r)
		return
	default: //Not Supported Endpoints
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed\n"))
		return
	}
}
func (d Datas) homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage EndPoint Hit")
}
func (d Datas) view(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EndPoint Hit: VIEW EndPoint")
	fmt.Fprintf(w, "VIEW EndPoint Hit\n\n")
	json.NewEncoder(w).Encode(d)
}
func (d Datas) saveJson(fileName string) {
	tm := time.Now().Unix()                                                                  //Timestamp
	file, _ := json.MarshalIndent(d, "", " ")                                                //Json to byte with indent
	current, _ := os.Getwd()                                                                 //Get current directory path
	ioutil.WriteFile(current+"/"+fileName+"/"+strconv.FormatInt(tm, 10)+".json", file, 0666) //Write in File
}

func (d Datas) loadJson(fileName string) *Datas {
	current, _ := os.Getwd()                                       //Get current directory path
	list, _ := filepath.Glob(current + "/" + fileName + "/*.json") //Read all .json files
	if list == nil {
		return &Datas{
			"Test_Key0": "Test_Value0",
			"Test_Key1": "Test_Value1",
		}
	}
	lastFile := last(list)
	byteSlice, _ := ioutil.ReadFile(lastFile)
	saved := Datas{}
	json.Unmarshal(byteSlice, &saved)
	return &saved

}
func handleRequests() {
	datas := newDatas()
	http.HandleFunc("/", datas.homepage)
	http.HandleFunc("/datas", datas.datas)
	go func() { //Go routine
		if _, err := os.Stat("tmp"); os.IsNotExist(err) { //Checks if 'tmp' folder exist
			os.Mkdir("tmp", 0755)
		}
		for { //Every 10 second,saves into a file
			time.Sleep(time.Second * sec)
			datas.saveJson("tmp")
		}
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))

}

//Returns Last Created File Name
func last(arr []string) string {
	if len(arr) == 2 {
		if strings.Compare(arr[0], arr[1]) == 1 {
			return arr[0]
		}
		return arr[1]
	} else if len(arr) == 1 {
		return arr[0]
	}
	var mid int = len(arr) / 2
	left := last(arr[:mid])
	right := last(arr[mid:])
	if strings.Compare(left, right) == 1 {
		return left
	}
	return right
}
