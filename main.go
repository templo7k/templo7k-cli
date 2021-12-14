package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var opt int

func download_pdf(post_name string) error {
	paper, err := http.Get("https://github.com/templo7k/templo7k-cli/raw/main/papers/" + post_name + ".pdf")

	if paper.StatusCode != http.StatusOK {
		fmt.Println("Bad status code")
	} else if paper.StatusCode == 200 {
		fmt.Println("PDF Downloaded")
	} else {
		fmt.Println("Paper does not exist (input only numbers)")
	}

	defer paper.Body.Close()

	out, err := os.Create(post_name + ".pdf")

	if err != nil {
		fmt.Println(err)
	}

	defer out.Close()

	_, err = io.Copy(out, paper.Body)
	return err
}

func verify_option(option int) string {
	valid_options := map[int]string{
		1: "AAAA",
		2: "BBBB",
	}

	return valid_options[option]
}

func choose_option() {
	fmt.Println("Welcome to the CLI version of templo7k.ninja!\n\nChoose the article you will read and we will automatically save it's pdf to your machine.")
	fmt.Println("1 - Programming\n2 - Cyber Security\n3 - Mathematics/Physics\n4 - Others\nInput option below (1-4) ")

	fmt.Scanln(&opt)
}

func main() {
	choose_option()
	chosed_option_verify := verify_option(opt)
	fmt.Println(chosed_option_verify)
}
