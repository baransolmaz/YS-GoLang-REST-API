package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestFileLoad(t *testing.T) {
	d := Datas{}
	current, _ := os.Getwd()
	name := []string{(current + "/tmp/0.json")}
	read := *d.loadJson(name)
	if len(read) != 3 {
		t.Errorf("Expected: 3, Size: %d", len(read))
	}
	if read["Key"] != "Value" {
		t.Errorf("Expected: Value, Recieved: %s", read["Key"])
	}
	if read["Test_Key0"] != "Test_Value0" {
		t.Errorf("Expected: Test_Value0, Recieved: %s", read["Test_Key0"])
	}
}
func TestFileSave(t *testing.T) {
	current, _ := os.Getwd()
	os.Remove(current + "/tmp/testsave.json")
	d := Datas{
		"Testing": "testing",
		"test2":   "test2",
	}
	d.saveJson("testsave")
	if _, err := os.Stat(current + "/tmp/testsave.json"); os.IsNotExist(err) { //Checks if 'testsave.json' file exist
		t.Errorf("File Not Created")
	}
	name := []string{(current + "/tmp/testsave.json")}
	read := *d.loadJson(name)
	if len(read) != 2 {
		t.Errorf("Expected: 2, Size: %d", len(read))
	}
	if read["Testing"] != "testing" {
		t.Errorf("Expected: testing, Recieved: %s", read["Testing"])
	}
	if read["test2"] != "test2" {
		t.Errorf("Expected: test2, Recieved: %s", read["test2"])
	}
	os.Remove(current + "/tmp/testsave.json")
}
func TestGetRequest(t *testing.T) {
	bodyReader := strings.NewReader("{\"Key\":\"Test_Key0\"}")
	req := httptest.NewRequest(http.MethodGet, "/requests", bodyReader)
	w := httptest.NewRecorder()
	d := Datas{}
	current, _ := os.Getwd()
	name := []string{(current + "/tmp/0.json")}
	read := *d.loadJson(name)
	read.requests(w, req)
	res := w.Result()

	defer res.Body.Close()
	if res.StatusCode == http.StatusInternalServerError {
		t.Errorf("Status Internal Server Error")
	} else if res.StatusCode == http.StatusBadGateway {
		t.Errorf("Status Bad Gateway")
	} else if res.StatusCode == http.StatusNotFound {
		t.Errorf("Status Not Found")
	}
}
func TestPutRequest(t *testing.T) {
	bodyReader := strings.NewReader("{\"Key\":\"Test_Key0\",\"Value\":\"Test_Put\"}")
	req := httptest.NewRequest(http.MethodPut, "/requests", bodyReader)
	w := httptest.NewRecorder()
	d := Datas{}
	current, _ := os.Getwd()
	name := []string{(current + "/tmp/0.json")}
	read := *d.loadJson(name)
	read.requests(w, req)
	res := w.Result()

	defer res.Body.Close()
	if res.StatusCode == http.StatusInternalServerError {
		t.Errorf("Status Internal Server Error")
	} else if res.StatusCode == http.StatusBadGateway {
		t.Errorf("Status Bad Gateway")
	} else if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status No Content")
	}
}
func TestViewRequest(t *testing.T) {
	req := httptest.NewRequest("VIEW", "/requests", nil)
	w := httptest.NewRecorder()
	d := Datas{}
	current, _ := os.Getwd()
	name := []string{(current + "/tmp/0.json")}
	read := *d.loadJson(name)
	read.requests(w, req)
	res := w.Result()

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status Not Ok")
	}
}
func TestDeleteRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/requests", nil)
	w := httptest.NewRecorder()
	d := Datas{}
	current, _ := os.Getwd()
	name := []string{(current + "/tmp/0.json")}
	read := *d.loadJson(name)
	read.requests(w, req)
	res := w.Result()

	defer res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status No Content")
	}
	if len(read) != 0 {
		t.Errorf("Expected: 0, Size: %d", len(read))
	}
	if read["Key"] != "" {
		t.Errorf("Expected: \"\", Recieved: %s", read["Key"])
	}
}
