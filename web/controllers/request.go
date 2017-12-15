package controllers

import (
	"net/http"
    //"github.com/gorilla/mux"
    //"fmt"
)
var history []string
var dataVal []string
var input string
func (app *Application) RequestHandler(w http.ResponseWriter, r *http.Request) {


    //router := mux.NewRouter().StrictSlash(true)
    //router.HandleFunc("/add/{input}", GetInput)

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

		history = append(history, data.TransactionId)
        //helloValue, err := app.Fabric.QueryHello()
        if err != nil {
                http.Error(w, "Unable to query the blockchain", 500)
            }
        dataVal = append(dataVal, helloValue)
	}

	renderTemplate(w, r, "request.html", data)
}

func GetRequestHistory() []string {
    return history
}

func GetRequestHistoryValues() []string {
    return dataVal
}
/*
func GetInput(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    input := vars["input"]
    fmt.Fprintln(w, "data show:", input)
}
*/