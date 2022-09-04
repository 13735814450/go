package service

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Student struct {
	Id int
	Name string
	Phone string
	Email string
	Sex string
}
func MoConnecToDB() *mgo.Collection {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("wxh").C("student")
	return c
}

func MoInsert() {
	c := MoConnecToDB()
	stu1 := Student{
		Name: "zhangsan",
		Phone: "13480989765",
		Email: "329832984@qq.com",
		Sex: "F",
	}
	stu2 := Student{
		Name: "liss",
		Phone: "13980989767",
		Email: "12832984@qq.com",
		Sex: "M",
	}
	err := c.Insert(&stu1, &stu2)
	if err != nil {
		log.Fatal(err)
	}
}
func MoGetDataViaSex() {
	c := MoConnecToDB()
	result := Student{}
	err := c.Find(bson.M{"sex": "M"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("student", result)
	students := make([]Student, 20)
	err = c.Find(nil).All(&students)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(students)

}

func MoUpdateDBViaId() {
	//id := bson.ObjectIdHex("5a66a96306d2a40a8b884049")
	c := MoConnecToDB()
	err := c.Update(bson.M{"email": "12832984@qq.com"}, bson.M{"$set": bson.M{"name": "haha", "phone": "37848"}})
	if err != nil {
		log.Fatal(err)
	}
}

func MoRemoveFromMgo() {
	c := MoConnecToDB()
	_, err := c.RemoveAll(bson.M{"phone": "13480989765"})
	if err != nil {
		log.Fatal(err)
	}
}
