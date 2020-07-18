package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/taosdata/driver-go/taosSql"
	"time"
)

var (
	db1       *sql.DB
	db2       *sql.DB
	err       error
	dbName    string
	tableName string
)

func Init(){
	tableName = "test"
	dbName = "test"
	db1, err = sql.Open("taosSql", "root:taosdata@/tcp(172.21.0.2:0)/")
	checkErr(err)
	//res, err := db1.Exec("create database test; ")
	//if err != nil {
	//	log.Println(res)
	//}
	//checkErr(err)
	db2, err = sql.Open("taosSql", "root:taosdata@/tcp(172.21.0.2:0)/test")
	checkErr(err)
}

func CreateTable(db *sql.DB, demot string) {
	st := time.Now().Nanosecond()
	// create table
	res, err := db.Exec("create table test.test (ts timestamp, id int, name binary(20), len tinyint, flag bool, notes binary(20), fv float, dv double)")
	checkErr(err)

	affectd, err := res.RowsAffected()
	checkErr(err)

	et := time.Now().Nanosecond()
	fmt.Printf("create table result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1E9)
}

func InsertData(db *sql.DB, dbName string) {
	st := time.Now().Nanosecond()
	// insert data into table
	stmt, err := db.Prepare("insert into ? values(?, ?, ?, ?, ?, ?, ?, ?) (?, ?, ?, ?, ?, ?, ?, ?) (?, ?, ?, ?, ?, ?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec(dbName, "now", 1000, "'haidian'", 6, true, "'AI world'", 6987.654, 321.987,
		"now+1s", 1001, "'changyang'", 7, false, "'DeepMode'", 12356.456, 128634.456,
		"now+2s", 1002, "'chuangping'", 8, true, "'database'", 3879.456, 65433478.456)
	checkErr(err)

	affectd, err := res.RowsAffected()
	checkErr(err)

	et := time.Now().Nanosecond()
	fmt.Printf("insert data result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1E9)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func runServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func runTdengine() {
	defer db1.Close()
	defer db2.Close()

	CreateTable(db1, tableName)
	InsertData(db2, tableName)
}

func main() {
	//runServer()
	Init()
	runTdengine()
}
