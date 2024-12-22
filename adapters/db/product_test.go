package db_test

import (
	"database/sql"
	"testing"

	"github.com/bruno3du/hexagonal/adapters/db"
	"github.com/bruno3du/hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
 			"id" varchar(255) NOT NULL,
 			"name" varchar(255) NOT NULL,
 			"price" float NOT NULL,
 			"status" varchar(255),
 			PRIMARY KEY ("id")
			)`

	stmt, err := db.Prepare(table)

	if err != nil {
		panic(err)
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price, status) VALUES(?, ?, ?, ?)")

	if err != nil {
		panic(err)
	}

	stmt.Exec("abc", "Product 1", 10.1, "disabled")
}

func TestPruductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.1, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "enabled"

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

}
