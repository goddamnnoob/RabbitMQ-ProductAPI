package repositories

import (
	"goddamnnoob/RabbitMQ-ProductAPI/integrations"
	"goddamnnoob/RabbitMQ-ProductAPI/models"
	"log"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func AddProduct(product *models.Product) (er error) {
	newProductQuery := "INSERT INTO products (product_id,product_name,product_description,product_images,product_price,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7)"
	connection, err := integrations.GetNewPostgresConnection()
	if err != nil {
		return err
	}
	_, err = connection.Exec(newProductQuery, product.Product_id.String(), product.Productname, product.Productdescription, pq.Array(product.Productimages), product.Productprice, product.Createdat, product.Updatedat)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	connection.Close()
	return nil

}

func GetProductURLsFromID(productid uuid.UUID) (product *models.Product, er error) {
	product = &models.Product{Product_id: productid}
	getProductQuery := "SELECT product_images from products where product_id = $1"
	connection, er := integrations.GetNewPostgresConnection()
	if er != nil {
		return nil, er
	}
	defer connection.Close()
	log.Println(productid.String())
	er = connection.QueryRow(getProductQuery, productid.String()).Scan(pq.Array(&product.Productimages))

	if er != nil {
		return nil, er
	}

	return product, nil
}

func UpdateCompressedImagePaths(product *models.Product) (er error) {
	updateCompressedImgePathQuery := "UPDATE products SET compressed_product_images = $1 where product_id=$2"
	conection, er := integrations.GetNewPostgresConnection()

	if er != nil {
		return er
	}
	_, er = conection.Exec(updateCompressedImgePathQuery, pq.Array(product.Compressed_product_images), product.Product_id.String())

	if er != nil {
		return er
	}

	defer conection.Close()
	return nil
}
