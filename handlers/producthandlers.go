package handlers

import (
	"encoding/json"
	"goddamnnoob/Zocket-assignment/models"
	"io"
	"io/ioutil"
	"net/http"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		io.WriteString(w, "Invalid HTTP Method")
		return
	}

	var user models.User
	var product models.Product
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	json.Unmarshal(body, &product)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("content-type", "application/json")
	output, err := json.Marshal(product)
	w.Write(output)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}
