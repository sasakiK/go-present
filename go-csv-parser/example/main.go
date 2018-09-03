package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kardianos/osext"
)

type Person struct {
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Address   *Address `json:address,omitempty`
}

type Address struct {
	City  string `json:"city"`
	State string `json:state`
}

func dirPrinter() {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(folderPath)
}

// main all start
func main() {
	dirPrinter()
	// read csv start
	csvFile, err := os.Open("data/people.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	// read csv end
	// insert into struct start
	var people []Person
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break // 空白行に達したらbreak
		} else if err != nil {
			log.Fatal(err)
		}
		people = append(people, Person{
			Firstname: line[0],
			Lastname:  line[1],
			Address: &Address{
				City:  line[2],
				State: line[3],
			},
		})
	}
	peopleJson, _ := json.Marshal(people)
	// insert into struct end
	fmt.Println(string(peopleJson))
}

// main all end
