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
}

type region struct {
	Name   Region
	Hotels []hotel
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("./index.tmpl.html"))
}

func main() {
	hotels := []region{
		{
			Name: Southern,
			Hotels: []hotel{
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
				},
				{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
				},
			},
		},
		{
			Name: Northern,
			Hotels: []hotel{
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
				},
				{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
				},
			},
		},
		{
			Name: Central,
			Hotels: []hotel{
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
				},
				{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
				},
			},
		},
	}

	err := tmpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatal(err)
	}
}
