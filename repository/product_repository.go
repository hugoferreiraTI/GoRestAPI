package repository

import (
	"database/sql"
	"fmt"
	"go_mod_api/model"
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
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var prodcutList []model.Product
	var productObje model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObje.Id,
			&productObje.Name,
			&productObje.Price,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		prodcutList = append(prodcutList, productObje)
	}

	rows.Close()
	return prodcutList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int

	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		"VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product

	err = query.QueryRow(id_product).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
	) //utilizado para executar a query e retornar os dados

	if err != nil {
		if err == sql.ErrNoRows { //não encontra o elemento procurado
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &product, nil
}
