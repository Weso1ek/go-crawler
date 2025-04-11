package main

import (
	"fmt"
	"net/url"
)

func normalizeURL(inputURL string) (string, error) {
	fmt.Println("Hello, World!")

	u, err := url.Parse(inputURL)
	if err != nil {
		panic(err)
	}
	
	return u.Host + u.Path, nil
}
