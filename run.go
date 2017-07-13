package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type Events []interface{}

//only using DB_user for now
type DBconf struct {
	DB_user string
	DB_pass string
	DB_host string
	DB_port string
}

func releaseTable(w http.ResponseWriter, r *http.Request) {
	conn := dbLogin()
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println(err)
	}
	//basic sql query to show releaseStatus table
	rows, err := db.Query("SELECT * FROM releaseStatus;")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}
	//creates object for data in each column
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		//gets hairy here but it seems to work
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	//marshal/unmarshal data from db, will likely break this out into another func
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		fmt.Println(err)
	}

	var e Events
	json.Unmarshal(jsonData, &e)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(e)
}

func deleteRow(w http.ResponseWriter, r *http.Request) {
	conn := dbLogin()
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println(err)
	}

	//check that row exists
	sqlcmd := "SELECT(relTicket) WHERE relTicket=?"
	err := db.QueryRow(sqlcmd)
	if relRow == 0 {
		log.Println(error)
	} else {
		//if it's all good, delete that row
		//need to update db function to delete row
		sql := "delete from relTicket=?;"
		rows, err := db.Exec(sql)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func dbLogin() string {
	file, _ := os.Open(".gitignore")
	decoder := json.NewDecoder(file)
	conf := DBconf{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error:", err)
	}
	return conf.DB_user
}

func main() {
	http.HandleFunc("/api", releaseTable)
	http.HandleFunc("/remove", deleteRow)
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.ListenAndServe(":9011", nil)
}
