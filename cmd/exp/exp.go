package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
	)
}

func main() {

	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5444",
		User:     "kim",
		Password: "1234",
		Database: "blog",
		SSLMode:  "disable",
	}

	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Connected to DB")

	// Create a table...
	_, err = db.Exec(`
	  CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE NOT NULL
	  );

	  CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		amout INT,
		description TEXT
	  );
	`)

	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created")

	//Insert some data..

	// name := "Mee"
	// email := "mee1091@gmail.com"
	// row := db.QueryRow(`
	//  INSERT INTO users(name,email)
	//  VALUES($1, $2) RETURNING id;`, name, email)
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User created. id = ", id)

	//query single record
	id := 7
	row := db.QueryRow(`
	SELECT name, email 
	FROM users
	WHERE  id = $1;
	`, id)
	var name, email string
	err = row.Scan(&name, &email)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User information: name: %s email: %s \n", name, email)

}
