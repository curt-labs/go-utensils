package mongo

import (
	"gopkg.in/mgo.v2"

	"os"
	"strings"
	"time"
)

var (
	MongoDatabase string
	MongoSession  *mgo.Session
)

func MongoConnectionString(database ...string) *mgo.DialInfo {
	var info mgo.DialInfo
	addr := os.Getenv("MONGO_URL")
	if addr == "" {
		addr = "127.0.0.1"
	}
	addrs := strings.Split(addr, ",")
	info.Addrs = append(info.Addrs, addrs...)

	info.Username = os.Getenv("MONGO_USERNAME")
	info.Password = os.Getenv("MONGO_PASSWORD")
	info.Database = os.Getenv("MONGO_DATABASE")
	info.Timeout = time.Second * 2
	info.FailFast = true
	if len(database) > 0 {
		info.Database = database[0]
	}
	info.Source = "admin"

	return &info
}

func InitMongo(database ...string) error {
	var err error
	var db_name string
	if len(database) > 0 {
		db_name = database[0]
	}
	if MongoSession == nil {
		connectionString := MongoConnectionString(db_name)
		MongoSession, err = mgo.DialWithInfo(connectionString)
		if err != nil {
			return err
		}
		MongoDatabase = connectionString.Database
	}
	return nil
}
