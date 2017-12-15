package controllers

import (
	"net/http"
	"fmt"
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	helloValue, err := app.Fabric.QueryHello()
	var dataVal = GetRequestHistoryValues()
	var historyValue = GetRequestHistory()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	data := &struct {
		Hello []string
		History []string
		HistoryVal []string
	}{
		Hello: helloValue,
		History: historyValue,
		HistoryVal: dataVal
	}
	fmt.Println("Request: ", r,"Data: ", data)
	renderTemplate(w, r, "home.html", data)
}
