package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type info struct {
	phone string
}

func main() {
	dir := map[string]info{
		"ziad":    {"075268618"},
		"sabrina": {"0685268848"},
	}

	action := flag.String("action", "", "L'action à effectuer (ajouter, supprimer, etc.)")
	name := flag.String("name", "", "Nom de la personne")
	phone := flag.String("phone", "", "Numero de la personne")
	//age := flag.String("age", "", "Age de la personne")

	flag.Parse()

	switch *action {
	case "list":
		listDir(dir)
	case "search":
		searchDir(dir, *name)
	case "add":
		addDir(dir, *name, *phone)
	case "update":
		updateDir(dir, *name, *phone)
	case "delete":
		deleteDir(dir, *name)
	default:
		fmt.Println("Unknown action")
	}
}

func paramVerif(name string, phone string, ageStr string) {
	if name == "" {
		fmt.Errorf("Name must not be empty")
	}
	if phone == "" {
		fmt.Errorf("Number must not be empty")
	} else if len(phone) == 10 {
		fmt.Errorf("Number must have at 10 digits")
	}
	_, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Errorf("Age must be a number")
	}
}

func listDir(dir map[string]info) {
	dirLen := len(dir)
	if dirLen == 0 {
		println("Empty directory")
		return
	} else {
		suffix := ""
		if dirLen > 1 {
			suffix = "s"
		}
		fmt.Printf("%d person%s found\n", dirLen, suffix)
	}
	for name, pinfo := range dir {
		printPerson(name, pinfo)
	}
}

func searchDir(dir map[string]info, name string) {
	personInfo, ok := dir[strings.ToLower(name)]
	if !ok {
		fmt.Println("Person not found in directory")
		return
	}
	printPerson(name, personInfo)
}

func printPerson(name string, personInfo info) {
	fmt.Printf("Name: %s \nPhone number: %s\n", name, personInfo.phone)
	fmt.Println("-----------------------")
}

func addDir(dir map[string]info, name string, phone string) {
	_, ok := dir[strings.ToLower(name)]
	if ok {
		fmt.Printf("%s already exist in the directory", name)
		return
	}

	dir[name] = info{phone: phone}

	fmt.Println("List updated : ")
	listDir(dir)
}

func deleteDir(dir map[string]info, name string) {
	_, ok := dir[strings.ToLower(name)]
	if !ok {
		fmt.Println("Person not found in directory")
		return
	}
	delete(dir, strings.ToLower(name))
}

func updateDir(dir map[string]info, name string, phone string) {
	_, ok := dir[strings.ToLower(name)]
	if !ok {
		fmt.Println("Person not found in directory")
		return
	}
	dir[strings.ToLower(name)] = info{
		phone: phone,
	}

	fmt.Println("Info updated : ")
	searchDir(dir, name)
}
