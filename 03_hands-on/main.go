package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type item struct {
	Name, Hour string
	Price      float64
}

type items []item

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	menu := items{
		item{
			Name:  "Cold cereal",
			Hour:  "Breakfast",
			Price: 3.99,
		},
		item{
			Name:  "Brisket",
			Hour:  "Dinner",
			Price: 12.99,
		},
		item{
			Name:  "Pizza",
			Hour:  "Lunch",
			Price: 8.99,
		},
	}
	err := tpl.Execute(os.Stdout, menu)
	if err != nil {
		log.Fatalln(err)
	}
}
