package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	var err error
	// my sql user root
	//db, err = sql.Open("sqlserver", "sqlserver://sa:P@ssw0rd@13.76.163.73:1433?database=techcoach")
	db, err = sql.Open("mysql", "root:P@ssw0rd@tcp(13.76.163.73:3306)/techcoach")
	if err != nil {
		panic(err)
	}

	//SQL server
	//query := "select * from cover where id=@id"
	//rows, err := db.Query(query, sql.Named("id", 1))

	//MySQL
	//query := "select * from cover where id=? or id=?"
	//rows, err := db.Query(query, 1, 2)

	// cover := Cover{
	// 	Id:   22,
	// 	Name: "sTone.",
	// }
	// err = AddCover(cover)
	// if err != nil {
	// 	panic(err)
	// }

	covers, err := GetCovers()
	if err != nil {
		panic(err)
	}

	for _, cover := range covers {

		fmt.Println(cover)
	}

}

type Cover struct {
	Id   int
	Name string
}

func GetCovers() ([]Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select * from cover"
	covers := []Cover{}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		cover := Cover{}
		err = rows.Scan(&cover.Id, &cover.Name)
		if err != nil {

			return nil, err
		}
		covers = append(covers, cover)
	}

	return covers, nil
}

func AddCover(cover Cover) error {

	err := db.Ping()
	if err != nil {
		return err
	}
	//MySQL
	query := "insert into cover value (?, ?)"
	result, err := db.Exec(query, cover.Id, cover.Name)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("Cannot insert")
	}

	return nil

}

func UpdateCover(cover Cover) error {

	err := db.Ping()
	if err != nil {
		return err
	}
	//MySQL
	query := "update cover set name=? where id=?"
	result, err := db.Exec(query, cover.Name, cover.Id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("Cannot Update")
	}

	return nil

}

func DeleteCover(cover Cover) error {

	err := db.Ping()
	if err != nil {
		return err
	}
	//MySQL
	query := "delete from cover where id=?"
	result, err := db.Exec(query, cover.Id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("Cannot Delete")
	}

	return nil

}
