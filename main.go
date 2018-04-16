package main

import "flag"
import "fmt"
import "os"
import "encoding/csv"
import "strings"
import "time"

func main() {
  csvFileName := flag.String("csv","problems.csv","a csv file in the format of 'question,answer'") // a flag for getting a text to use as csv file name from the CLI
  timeLimit := flag.Int("limit",30,"the time limit for the quiz in second") // a flag for getting a number to use as the quiz time limit
  flag.Parse() // you should do parse every time

  file, err := os.Open(*csvFileName) // opening a file just like in c++
  if err != nil {
    exit(fmt.Sprintf("Faild to open CSV File: %s\n", *csvFileName)) // there is func that close the file and shows a message
  }
  r := csv.NewReader(file) // identify a read var as csv filr reader
  lines, err := r.ReadAll() // to read the whole csv lines
  if err != nil {
    exit("Faild to parse the provided CSV file.") // once again if there is an error close the file
  }
  problems := parseLinse(lines) // there is parseLinse func that takes an array of csv lines and break it into a Problem Object
  timer := time.NewTimer(time.Duration(*timeLimit) * time.Second) // the quiz timer setter

  correct := 0 // refrence to the number of correct answers
  for i, p := range problems {
    fmt.Printf("Problem #%d: %s =", i+1, p.q)
    answerCh := make(chan string) // making a chanel to check if the timer is out or not yet
    go func() {                  // its called a Closures function it the same as anonymous function but this one takes
      var answer string          // a variable from outside the functio which it in this case is  answerCh
      fmt.Scanf("%s\n", &answer)
      answerCh <- answer        // we r passing the asnwer that cames from the flag to a chanel so we can check it in line 37
    }()                       // we r saying that we want to call this func when we put this putting this parentheses
    select {
    case <-timer.C:         // if the time is out
      fmt.Printf("\n Time out and You scored %d of %d.\n", correct, len(problems))
      return               // print the score and break ..... we can't put break cuz it will break the select not the for loop
    case answer := <- answerCh:  // in case there is answer passed in the answerCh from the flag
      if answer == p.a {        // checking if it a right answer
        correct++
        fmt.Printf("Correct! \n")
      }else{
        fmt.Printf("Wrong! it is %s \n",p.a)
      }
    }
  }
  fmt.Printf("You scored %d of %d.\n", correct, len(problems))
}

func parseLinse(lines [][]string) []problem { // function that put the lins in array on problem objects
  ret := make([]problem, len(lines))
  for i, line := range lines {
    ret[i] = problem{
      q: line[0],
      a: strings.TrimSpace(line[1]),
    }
  }
  return ret
}

type problem struct { // a class so we can make objects from
  q string
  a string
}

func exit(msg string) { // function to close the csv file and print a msg
  fmt.Println(msg)
  os.Exit(1)
}
