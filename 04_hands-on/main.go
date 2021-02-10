package main

import "fmt"

// Menu stores values for restaurant
type Menu struct {
	Dish, Hour string
	Price      float64
}

// Restaurant stores values for menu
type Restaurant struct {
	Name string
	Menu
}

// encapsulates into slice of structs so I can have multiple restaurant names
type restaurants []Restaurant

func main() {
	r := restaurants{
		Restaurant{
			Name: "Jordan's pizza palace",
			Menu: Menu{
				Dish:  "potatos",
				Hour:  "dinner",
				Price: 34.99,
			},
		},
		Restaurant{
			Name: "Pizza Palace",
			Menu: Menu{
				Dish:  "Pizza",
				Hour:  "dinner",
				Price: 34.99,
			},
		},
	}
	fmt.Println(r)
}
