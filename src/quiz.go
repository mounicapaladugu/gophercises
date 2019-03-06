package main

import (
	"sort"
	"os"
	"fmt"
	"log"
	"encoding/csv"
	"bufio"
	"time"
	"io"
	"sync"
	"github.com/urfave/cli"
)

type Quiz struct {
	Question string
	Answer string
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main(){
		//filepath string
		var filepath string
		var totaltime int
		var random bool
		var points int
		//this it to fire off a goroutine that keeps track of all goroutines and close gracefully after they end. 
		var wg sync.WaitGroup

		//cli to parse filepath from user. Has a default location to load from incase of no input
		app := cli.NewApp()
		app.Name = "Quiz"
		app.Usage = "Fun Quiz: A Mind Bender"
	
		//global level flags 
		app.Flags = []cli.Flag{
			cli.StringFlag{
				Name:        "filepath",
				Value:       "../data/problems.csv",
				Usage:       "filepath to load the csv file",
				Destination: &filepath,
			},
			cli.IntFlag{
				Name: "totaltime",
				Value: 5,
				Usage: "Set total time of Quiz",
				Destination: &totaltime,
			},
			cli.BoolFlag{
				Name: "random",
				Usage: "Ask questions in random order",
				Destination: &random,
			},
		}

		app.Action = func(c *cli.Context) error {
			fmt.Println("Welcome to the Quiz")
			return nil
		}

		sort.Sort(cli.FlagsByName(app.Flags))
		err := app.Run(os.Args)
		if err != nil {
			log.Fatal(err)
		}

		Quiz, err := parseCsv(filepath)
		checkError(err)
		
		//wait for Enter to start quiz
		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')

		//creating a channel to communicate 
		communicate := make(chan string)
		//TODO:implement context 
		/*
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(totaltime)*time.Second)
		defer cancel() //this is to make sure all paths cancel the context to avoid context leak 
		*/

		wg.Add(1)
		timeout := time.After(time.Duration(totaltime)*time.Minute)
		//starting go routines for quiz and userinput 
		go func(){
		wrap:
			for i :=0;i < len(Quiz); i++ {
				go userInput(os.Stdout, os.Stdin, Quiz[i].Question, communicate)
				select {
				case <-timeout:
					fmt.Fprintf(os.Stderr, "Time's up!")
					break wrap
				
				case input, ok := <- communicate:
					if !ok {
						break wrap
					}
					if Quiz[i].Answer == input {
						points++
					}

				}
			}
			wg.Done()
		}()
		wg.Wait()
		close(communicate)

		fmt.Fprintf(os.Stdout, "%d correctly answered out of %d \n", points, len(Quiz))
		
}

func userInput(w io.Writer, r io.Reader, question string, sendTo chan string){
	reader := bufio.NewReader(r)
	fmt.Fprintln(w, "Question :" + question)
	fmt.Fprint(w, "Answer: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		//Incase of error, close channel
		close(sendTo)
		if err == io.EOF {
			return
		}
		log.Fatalln(err)
	}
	sendTo <- input
}

func parseCsv(filepath string)([]Quiz, error){
	//load the csv file
	f, err := os.Open(filepath)
	checkError(err)
	defer f.Close()

	//create a new reader
	r := csv.NewReader(bufio.NewReader(f))
	output := []Quiz{}

	//read the contents of csv and populate json
	for {
		line, error := r.Read()
		if error == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Error while parsing csv, %v", err)
		}
		if len(line) != 2 {
			return nil, fmt.Errorf("Unexpected number of fields for line: %v", line)
		}

		//answer, _ := strconv.Atoi(line[1])

		output = append(output, Quiz{
			Question: line[0],
			Answer: line[1],
		})
	}
	return output, nil
}
