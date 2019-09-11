package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fileName := os.Args[1]
	jsonFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	newFile := strings.Split(fileName, ".")

	abi, _ := json.Marshal(result["abi"])
	err = ioutil.WriteFile("."+newFile[1]+".abi", abi, 0644)

	fmt.Println(newFile[1] + ".abi file created")
}
