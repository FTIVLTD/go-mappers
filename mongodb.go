package mappers

import (
	"fmt"
	"strconv"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
Mongo - MongoDB mapper
*/
type Mongo struct {
	DBConfig   DBConfig
	Collection string
	Conn       *mgo.Session
	limit      int
}

/*
Connect - connecting to DB
*/
func (m *Mongo) Connect() error {
	m.log("Connecting to: ", m.getDBInfo())
	session, err := mgo.Dial(m.prepareConnectionString())
	m.Conn = session
	if err != nil {
		return err
	}
	return nil
}

func (m *Mongo) Limit(limit int) Mongo {
	m.limit = limit
	return m
}

/*
Create - inserting new enity
*/
func (m *Mongo) Create(data interface{}) error {
	if m.Conn == nil {
		err := m.Connect()
		if err != nil {
			fmt.Println("Error connecting: ", err)
			return nil
		}
	}
	c := m.Conn.DB(m.DBConfig.Database).C(m.Collection)
	return c.Insert(data)
}

/*
Search - Searching data in db between fromDate and toDate and substr in "message"
*/
func (m *Mongo) Search(query bson.M) (interface{}, error) {
	if m.Conn == nil {
		err := m.Connect()
		if err != nil {
			fmt.Println("Error connecting: ", err)
			return nil, err
		}
	}
	c := m.Conn.DB(m.DBConfig.Database).C(m.Collection)
	var data []interface{}
	err := c.Find(query).Limit(m.limit).All(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

/*
Close - closing connection
*/
func (m *Mongo) Close() error {
	if m.Conn == nil {
		return nil
	}
	m.Conn.Close()
	return nil
}

/*
Converts db config into connection string
[mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
*/
func (m *Mongo) prepareConnectionString() string {
	c := m.DBConfig
	str := "mongodb://"
	if c.User != "" {
		str += c.User + c.Password + "@"
	}
	str += c.Host
	if c.Port != 0 {
		str += ":" + strconv.Itoa(c.Port)
	}
	if c.Database != "" {
		str += "/" + c.Database
	}
	return str
}

func (m *Mongo) getDBInfo() string {
	c := m.DBConfig
	return "mongodb://" + c.User + "***@" + c.Host + ":" + strconv.Itoa(c.Port) + "/" + c.Database
}

func (m *Mongo) log(data ...interface{}) error {
	fmt.Println(data...)
	return nil
}
