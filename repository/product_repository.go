package repository

import (
	"database/sql"
	"fmt"
	"products-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.Product{}, err
	}

	var productsList []model.Product
	var productObj model.Product
	for rows.Next() {
		err = rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return []model.Product{}, err
		}
		productsList = append(productsList, productObj)
	}

	return productsList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int

	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *ProductRepository) GetProductByID(id int) (model.Product, error) {
	query := "SELECT * FROM product WHERE id = $1"
	row := pr.connection.QueryRow(query, id)

	var product model.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Product{}, fmt.Errorf("no product found with id %d", id)
		}
		fmt.Println("Error scanning row:", err)
		return model.Product{}, err
	}

	return product, nil
}
