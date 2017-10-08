package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type users struct {
	ID            int
	Name, Address string
}

func main() { //start of main func
	fmt.Println("starting up")
	r := mux.NewRouter()
	r.HandleFunc("/users", GetAllUsersFunc).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8888", nil)
} //end of main func

func GetAllUsersFunc(w http.ResponseWriter, r *http.Request) { //start of GetAllUserFunc

	fmt.Println("before--reaching in homehandler funcn")
	//fmt.Fprintf(w, "home")
	db, err := sql.Open("mysql", "root:root@/test")
	//db, err := sql.Open("mysql", "root:passapp@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	//fmt.Println(err)

	// fmt.Println(db)
	db.SetMaxIdleConns(100)
	rows, err := db.Query("SELECT * FROM user ")
	if err != nil {
		fmt.Print(err.Error())
	}
	//fmt.Println(rows)
	defer rows.Close()
	setDataRow := users{}
	results := []users{}
	for rows.Next() { //start of loop for fetching rows from db
		var id int
		var name, address string
		//setDataAppend :=users{}
		rows.Scan(&id, &name, &address)
		//fmt.Println(rows)
		setDataRow.ID = id
		setDataRow.Name = name
		setDataRow.Address = address
		// setDataRow = append(setDataRow , users)
		results = append(results, setDataRow)
		//fmt.Println(setDataRow)
		fmt.Println(results)

		//converting struct data into json
		b, err := json.Marshal(results) //will return data in byte formate
		if err != nil {
			fmt.Println(err)
			return
		}
		// string(b)-->will convert byte data into string but not will be formated
		fmt.Println(string(b)) //printing on console as formated bcz formate package using
		//json.NewEncoder(w).Encode(string(b))//not formated jason data as a  response from server
		fmt.Fprintf(w, string(b)) //formated json data as a response

	} //end of loop for fetching rows from db
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	fmt.Println("after--reaching in homehandler funcn")
} //end of GetAllUserFunc
