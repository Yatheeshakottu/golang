package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type DB struct {
	Amount    string
	TimeStamp string
	Location  string
}

var (
	mp = make(map[string]DB)
)

func statistics(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "Get":
		fmt.Fprintf(w, "")
	default:
		fmt.Println("Invalid Request")
	}
}

func transactions(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "POST":
		b, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		data, _ := json.Marshal(b)
		var db DB
		json.Unmarshal(data, &db)
		mp[time.Now().Format(time.RFC3339)] = DB{
			Amount:    db.Amount,
			TimeStamp: db.TimeStamp,
			Location:  db.Location,
		}
	case "DELETE":
		mp = make(map[string]DB)
	case "PUT":
		for k := range mp {
			db := mp[k]
			db.Location = ""
			mp[k] = db
		}
	default:
		fmt.Println("Invalid Request")
	}
}

func main() {

	http.HandleFunc("/statistics", statistics)
	http.HandleFunc("/transactions", transactions)

	http.ListenAndServe(":8090", nil)
}
