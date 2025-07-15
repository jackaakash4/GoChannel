package main

import (
    "fmt"
    "flag"
    "os"
    "encoding/csv"
    "strings"
)

type problem struct{
    question string
    answer string
}

func exit(msg string){
    fmt.Println(msg)
    os.Exit(1)
}

func parseLines(lines [][]string) []problem{
    ret := make([]problem, len(lines))
    for i, line := range lines{
        ret[i] = problem{
                    question: line[0],
                    answer: strings.TrimSpace(line[1]),
                }
            }
            return ret
        }


func main(){

    csvFilename := flag.String("csv", "problem.csv", "a csv file containg a quiz questions and answers")
    flag.Parse()

    file, err := os.Open(*csvFilename)
    if err != nil {
        exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFilename))
    }
    
    r := csv.NewReader(file)
    lines, err := r.ReadAll()
    
    if err != nil{
        exit("Failed to parse the csv file")

    }
    
    problems := parseLines(lines)
    
    correct := 0
    for i, prob := range problems{
        fmt.Printf("Problem %d: %s = \n", i, prob.question)
        var answer string
        fmt.Scanf("%s", &answer)
        if answer == prob.answer {
            correct++
        }
}
fmt.Printf("You've scored %d out of %d. \n", correct, len(problems))
}

