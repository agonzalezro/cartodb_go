package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/agonzalezro/cartodb_go"
)

func main() {
	apiKey := os.Getenv("CARTODB_API_KEY")
	username := os.Getenv("CARTODB_USERNAME")

	client := cartodb.NewAPIKeyClient(apiKey, username, "", "", "")
	response, err := client.SQL("select * from tweets", "GET", "json")
	if err != nil {
		fmt.Print(err)
		return
	}

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("Response: %+v", response)
	fmt.Printf("Body: %s", body)
	return
}
