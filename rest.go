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
	current, _ := os.Getwd()                          //Get current directory path
	list, _ := filepath.Glob(current + "/tmp/*.json") //Read all .json files
	return d.loadJson(list)
}
func (d Datas) get(w http.ResponseWriter, r *http.Request) {
	bodybytes, err := ioutil.ReadAll(r.Body) //To read wanted key
	if err != nil {                          //To check the reading if it is successful or not
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var data Data
	err = json.Unmarshal(bodybytes, &data) //Convert []byte to json
	if err != nil {                        //To check whether the key was sent
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("No Key Found"))
		return
	}
	if _, ok := d[data.Key]; !ok { //To check the wanted key exist
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Key Not Found\n"))
		return
	}
	//fmt.Println(data) //To see wanted key
	fmt.Println("EndPoint Hit: GET EndPoint")
	json.NewEncoder(w).Encode(d[data.Key])
}
func (d Datas) put(w http.ResponseWriter, r *http.Request) {
	bodybytes, err := ioutil.ReadAll(r.Body) //To read data
	if err != nil {                          //To check the reading if it is successful or not
		w.WriteHeader(http.StatusInternalServerError) //500
		w.Write([]byte(err.Error()))
		return
	}
	var data Data
	err = json.Unmarshal(bodybytes, &data) //Convert []byte to json
	if err != nil {                        //To check whether the key was sent
		w.WriteHeader(http.StatusBadGateway) //502
		w.Write([]byte(err.Error()))
		return
	}
	/* if _, ok := d[data.Key]; !ok {//To check the wanted key exist
		w.WriteHeader(http.StatusNotFound)//404
		w.Write([]byte("Key Not Found\n"))
		return
	} */
	//fmt.Println(data) //To see wanted key
	d[data.Key] = data.Value //Update or add new key/value
	fmt.Println("EndPoint Hit: SET EndPoint")
	w.WriteHeader(http.StatusNoContent) //204
}
func (d *Datas) delete(w http.ResponseWriter, r *http.Request) {
	*d = Datas{} //Empty Map
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
	file, _ := json.MarshalIndent(d, "", " ")                      //Json to byte with indent
	current, _ := os.Getwd()                                       //Get current directory path
	ioutil.WriteFile(current+"/tmp/"+fileName+".json", file, 0666) //Write in File
}
func (d Datas) loadJson(fileName []string) *Datas {
	if fileName == nil { //No Json file exist
		return &Datas{ //Create New
			"Test_Key0": "Test_Value0",
			"Test_Key1": "Test_Value1",
		}
	}
	//Json Files exist
	lastFile := last(fileName)                //Get Last Created File
	byteSlice, _ := ioutil.ReadFile(lastFile) //Read file
	saved := Datas{}
	json.Unmarshal(byteSlice, &saved) //Save in
	return &saved
}
func handleRequests() {
	datas := newDatas()
	http.HandleFunc("/", datas.homepage)
	http.HandleFunc("/datas", datas.datas)
	go func() { //Go routine
		if _, err := os.Stat("tmp"); os.IsNotExist(err) { //Checks if 'tmp' folder exist
			os.Mkdir("tmp", 0755) //if not creates
		}
		for { //Every 10 second,saves into a file
			time.Sleep(time.Second * sec)
			tm := time.Now().Unix() //Timestamp
			datas.saveJson(strconv.FormatInt(tm, 10))
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
