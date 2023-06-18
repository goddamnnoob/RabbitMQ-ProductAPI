package handlers

import (
	"encoding/json"
	"goddamnnoob/RabbitMQ-ProductAPI/models"
	"goddamnnoob/RabbitMQ-ProductAPI/services"
	"io"
	"io/ioutil"
	"net/http"
)

type Httpres struct {
	Result string `json:"Result"`
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	if r.Method != "POST" {
		output, _ := json.Marshal(Httpres{
			Result: "Method Not Supported for this endpoint",
		})
		http.Error(w, string(output), 400)
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
	result, err := services.AddProduct(&product, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	output, _ := json.Marshal(Httpres{
		Result: result,
	})
	w.Write(output)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}
