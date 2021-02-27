package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func loadFromJSON(fileName string, key interface{}) error {
	stream, err := os.Open(fileName)
	if err != nil {
		return err
	}

	decodeJson := json.NewDecoder(stream)
	err = decodeJson.Decode(key)
	if err != nil {
		return err
	}
	stream.Close()
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}
	fileName := arguments[1]
	var myRecord Record
	err := loadFromJSON(fileName, &myRecord)
	if err == nil {
		fmt.Println(myRecord)
	} else {
		fmt.Println(err)
	}

}
