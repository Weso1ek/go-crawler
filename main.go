package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	fmt.Println("starting crawl of: ", os.Args[1])

	//respBody, err := getHTML(os.Args[1])
	_, err := getHTML(os.Args[1])

	if err != nil {
		os.Exit(1)
	}

	//fmt.Println(respBody)

	pages := make(map[string]int)

	crawlPage(os.Args[1], os.Args[1], pages)
}
