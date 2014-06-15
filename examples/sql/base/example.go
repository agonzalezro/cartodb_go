package main

import (
	"fmt"
	"io/ioutil"

	"github.com/agonzalezro/cartodb_go"
)

func main() {
	client := cartodb.NewBaseClient("alex", "", "", "")
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
