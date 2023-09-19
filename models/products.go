package models

import "web-application/db"

type Product struct {
	Name, Description string
	Price             float64
	Id, Qty           int
}

func GetAllProducts() []Product {
	db := db.DatabaseConnection()

	selectAllProducts, err := db.Query("SELECT * FROM products ORDER BY id ASC")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, qty int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &qty)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Qty = qty

		products = append(products, p)

	}
	defer db.Close()
	return products
}

func GetProduct(id string) Product {
	db := db.DatabaseConnection()

	selectProduct, err := db.Query("SELECT * FROM products WHERE id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for selectProduct.Next() {
		var id, qty int
		var name, description string
		var price float64

		err = selectProduct.Scan(&id, &name, &description, &price, &qty)
		if err != nil {
			panic(err.Error())
		}
		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Qty = qty

	}

	defer db.Close()
	return product
}

func CreateNewProduct(name, description string, price float64, qty int) {
	db := db.DatabaseConnection()

	insertData, err := db.Prepare("INSERT INTO products(name, description, price, qty) VALUES($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, qty)
	defer db.Close()
}

func UpdateProduct(id, qty int, name, description string, price float64) {
	db := db.DatabaseConnection()

	updateData, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, qty=$4 WHERE id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateData.Exec(name, description, price, qty, id)
	defer db.Close()
}

func DeleteProduct(id string) {

	db := db.DatabaseConnection()

	deleteData, err := db.Prepare("DELETE FROM products WHERE id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteData.Exec(id)
	defer db.Close()
}
