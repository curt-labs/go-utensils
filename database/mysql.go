package database

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"os"
)

type Scanner interface {
	Scan(...interface{}) error
}

var (
	EmptyDb = flag.String("cleaner", "", "bind empty database with structure defined")
	DB      *sql.DB
	VcdbDB  *sql.DB
	Driver  = "mysql"
)

func InitDB() (*sql.DB, error) {
	var err error
	var DB *sql.DB
	if DB == nil {
		DB, err = sql.Open(Driver, ConnectionString())
		if err != nil {
			return nil, err
		}
	}
	return DB, nil
}

func InitVCDB() (*sql.DB, error) {
	var err error
	var VcdbDB *sql.DB
	if VcdbDB == nil {
		VcdbDB, err = sql.Open(Driver, VcdbConnectionString())
		if err != nil {
			return nil, err
		}
	}
	return VcdbDB, nil
}

func ConnectionString() string {
	if addr := os.Getenv("DATABASE_HOST"); addr != "" {
		proto := os.Getenv("DATABASE_PROTOCOL")
		user := os.Getenv("DATABASE_USERNAME")
		pass := os.Getenv("DATABASE_PASSWORD")
		db := os.Getenv("CURT_DEV_NAME")

		return fmt.Sprintf("%s:%s@%s(%s)/%s?loc=%s", user, pass, proto, addr, db, "America%2FChicago")
	}

	if EmptyDb != nil && *EmptyDb != "" {
		return "root:@tcp(127.0.0.1:3306)/CurtData_Empty?&loc=America%2FChicago"
	}
	return "root:@tcp(127.0.0.1:3306)/CurtData?loc=America%2FChicago"
}

func VcdbConnectionString() string {
	if addr := os.Getenv("DATABASE_HOST"); addr != "" {
		proto := os.Getenv("DATABASE_PROTOCOL")
		user := os.Getenv("DATABASE_USERNAME")
		pass := os.Getenv("DATABASE_PASSWORD")
		db := os.Getenv("VCDB_NAME")
		return fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=true&loc=%s", user, pass, proto, addr, db, "America%2FChicago")
	}

	return "root:@tcp(127.0.0.1:3306)/vcdb?parseTime=true&loc=America%2FChicago"
}
