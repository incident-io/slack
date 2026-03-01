package main

import (
	"fmt"

	"github.com/incident-io/slack"
)

func main() {
	api := slack.New("YOUR_TOKEN_HERE")
	params := slack.UploadFileParameters{
		Title:    "Batman Example",
		File:     "example.txt",
		Filename: "example.txt",
		FileSize: 100,
	}
	file, err := api.UploadFile(params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Title: %s\n", file.ID, file.Title)

	err = api.DeleteFile(file.ID)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("File %s deleted successfully.\n", file.ID)
}
