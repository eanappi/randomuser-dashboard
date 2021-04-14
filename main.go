package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/eanappi/randomuser"
)

func renderDashboard(w http.ResponseWriter, r *http.Request) {
	var data []map[string]string

	users, err := randomuser.NewUsers(10)
	if err != nil {
		fmt.Errorf("The service is offline by: %s", err)
	}
	for index := range users.Results {
		data = append(data, users.Summary(index))
	}
	dashboardTemplate, _ := template.ParseFiles("templates/index.tmpl")
	err = dashboardTemplate.Execute(w, data)
	if err != nil {
		log.Println("Error on template :", err)
		return
	}
}

func main() {
	http.HandleFunc("/", renderDashboard)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("Error initializing server:", err)
		return
	}
}
