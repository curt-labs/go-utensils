package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"os"
)

type Scanner interface {
	Scan(...interface{}) error
}

var (
	DB      *sql.DB
	AdminDB *sql.DB
	B2BDB   *sql.DB
	VcdbDB  *sql.DB
	Driver  = "mysql"
)

// InitDB Initiates a conenction to the CurtData database
func InitDB() (*sql.DB, error) {
	var err error
	if DB == nil {
		db := "CurtData"
		if d := os.Getenv("CURT_DEV_NAME"); d != "" {
			db = d
		}
		DB, err = sql.Open(Driver, ConnectionString(db, false))
		if err != nil {
			return nil, err
		}
	}
	return DB, nil
}

// InitVCDB Initiates a conenction to the VCDB database
func InitVCDB() (*sql.DB, error) {
	var err error
	if VcdbDB == nil {
		db := "vcdb"
		if d := os.Getenv("VCDB_NAME"); d != "" {
			db = d
		}
		VcdbDB, err = sql.Open(Driver, ConnectionString(db, true))
		if err != nil {
			return nil, err
		}
	}
	return VcdbDB, nil
}

// InitAdmin Initiates a conenction to the admin database
func InitAdmin() (*sql.DB, error) {
	var err error
	if AdminDB != nil {
		err = AdminDB.Ping()
		if err == nil {
			return AdminDB, nil
		}
	}

	db := "admin"
	if d := os.Getenv("ADMIN_NAME"); d != "" {
		db = d
	}
	AdminDB, err = sql.Open(Driver, ConnectionString(db, true))

	return AdminDB, err
}

// InitB2B Initiates a conenction to the mysql version of the old b2b database
func InitB2B() (*sql.DB, error) {
	var err error
	if B2BDB != nil {
		err = B2BDB.Ping()
		if err == nil {
			return B2BDB, nil
		}
	}

	db := "b2b"
	if d := os.Getenv("B2B_NAME"); d != "" {
		db = d
	}
	B2BDB, err = sql.Open(Driver, ConnectionString(db, true))

	return B2BDB, err
}

// ConnectionString Generates a MySQL connection string
func ConnectionString(db string, parseTime bool) string {
	if addr := os.Getenv("DATABASE_HOST"); addr != "" {
		proto := os.Getenv("DATABASE_PROTOCOL")
		user := os.Getenv("DATABASE_USERNAME")
		pass := os.Getenv("DATABASE_PASSWORD")

		return fmt.Sprintf("%s:%s@%s(%s)/%s?loc=%s", user, pass, proto, addr, db, "America%2FChicago")
	}

	return fmt.Sprintf("root:@tcp(127.0.0.1:3306)/%s?parseTime=%t&loc=America%%2FChicago", db, parseTime)
}
