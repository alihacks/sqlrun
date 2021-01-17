package common

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
)

type SqlRunConfig struct {
	ServerName string
	Port       uint16
	UserName   string
	Password   string
	Database   string
	Query      string
}

func RunSql(config SqlRunConfig) error {
	connString := fmt.Sprintf("server=%s", config.ServerName)

	if config.UserName != "" {
		connString += fmt.Sprintf(";user id=%s;password=%s", config.UserName, config.Password)
	}

	if config.Port > 0 {
		connString += fmt.Sprintf(";port=%d", config.Port)
	}

	if config.Database != "" {
		connString += fmt.Sprintf(";database=%s", config.Database)
	}

	//fmt.Printf(" connString:%s\n", connString)

	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer db.Close()

	var rows *sql.Rows
	rows, err = db.Query(config.Query)
	if err != nil {
		log.Fatal("Query failed:", err.Error())
		return err
	}

	var columns []string
	columns, err = rows.Columns()

	readCols := make([]interface{}, len(columns))
	writeCols := make([]sql.NullString, len(columns))
	for i, _ := range writeCols {
		readCols[i] = &writeCols[i]
	}

	fmt.Printf("%s\n", strings.Join(columns, ","))

	for rows.Next() {
		err := rows.Scan(readCols...)
		if err != nil {
			log.Fatal("Scan failed:", err.Error())
			return err
		}

		for _, value := range writeCols {
			fmt.Printf("%s,", value.String)
		}
		fmt.Printf("\n")

	}
	if err = rows.Err(); err != nil {
		log.Fatal("Scan failed:", err.Error())
		return err
	}

	return nil
}
