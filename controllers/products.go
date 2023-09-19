package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"web-application/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		priceStr := r.FormValue("price")
		qtyStr := r.FormValue("qty")

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Println("Error when converting price", err)
		}

		qty, err := strconv.Atoi(qtyStr)
		if err != nil {
			log.Println("Error when converting qty", err)
		}

		models.CreateNewProduct(name, description, price, qty)

	}
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.GetProduct(productId)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idStr := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		priceStr := r.FormValue("price")
		qtyStr := r.FormValue("qty")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("Error when converting id", err)
		}

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Println("Error when converting price", err)
		}

		qty, err := strconv.Atoi(qtyStr)
		if err != nil {
			log.Println("Error when converting qty", err)
		}

		models.UpdateProduct(id, qty, name, description, price)

	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	if productId != "" {
		models.DeleteProduct(productId)
	}

	http.Redirect(w, r, "/", 301)
}
