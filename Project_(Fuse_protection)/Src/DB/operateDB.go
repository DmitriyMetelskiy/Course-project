package database
/*
	functions for operating DB tables (collections)
*/
import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"github.com/labstack/echo"
	//"net/http"
)

func registerUser(login string, pwd string) {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	userCollection := session.DB("userdb").C("users")
	err = userCollection.Insert(&User{Id: bson.NewObjectId(), Login: login, Password: pwd, Discount: 0, Currency: 0.0})
	if err != nil {
		fmt.Println(err)
	}
}

//func removeProduct(c echo.Context) {
//	session, err := mgo.Dial("mongodb://127.0.0.1")
//	if err != nil {
//		panic(err)
//	}
//	defer session.Close()
//	productCollection := session.DB("productdb").C("products")
//	fmt.Println(productCollection)
//}

func main() {
	InitiateDB()
	// do some stuff
	//removeProduct("as")
}
