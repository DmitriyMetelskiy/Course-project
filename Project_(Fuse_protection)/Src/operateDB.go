package main
/*
	functions for operating DB tables (collections)
*/
import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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



func main() {
	InitiateDB()
	// do some stuff
}
