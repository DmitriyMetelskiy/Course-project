package database
/*
	creating initial database structure on mock data
 */
import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id bson.ObjectId `bson:"_id"`		// id
	Login string `bson:"login"`			// login
	Password string `bson:"password"`	// pwd
	Discount int `bson:"discount"`		// discount in % for not-under-discount products
	Currency float64 `bson:"currency"`	// money in the account
}

type Product struct {
	// if there is no id specified, mongoDB will automatically add one
	Type string `bson:"type"`		// knee/shin/elbow pads, gloves, helmets, combo, wrist/ankle support and other...
	Model string `bson:"model"`		// alpha/delta/echo
	Size string `bson:"size"`		// S/M/L/XL
	Price float64 `bson:"price"`	// product's price
	Location int `bson:"location"`	// [available in ...] shop id
// todo optimizations: store not location, but slice of locations with amount of stored items there
}

/*
type Shop struct {
	Id bson.ObjectId `bson:"_id"`		// shop id
	Address string `bson:"address"`		// show address
	Phone string `bson:"phone"`			// shop phone number
*/

func InitiateDB() {
// open connection
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
// get collections
	userCollection := session.DB("userdb").C("users")
	productCollection := session.DB("productdb").C("products")
// "array" of test users
	testUsers := []User {
		{Id: bson.NewObjectId(), Login: "dimon", Password: "12456", Discount: 3, Currency: 100.0},
		{Id: bson.NewObjectId(), Login: "timon", Password: "qwert", Discount: 0, Currency: 50.0},
		{Id: bson.NewObjectId(), Login: "pumba", Password: "qwer", Discount: 0, Currency: 164.5},
		// ...
	}
// add mock users in 'users' collection
	err = userCollection.Insert(testUsers)
	if err != nil {
		fmt.Println(err)
	}
// "array" of test products
	testProducts := []Product {
		{Type: "knee pads", Model: "delta", Size: "M", Price: 135, Location: 0},
		{Type: "elbow pads", Model: "delta", Size: "M", Price: 65, Location: 0},
		{Type: "wrist support", Model: "alpha", Size: "M", Price: 40, Location: 1},
		{Type: "shin pads", Model: "alpha", Size: "S", Price: 80, Location: 0},
		// ...
	}
// add mock products in 'products' collection
	err = productCollection.Insert(testProducts)
	if err != nil {
		fmt.Println(err)
	}
}