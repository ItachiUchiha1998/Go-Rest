package store

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

const SERVER = "mongodb://localhost:27017"
const DBNAME = "dummyStore"
const COLLECTION = "store"

var productId = 10

func (r Repository) GetProducts() Products { // get all products
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Error in making connection to db..")
	}
	defer session.Close()
	c := session.DB(DBNAME).C(COLLECTION)
	results := Products{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

func (r Repository) GetProductById(id int) Product {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	var result Product

	fmt.Println("ID in GetProductById", id)

	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	return result
}

func (r Repository) GetProductsByString(query string) Products {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(COLLECTION)

	result := Products{}

	qs := strings.Split(query, " ")
	and := make([]bson.M, len(qs))
	for i, q := range qs {
		and[i] = bson.M{"title": bson.M{
			"$regex": bson.RegEx{Pattern: ".*" + q + ".*", Options: "i"},
		}}
	}
	filter := bson.M{"$and": and}

	if err := c.Find(&filter).Limit(5).All(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	return result
}

func (r Repository) AddProduct(product Product) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	productId += 1
	product.ID = productId
	session.DB(DBNAME).C(COLLECTION).Insert(product)
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added New Product ID- ", product.ID)

	return true
}

func (r Repository) UpdateProduct(product Product) bool { // update
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	err = session.DB(DBNAME).C(COLLECTION).UpdateId(product.ID, product)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Updated Product ID - ", product.ID)

	return true
}

func (r Repository) DeleteProduct(id int) string { // delete a product
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	if err = session.DB(DBNAME).C(COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	fmt.Println("Deleted Product ID - ", id)
	return "OK"
}
