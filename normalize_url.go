package main

import (
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	u, err := url.Parse(inputURL)
	if err != nil {
		panic(err)
	}

	return u.Host + u.Path, nil
}

func getURLsFomHTML(htmlBody, rawBaseURL string) ([]string, error) {

	var urls []string
	var urlsParsed []string

	reader := strings.NewReader(htmlBody)

	parsed, err := html.Parse(reader)
	if err != nil {
		return urls, err
	}

	processLinks(parsed, &urlsParsed)

	for _, j := range urlsParsed {
		u, err := url.Parse(j)

		if err != nil {
			return urls, err
		}

		if u.Host == "" {
			urls = append(urls, rawBaseURL+j)
		}

		if u.Host != rawBaseURL && u.Host != "" {
			urls = append(urls, j)
		}
	}

	return urls, nil
}

func processLinks(n *html.Node, urls *[]string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				*urls = append(*urls, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		processLinks(c, urls)
	}
}
