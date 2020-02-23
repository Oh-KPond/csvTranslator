package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Book struct {
	id           int
	mergeResults *mergeResults
}

type mergeResults struct {
	block      string
	chart      string
	hat        string
	sell       string
	success    string
	company    string
	subtract   string
	event      string
	particular string
	deal       string
	swim       string
	term       string
	opposite   string
	wife       string
	shoe       string
	shoulder   string
	spread     string
	arrange    string
	camp       string
	invent     string
	cotton     string
	born       string
	determine  string
	quart      string
	nine       string
	truck      string
	noise      string
	level      string
	chance     string
	gather     string
	shop       string
	stretch    string
	throw      string
	shine      string
	property   string
	column     string
	molecule   string
	wrong      string
	gray       string
	repeat     string
	require    string
	broad      string
	prepare    string
	salt       string
	nose       string
	plural     string
	anger      string
	claim      string
	continent  string
}

func main() {
	fmt.Println("Let's do this!")
}

// func merge(file string) {

// 	csvFile, _ := os.Open("people.csv")
// }

func getMergeKey() map[string]map[string]string {
	var mergeKey map[string]map[string]string

	file, _ := ioutil.ReadFile("dictionary/languageDict.json")
	err := json.Unmarshal([]byte(file), &mergeKey)
	if err != nil {
		log.Println(err)
	}
	return mergeKey
}

func getLanguage(fileName string) string {
	return strings.Fields(fileName)[0]
}
