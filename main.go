package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {

	// User picks a document to format
	fmt.Println("Hello, which document do you need to compile?")
	documents, _ := ioutil.ReadDir("./documents")
	for _, document := range documents {
		fmt.Println("\n" + document.Name())
	}

	var docChoice string
	fmt.Printf("::> ")
	fmt.Scanln(&docChoice)

	_, err := os.Lstat("documents/" + docChoice)
	if err != nil {
		panic("Can't find that Document! :O")
	}

	fmt.Println("You have picked " + docChoice)

	templates, _ := ioutil.ReadDir("documents/" + docChoice)
	finalizer := false

	for !finalizer {
		chosenParagraph, _ := chooseParagraph(templates)
		if chosenParagraph == "q" {
			finalizer = true
			fmt.Println("Done, compiling document")
		}
	}

}

func chooseParagraph(templates []os.FileInfo) (chosenTemplate string, err error) {

	type TemplatePair struct {
		template string
		idx      int
	}

	// User picks the next paragraph
	templateList := []TemplatePair{}
	var paraChoiceNo string

	fmt.Println("Pick the next paragraph:")
	for idx, template := range templates {
		templateList = append(templateList, TemplatePair{template.Name(), idx})
		fmt.Println("[" + strconv.Itoa(idx) + "]: " + template.Name())
	}
	fmt.Println("[q]: finished")

	fmt.Scanln(&paraChoiceNo)
	if paraChoiceNo == "q" {
		return "q", nil
	}
	for _, list := range templateList {
		if s, _ := strconv.Atoi(paraChoiceNo); s == list.idx {
			return list.template, nil
		}
	}
	return "", errors.New("Error indexing paragraphs")
}
