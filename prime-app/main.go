package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func prompt() {
	fmt.Print("-> ")
}

func main() {
	// print a welcome msg
	prompt()

	// Create a channel to indicate when the user wants to quit
	doneChan := make(chan bool)

	// start a go routine to read user input  and run program
	go readUserInput(os.Stdin, doneChan)

	// block until done chan gets value
	<-doneChan

	// close the chan
	close(doneChan)

	// say good bye
	fmt.Println("Bid adieu!")
}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user Input
	scanner.Scan()

	// check to see if the user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert user input to string
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter an integer!", false
	}

	_, msg := isPrime(num)

	return msg, false
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is non-prime, by definition!", n)
	}

	if n < 0 {
		return false, "Negative numbers are non-prime, by definition!"
	}

	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is non-prime, since it is divisible by %d!", n, i)
		}
	}

	return true, fmt.Sprintf("%d is prime!", n)
}
