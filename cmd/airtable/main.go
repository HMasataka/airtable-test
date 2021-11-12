package main

import (
	"fmt"

	"github.com/mehanizm/airtable"
)

func main() {
	client := airtable.NewClient("key2aujiePk1myeR1")
	table := client.GetTable("appD8wsx3DiEMakAr", "Table1")
	records, err := table.GetRecords().
		FromView("view1").
		Do()
	if err != nil {
		panic(err)
	}
	for _, r := range records.Records {
		fmt.Printf("%+v\n", r)
	}

	fmt.Println()
	rrr, err := table.GetRecord("recLrbPrGWLHWtp6U")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", rrr)
}
