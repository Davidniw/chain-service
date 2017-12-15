package controllers

import (
	"net/http"
	"fmt"
)
type Hell struct {
		Guid string
		Value string
	}
var dataStruct []Hell
func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	helloValue, err := app.Fabric.QueryHello()
	var dataVal = GetRequestHistoryValues()
	var dataValT = GetRequestHistoryT()
	var historyValue = GetRequestHistory()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	data := &struct {
		Hello string
		History []string
		HistoryVal []string
		HistoryObj []Hell
	}{
		Hello: helloValue,
		History: historyValue,
		HistoryVal: dataVal,
		HistoryObj: dataValT,
	}
	fmt.Println("Request: ", r,"Data: ", data)
	renderTemplate(w, r, "home.html", data)
}
