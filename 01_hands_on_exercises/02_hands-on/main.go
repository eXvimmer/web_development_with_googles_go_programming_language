package main

import (
	"html/template"
	"log"
	"os"
)

type Region string

const (
	Northern = Region("Northern")
	Central  = Region("Central")
	Southern = Region("Southern")
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  Region
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("./index.tmpl.html"))
}

func main() {
	hotels := []hotel{
		{
			Name:    "Hotel California",
			Address: "42 Sunset Boulevard",
			City:    "Los Angeles",
			Zip:     "95612",
			Region:  Southern,
		},
		{
			Name:    "Hotel Babylon",
			Address: "13 somewhere close",
			City:    "Los Angeles",
			Zip:     "95613",
			Region:  Northern,
		},
		{
			Name:    "Hotel Awesome",
			Address: "13 somewhere far far away",
			City:    "Los Angeles",
			Zip:     "95614",
			Region:  Central,
		},
	}

	err := tmpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatal(err)
	}
}
