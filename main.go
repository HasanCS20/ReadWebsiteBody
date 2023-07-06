package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var url string

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the web URL: ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimRight(url, "\n")
	respone, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response Type %T\n: ", respone)
	defer respone.Body.Close()

	databyte, err := ioutil.ReadAll(respone.Body)

	if err != nil {
		panic(err)
	}
	content := string(databyte)

	if err != nil {
		panic(err)
	}

	fileName := "file.txt"
	fmt.Println(fileName)
	CreateFile(fileName, content)
}

func CreateFile(filename string, textContent string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	_, erro := file.WriteString(textContent)
	if erro != nil {
		panic(erro)
	}
}
