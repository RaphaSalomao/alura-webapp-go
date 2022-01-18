package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/RaphaSalomao/alura-webapp-go/database"
	"github.com/RaphaSalomao/alura-webapp-go/model"
	"github.com/google/uuid"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := database.FindAllProducts()
	fmt.Println(products)
	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		var err error
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			fmt.Println(err)
		}
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			fmt.Println(err)
		}
		p := model.Product{
			Name:        name,
			Description: description,
			Price:       price,
			Quantity:    quantity,
		}
		database.CreateProduct(&p)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	uuid, _ := uuid.Parse(id)
	database.DeleteProduct(uuid)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	uuid, _ := uuid.Parse(id)
	p := database.FindProduct(uuid)
	templates.ExecuteTemplate(w, "Edit", p)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	uuid, _ := uuid.Parse(id)
	p := database.FindProduct(uuid)
	p.Name = r.FormValue("name")
	p.Description = r.FormValue("description")
	p.Price, _ = strconv.ParseFloat(r.FormValue("price"), 64)
	p.Quantity, _ = strconv.Atoi(r.FormValue("quantity"))

	database.UpdateProduct(&p)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
