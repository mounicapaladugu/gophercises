package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

type Line struct {
	Question string `json:"question"`
	Answer   int    `json:"answer"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadCsvFile(filepath string) ([]Line, error) {
	//load the csv file
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()

	//create a new reader
	r := csv.NewReader(bufio.NewReader(f))

	var lines []Line

	//read the contents of csv and populate json
	for {
		line, error := r.Read()
		if error == io.EOF {
			break
		}
		check(error)
		answer, _ := strconv.Atoi(line[1])

		lines = append(lines, Line{
			Question: line[0],
			Answer:   answer,
		})
	}
	return lines, nil

	//printing for debugging purpose
	/*
		linesJSON, _ := json.MarshalIndent(lines, "", "    ")
		fmt.Println(string(linesJSON))
	*/
}

func readInput(input chan<- int) {
	for {
		var u int
		_, err := fmt.Scanf("%d \n", &u)
		check(err)

		input <- u
	}
}

func checkAnswer() {

}

func testScore() {

}

func askQuestions() {

}

func main() {
	//filepath string
	var filepath string

	//cli to parse filepath from user. Has a default location to load from incase of no input
	app := cli.NewApp()
	app.Name = "Quiz"
	app.Usage = "Fun Quiz: A Mind Bender"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "filepath",
			Value:       "data/problems.csv",
			Usage:       "filepath to load the csv file",
			Destination: &filepath,
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("Welcome to the Quiz")
		Line, _ := ReadCsvFile(filepath)

		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		var points int
		userInput := make(chan int)

		go readInput(userInput)

		
		for i := 0; i < len(Line); i++ {
				
			fmt.Println(Line[i].Question)
			fmt.Print("Enter Answer \n")

			select {
			case userAnswer := <-userInput:
				if userAnswer == Line[i].Answer {
					fmt.Println("Correct Answer:", userAnswer)
					points++
				} else {
					fmt.Println("Wrong Answer!")
				}
			}
		}
		

		fmt.Printf("Total correct answers %d", points)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
