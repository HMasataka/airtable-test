package main

import (
	"fmt"

	at "github.com/mehanizm/airtable"
)

func main() {
	client := at.NewClient("key2aujiePk1myeR1")
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
}
