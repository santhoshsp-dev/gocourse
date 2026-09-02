package sqlconnect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("HOST")
	dbport := os.Getenv("DB_PORT")
	// connectionString := "root:Eagledon@7@tcp(127.0.0.1:3306)/" + dbname
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, dbport, dbname)
	// connectionString:="postgres://username:password@localhost:5432/database_name"
	// connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, dbport, dbname)
	db, err := sql.Open("mysql", connectionString)
	// db, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		// panic(err)
		return nil, err
	}
	return db, nil
}

/*
package sqlconnect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDb() (*sql.DB, error) {
	fmt.Println("Trying to connect to PostgreSQL...")

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("HOST")
	dbport := os.Getenv("DB_PORT")

	if host == "" {
		host = "localhost"
	}
	if dbport == "" {
		dbport = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "school"
	}

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		dbport,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	// Actually verify that PostgreSQL is reachable.
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL")

	return db, nil
}
*/
