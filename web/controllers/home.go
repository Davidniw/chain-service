package controllers

import (
	"net/http"
	"fmt"
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	helloValue, err := app.Fabric.QueryHello()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	data := &struct {
		Hello string
	}{
		Hello: helloValue,
	}
	fmt.Println("Request: ", r,"Data: ", data)
	renderTemplate(w, r, "home.html", data)
}
