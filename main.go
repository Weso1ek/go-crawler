package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	fmt.Println("starting crawl of: ", os.Args[1])

	rawBaseURL := os.Args[1]
	maxConcurrency, _ := strconv.Atoi(os.Args[2])
	maxPages, _ := strconv.Atoi(os.Args[3])

	//const maxConcurrency = 3
	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	//for normalizedURL, count := range cfg.pages {
	//	fmt.Printf("%d - %s\n", count, normalizedURL)
	//}

	printReport(cfg.pages, rawBaseURL)

	////respBody, err := getHTML(os.Args[1])
	//_, err := getHTML(os.Args[1])
	//
	//if err != nil {
	//	os.Exit(1)
	//}
	//
	////fmt.Println(respBody)
	//
	//pages := make(map[string]int)
	//
	//crawlPage(os.Args[1], os.Args[1], pages)
}
