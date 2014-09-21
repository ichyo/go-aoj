package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ichyo/go-aoj"
)

func Usage() {
	usage := `
Usage:

	aoj command [argument]

The commands are:

	user		show user infomation
	problem		show problem information
	submit		submit solution
`
	fmt.Println(usage)
	os.Exit(1)
}

func UsageUser() {
	usage := "Usage: aoj user user_id"
	fmt.Println(usage)
	os.Exit(1)
}

func UsageProblem() {
	usage := "Usage: aoj problem problem_id"
	fmt.Println(usage)
	os.Exit(1)
}

func User() {
	if len(os.Args) <= 2 || os.Args[2] == "" {
		UsageUser()
	}

	id := os.Args[2]
	user, err := aoj.GetUser(id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID:", user.ID)
	fmt.Println("Name:", user.Name)
	fmt.Println("Affiliation:", user.Affiliation)
	fmt.Println("RegisterDate:", user.RegisterDate)
	fmt.Println("LastSubmitDate:", user.LastSubmitDate)
}

func Problem() {
	if len(os.Args) <= 2 || os.Args[2] == "" {
		UsageProblem()
	}

	id := os.Args[2]
	p, err := aoj.GetProblem(id, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID:", p.ID)
	fmt.Println("Name:", p.Name)
}

func Submit() {
}

func main() {
	if len(os.Args) <= 1 {
		Usage()
	}
	switch os.Args[1] {
	case "user":
		User()
	case "problem":
		Problem()
	case "submit":
		Submit()
	}
}
