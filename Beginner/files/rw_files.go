package main

import (
	"fmt"
	"io/ioutil"
	"os"
)
func main() {
	filename := "localfile.data"
	mydata := []byte("This is ramdom string in a file\n")
	err := ioutil.WriteFile(filename, mydata, 0777)
	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	f, err := os.OpenFile(filename, os.O_APPEND | os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	if _, err = f.WriteString("New data that wasn't there originally\n"); err != nil {
		panic(err)
	}
	defer f.Close()
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

}
