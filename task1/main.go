package main

import (
	"archive/zip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

type JSON struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

type Character struct {
	XMLName   xml.Name `xml:"character"`
	FirstName string   `xml:"name_first"`
	LastName  string   `xml:"name_last"`
	Age       int      `xml:"age"`
}

func createFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Smth goes wrong: ", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Println(file.Name())
}

func removeFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("Smth goes wrong: ", err)
		os.Exit(1)
	}
}

func unzip() {
	app := "unzip"
	arg1 := "-l"
	arg2 := "archive.zip"

	shell := exec.Command(app, arg1, arg2)
	stdout, err := shell.Output()
	if err != nil {
		fmt.Println("Smth goes wrong: ", err)
		os.Exit(1)
		return
	}
	fmt.Println(string(stdout))
}

//--------------------------------------------------------------\\

func TaskOne() {
	app := "df"
	arg := "-h"

	shell := exec.Command(app, arg)
	stdout, err := shell.Output()
	if err != nil {
		fmt.Println("Smth goes wrong: ", err)
		os.Exit(1)
		return
	}

	fmt.Println(string(stdout))

}

func TaskTwo() {
	var password string
	fmt.Println("Enter your password: ")
	fmt.Fscan(os.Stdin, &password)

	fileName := password + ".txt"

	createFile(fileName)

	data := []byte(password)
	_ = ioutil.WriteFile(fileName, data, 0777)
	tmpl, _ := ioutil.ReadFile(fileName)
	fmt.Println(string(tmpl))
	fmt.Println("Do you want delete file? 0 - no; 1 - yes")
	var del string
	fmt.Fscan(os.Stdin, &del)
	if del == "1" {
		removeFile(fileName)
	}
	if del == "0" {
		fmt.Println("File is not deleted!")
	}

}

func TaskThree() {
	var info = [4]JSON{
		{"Dima", 19, "Mos"},
		{"Daniil", 19, "Mos"},
		{"Donya", 19, "Mos"},
		{"Zlata", 20, "Chel"},
	}

	createFile("JSON.json")

	tmpl, _ := json.Marshal(info)
	_ = ioutil.WriteFile("JSON.json", tmpl, 0777)
	tmpl, _ = ioutil.ReadFile("JSON.json")
	fmt.Println(string(tmpl))
	fmt.Println("Do you want delete file? 0 - no; 1 - yes")
	var del string
	fmt.Scanf("%s\n", &del)
	if del == "1" {
		removeFile("JSON.json")
	}
	if del == "0" {
		fmt.Println("File is not deleted!")
	}
}

func TaskFour() {
	createFile("XML.xml")
	var firstName, lastName string
	var age int
	fmt.Println("Enter your name: ")
	fmt.Fscan(os.Stdin, &firstName)
	fmt.Println("Enter your last name: ")
	fmt.Fscan(os.Stdin, &lastName)
	fmt.Println("Enter your age: ")
	fmt.Fscan(os.Stdin, &age)
	character := &Character{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

	file, _ := xml.MarshalIndent("XML.xml", "", "")

	_ = ioutil.WriteFile("XML.xml", file, 0644)
	tmpl, _ := ioutil.ReadFile("XML.xml")

	_ = xml.Unmarshal([]byte(tmpl), &character)
	fmt.Println("Name: ", character.FirstName, " Last name: ", character.LastName, " Age: ", character.Age)

	fmt.Println("Do you want delete file? 0 - no; 1 - yes")
	var del string
	fmt.Scanf("%s\n", &del)
	if del == "1" {
		removeFile("XML.xml")
	}
	if del == "0" {
		fmt.Println("File is not deleted!")
	}
}

func TaskFive() {
	fmt.Println("Making archive...")
	archive, err := os.Create("archive.zip")
	if err != nil {
		panic(err)
	}

	defer archive.Close()

	zipWriter := zip.NewWriter(archive)

	fmt.Println("Opening file...")
	file, err := os.Open("index.html")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fmt.Println("Adding file to archive...")
	w, err := zipWriter.Create("html/index.html")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w, file); err != nil {
		panic(err)
	}
	fmt.Println("Closing archive...")
	zipWriter.Close()
	unzip()
	fmt.Println("Do you want delete file? 0 - no; 1 - yes")
	var del string
	fmt.Scanf("%s\n", &del)
	if del == "1" {
		removeFile("archive.zip")
	}
	if del == "0" {
		fmt.Println("File is not deleted!")
	}
}

func main() {
	//fmt.Println("firstTask is starting...")
	//TaskOne()
	fmt.Println("secondTask is starting...")
	TaskTwo()
	fmt.Println("thirdTask is starting...")
	TaskThree()
	fmt.Println("forthTask is starting...")
	TaskFour()
	//fmt.Println("fifthTask is starting...")
	//TaskFive()
}
