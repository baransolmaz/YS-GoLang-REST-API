package main

import (
	"os"
	"testing"
)

func TestFileRead(t *testing.T) {
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
