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

type restaurant struct {
	Name string
	Menu []menu
}

type menu struct {
	Type  Meal
	Items []menuItem
}

type menuItem struct {
	Name        string
	Description string
	Price       float64
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.tmpl.html"))
}

func main() {
	restaurants := []restaurant{
		{
			Name: "Federicos",
			Menu: []menu{
				{
					Type: Breakfast,
					Items: []menuItem{
						{
							Name:        "Oatmeal",
							Description: "yum yum",
							Price:       4.95,
						},
						{
							Name:        "Cheerios",
							Description: "American eating food traditional now",
							Price:       3.95,
						},
						{
							Name:        "Juice Orange",
							Description: "Delicious drinking in throat squeezed fresh",
							Price:       2.95,
						},
					},
				},
				{
					Type: Lunch,
					Items: []menuItem{
						{
							Name:        "Hamburger",
							Description: "Delicous good eating for you",
							Price:       6.95,
						},
						{
							Name:        "Cheese Melted Sandwich",
							Description: "Make cheese bread melt grease hot",
							Price:       3.95,
						},
						{
							Name:        "French Fries",
							Description: "French eat potatoe fingers",
							Price:       2.95,
						},
					},
				},
				{
					Type: Dinner,
					Items: []menuItem{
						{
							Name:        "Pasta Bolognese",
							Description: "From Italy delicious eating",
							Price:       7.95,
						},
						{
							Name:        "Steak",
							Description: "Dead cow grilled bloody",
							Price:       13.95,
						},
						{
							Name:        "Bistro Potatoe",
							Description: "Bistro bar wood American bacon",
							Price:       6.95,
						},
					},
				},
			},
		},
		{
			Name: "Domenicos",
			Menu: []menu{
				{
					Type: Breakfast,
					Items: []menuItem{
						{
							Name:        "Oatmeal",
							Description: "yum yum",
							Price:       4.95,
						},
						{
							Name:        "Cheerios",
							Description: "American eating food traditional now",
							Price:       3.95,
						},
						{
							Name:        "Juice Orange",
							Description: "Delicious drinking in throat squeezed fresh",
							Price:       2.95,
						},
					},
				},
				{
					Type: Lunch,
					Items: []menuItem{
						{
							Name:        "Hamburger",
							Description: "Delicous good eating for you",
							Price:       6.95,
						},
						{
							Name:        "Cheese Melted Sandwich",
							Description: "Make cheese bread melt grease hot",
							Price:       3.95,
						},
						{
							Name:        "French Fries",
							Description: "French eat potatoe fingers",
							Price:       2.95,
						},
					},
				},
				{
					Type: Dinner,
					Items: []menuItem{
						{
							Name:        "Pasta Bolognese",
							Description: "From Italy delicious eating",
							Price:       7.95,
						},
						{
							Name:        "Steak",
							Description: "Dead cow grilled bloody",
							Price:       13.95,
						},
						{
							Name:        "Bistro Potatoe",
							Description: "Bistro bar wood American bacon",
							Price:       6.95,
						},
					},
				},
			},
		},
	}

	err := tmpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatal(err)
	}
}
