package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Desc string `json:"description"`
	Code int    `json:"code"`
}

func main() {
	received := `
[
  {
    "description": "StatusOK",
    "code": 200
  },
  {
    "description": "StatusMovedPermanently",
    "code": 301
  },
  {
    "description": "StatusFound",
    "code": 302
  },
  {
    "description": "StatusSeeOther",
    "code": 303
  },
  {
    "description": "StatusTemporaryRedirect",
    "code": 307
  },
  {
    "description": "StatusBadRequest",
    "code": 400
  },
  {
    "description": "StatusUnauthorized",
    "code": 401
  },
  {
    "description": "StatusPaymentRequired",
    "code": 402
  },
  {
    "description": "StatusForbidden",
    "code": 403
  },
  {
    "description": "StatusNotFound",
    "code": 404
  },
  {
    "description": "StatusMethodNotAllowed",
    "code": 405
  },
  {
    "description": "StatusTeapot",
    "code": 418
  },
  {
    "description": "StatusInternalServerError",
    "code": 500
  }
]
  `

	var data []Data
	json.Unmarshal([]byte(received), &data)
	for _, v := range data {
		fmt.Println("[code:", v.Code, ", description:", v.Desc, "]")
	}
}
