package main

import (
	"os"
	"testing"
)

func TestFileLoad(t *testing.T) {
	d := Datas{}
	current, _ := os.Getwd()
	name := []string{(current + "/tmp/test.json")}
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
