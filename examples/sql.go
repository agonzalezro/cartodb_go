package main

import (
	"fmt"
	"io/ioutil"

	"github.com/agonzalezro/cartodb_go"
)

func main() {
	apiKey := "dd3cc8a9221e19576b8ac025c03e7b00738f874c"
	cartodbDomain := "alex"

	client := cartodb_go.NewCartoDBAPIKey(apiKey, cartodbDomain, "", "", "")
	response, err := client.Sql("select * from tweets", false, "json")
	if err != nil {
		fmt.Print(err)
		return
	}

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("Response: %+v", response)
	fmt.Printf("Body: %s", body)
	return
}
