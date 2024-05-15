package scrapper

import (
	"fmt"
	"io"
	"net/http"
)

var url string
var plainHTML string

func GetUrl() string {

	fmt.Println("Enter the address of the website you want to parse")

	_, err := fmt.Scan(&url)
	if err != nil {
		//make check the validity of the url
		fmt.Print("\nPlease enter a valid URL\n")
	}
	return url
}

func GetPlainHTML() string {
	req, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Print(err)
		}
	}(req.Body)

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Print(err)
	}
	plainHTML = string(body)
	return plainHTML
}
