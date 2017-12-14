package controllers

import (
	"net/http"
	"fmt"
)
var dataVal []string
func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	helloValue, err := app.Fabric.QueryHello()
	dataVal = append(dataVal, helloValue)
	var historyValue = GetRequestHistory()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	data := &struct {
		Hello []string
		History []string
	}{
		Hello: dataVal,
		History: historyValue,
	}
	fmt.Println("Request: ", r,"Data: ", data)
	renderTemplate(w, r, "home.html", data)
}
