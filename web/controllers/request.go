package controllers

import (
	"net/http"
    "github.com/gorilla/mux"
)
var history []string
func (app *Application) RequestHandler(w http.ResponseWriter, r *http.Request) {

    router.HandleFunc("/add/{input}", TodoShow)
    vars := mux.Vars(r)
    input := vars["input"]
    fmt.Fprintln(w, "input show:", input)

    vars := mux.Vars(r)
    input := vars["input"]

	data := &struct {
		TransactionId string
		Success bool
		Response bool
	}{
		TransactionId: "",
		Success: false,
		Response: false,
	}
	if r.FormValue("submitted") == "true" {
		helloValue := r.FormValue("hello")
		txid, err := app.Fabric.InvokeHello(helloValue)
		if err != nil {
			http.Error(w, "Unable to invoke hello in the blockchain", 500)
		}
		data.TransactionId = txid
		data.Success = true
		data.Response = true
	}
	history = append(history, data.TransactionId)
	//fmt.Println("History:",history)
	renderTemplate(w, r, "request.html", data)
}

func GetRequestHistory() []string {
    return history
}