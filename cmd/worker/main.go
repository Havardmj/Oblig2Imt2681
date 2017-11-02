package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/havard/Oblig2/cmd/database"
	"encoding/json"
	"time"
)

func ObtainCurrency(input string, db database.MgoDB) {

	content, err := http.Get(input)
	if err != nil {
		fmt.Printf("Could'nt Obtain new Currencies from fixer.io %v", err)
	}
	body, err := ioutil.ReadAll(content.Body)
	if err != nil {
		fmt.Printf("Read content Body failed in ObtainCurrency func : %v", err)
	}
	defer content.Body.Close()
	Data := database.Currency{}
	err = json.Unmarshal(body, &Data)
	if err != nil {
		fmt.Printf("Unmarshaling of data failed in ObtainCurrency func : %v", err)
	}

	db.Addcurrency(Data)


}
func main() {
	db := database.MgoDB {
		"mongodb://admin:imt2681@dds141274.mlab.com:41274/cloudimt2681",
		"cloudimt2681",
		"currency",
		"webhooks",
	}
	url := "http://api.fixer.io/latest?base=EUR"
	db.Init()
	ObtainCurrency(url, db)
	delay := time.Minute * 15
	for {
		time.Sleep(delay)
		db.Init()
		ObtainCurrency(url, db)
	}

}