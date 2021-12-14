package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func error(msg string) string {
	fmt.Fprintf(os.Stderr, "An error has occured: %s", msg)
	return "err"
}

func main() {
	fmt.Println("Welcome to the CLI version of templo7k.ninja!\n\nChoose the article you will read and we will automatically save it's pdf to your machine.")
	fmt.Println("1 - Programming\n2 - Cyber Security\n3 - Mathematics/Physics\n4 - Others\nInput option below (1-4) ")

	var topic uint8
	var opt int

	fmt.Scanln(&topic)

	switch topic {
	case 1:
		fmt.Println("\nProgramming\n1 - Bit Shift Operators")
		fmt.Scanln(&opt)
		defer programming(opt)
	case 2:
		fmt.Println("No paper yet")
	case 3:
		fmt.Println("No paper yet")
	case 4:
		fmt.Println("No paper yet")
	default:
		error("An error occured, input only a number between 1-4")
	}

}

func programming(prog int) int {
	if prog == 1 {
		out, err := os.Create("bit_shift.pdf")
		paper, err := http.Get("https://github.com/templo7k/templo7k-cli/raw/main/papers/bit_shift.pdf")

		defer out.Close()

		defer paper.Body.Close()

		_, err = io.Copy(out, paper.Body)

		fmt.Println("PDF Downloaded")

		if err != nil {
			fmt.Println(err)
		}
		if paper.StatusCode != http.StatusOK {
			error("Bad status code")
		}
	} else {
		error("Paper does not exist (input only numbers)")
	}
	return 0
}
