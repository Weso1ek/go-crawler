package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	
	if len(cfg.pages) > cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	// skip other websites
	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
		return
	}

	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	nextURLs, err := getURLsFromHTMLBootdev(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v", err)
		return
	}

	for _, nextURL := range nextURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}

//func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
//	currentURL, err := url.Parse(rawCurrentURL)
//	if err != nil {
//		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
//		return
//	}
//	baseURL, err := url.Parse(rawBaseURL)
//	if err != nil {
//		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
//		return
//	}
//
//	// skip other websites
//	if currentURL.Hostname() != baseURL.Hostname() {
//		return
//	}
//
//	normalizedURL, err := normalizeURLBootdev(rawCurrentURL)
//	if err != nil {
//		fmt.Printf("Error - normalizedURL: %v", err)
//		return
//	}
//
//	// increment if visited
//	if _, visited := pages[normalizedURL]; visited {
//		pages[normalizedURL]++
//		return
//	}
//
//	// mark as visited
//	pages[normalizedURL] = 1
//
//	fmt.Printf("crawling %s\n", rawCurrentURL)
//
//	htmlBody, err := getHTMLBootdev(rawCurrentURL)
//	if err != nil {
//		fmt.Printf("Error - getHTML: %v", err)
//		return
//	}
//
//	nextURLs, err := getURLsFromHTMLBootdev(htmlBody, rawBaseURL)
//	if err != nil {
//		fmt.Printf("Error - getURLsFromHTML: %v", err)
//		return
//	}
//
//	for _, nextURL := range nextURLs {
//		crawlPage(rawBaseURL, nextURL, pages)
//	}
//}
