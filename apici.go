package main

import (
	//	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const EXIT_STATUS = 2

type infoCollection struct {
	Info struct {
		PostmanId string `json:"_postman_id"`
		Schema    string `json:"schema"`
	} `json:"info"`
}

type postmanCollection struct {
	Info           infoCollection   `json:"info"`
	ItemCollection []itemCollection `json:"item"`
}

type requestSchema struct {
	Method string `json:"method"`
	Url    struct {
		Raw string `json:"raw"`
	} `json:"url"`
}

type itemCollection struct {
	FolderName   string `json:"name"`
	RequestItems []struct {
		RequestName   string        `json:"name"`
		RequestSchema requestSchema `json:"request"`
	} `json:"item"`
}

func readFile() (string, error) {
	/*
		if r := len(os.Args); r < 2 {
			return "", errors.New("Filename is required")
		}

		file, _ := os.Open(os.Args[1])

		if _, err := file.Stat(); err != nil {
			return "", err
		}

		return os.Args[1], nil
	*/
	return "a.json", nil
}

func CheckIfPostmanCollection(j []byte) error {
	dat := infoCollection{}

	if err := json.Unmarshal(j, &dat); err != nil {
		return err
	}

	if dat.Info.PostmanId != "" {
		return nil
	}

	return errors.New("Not a postman collection")
}

func ParseRequest(j []byte) ([]itemCollection, error) {

	req := postmanCollection{}

	if err := json.Unmarshal(j, &req); err != nil {
		fmt.Println(req, err)
		return req.ItemCollection, err
	}

	return req.ItemCollection, nil
}

func main() {
	file, err := readFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(EXIT_STATUS)
	}

	buffer, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println(err)
		os.Exit(EXIT_STATUS)
	}

	if err := CheckIfPostmanCollection(buffer); err != nil {
		fmt.Println("Postman Collection: ", err)
		os.Exit(EXIT_STATUS)
	}

	fmt.Println("We have a postman collection")
	fmt.Println("============================")

	reqs, err := ParseRequest(buffer)

	if err != nil {
		fmt.Println(err)
		os.Exit(EXIT_STATUS)
	}

	for _, item := range reqs[0:1] {
		fmt.Println("Running collection - ", item.FolderName)
		for k, reqitems := range item.RequestItems {
			fmt.Printf("\t Example #%d: Request item - %s \n", k+1, reqitems.RequestName)
			fmt.Println(reqitems.RequestSchema.Url.Raw)
		}
	}
}
