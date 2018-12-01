package main

import (
	"fmt"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	session, err := mgo.Dial("mongodb://localhost")
	fmt.Println("Starting...")
	if err != nil {
		e.Logger.Fatal(err) //add log to ECHO.
		//return err
	}

	h := &handler{
		mongo: session,
		db:    "phonebooks",
		col:   "phonebook",
	}

	e.GET("/", h.welcome)
	e.GET("/api/list", h.list)
	e.POST("/api/insert", h.create)
	e.GET("/api/search/:firstname", h.search)
	e.PUT("/api/update/:firstname", h.update)
	e.DELETE("/api/delete/:firstname", h.delete)
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":1324"))
}

type phonebook struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Firstname string        `json:"firstname" bson:"firstname"`
	Lastname  string        `json:"lastname" bson:"lastname"`
	Telephone string        `json:"telephone" bson:"telephone"`
	Address   string        `json:"address" bson:"address"`
}

type handler struct {
	mongo *mgo.Session
	db    string
	col   string
}

func (h *handler) welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Beam Test!!")
}

func (h *handler) list(c echo.Context) error {
	conn := h.mongo.Copy()
	defer conn.Close()
	var ts []phonebook
	if err := conn.DB(h.db).C(h.col).Find(nil).All(&ts); err != nil {
		return err
	}
	c.JSON(http.StatusOK, ts)
	return nil
}

func (h *handler) create(c echo.Context) error {
	id := bson.NewObjectId()
	var t phonebook
	if err := c.Bind(&t); err != nil {
		return err
	}
	t.ID = id

	conn := h.mongo.Copy()
	defer conn.Close()
	if err := conn.DB(h.db).C(h.col).Insert(t); err != nil {
		return err
	}

	c.JSON(http.StatusOK, t)
	return nil
}

func (h *handler) search(c echo.Context) error {
	conn := h.mongo.Copy()
	defer conn.Close()
	var t phonebook
	fname := c.Param("firstname")
	fmt.Println(fname)
	if err := conn.DB(h.db).C(h.col).Find(bson.M{"firstname": fname}).One(&t); err != nil {
		return err
	}
	c.JSON(http.StatusOK, t)
	return nil
}

func (h *handler) update(c echo.Context) error {
	conn := h.mongo.Copy()
	defer conn.Close()
	var t phonebook

	fname := c.Param("firstname")
	//fmt.Println(fname)
	if err := conn.DB(h.db).C(h.col).Find(bson.M{"firstname": fname}).One(&t); err != nil {
		fmt.Println(err)
		return err
	}

	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}

	// Update
	colQuerier := bson.M{"firstname": fname}

	change := bson.M{"$set": bson.M{"firstname": m["firstname"], "lastname": m["lastname"], "telephone": m["telephone"], "address": m["address"]}}
	err := conn.DB(h.db).C(h.col).Update(colQuerier, change)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, m)
	return nil
}

func (h *handler) delete(c echo.Context) error {
	conn := h.mongo.Copy()
	defer conn.Close()
	//id := bson.ObjectIdHex(c.Param("id"))

	fname := c.Param("firstname")
	fmt.Println(fname)
	if err := conn.DB(h.db).C(h.col).Remove(bson.M{"firstname": fname}); err != nil {
		return err
	}
	c.JSON(http.StatusOK, echo.Map{
		"result": "success",
	})
	return nil
}
