package main

import (
	"html/template"
	"log"
	"os"
)

type Meal string

const (
	Breakfast = Meal("Breakfast")
	Lunch     = Meal("Lunch")
	Dinner    = Meal("Dinner")
)

type menuItem struct {
	Name        string
	Description string
	Meal        Meal
	Price       float64
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.tmpl.html"))
}

func main() {
	menu := []menuItem{
		{
			Name:        "Oatmeal",
			Description: "yum yum",
			Meal:        Breakfast,
			Price:       4.95,
		},
		{
			Name:        "Hamburger",
			Description: "Delicous good eating for you",
			Meal:        Lunch,
			Price:       6.95,
		},
		{
			Name:        "Pasta Bolognese",
			Description: "From Italy delicious eating",
			Meal:        Dinner,
			Price:       7.95,
		},
	}

	err := tmpl.Execute(os.Stdout, menu)
	if err != nil {
		log.Fatal(err)
	}
}
