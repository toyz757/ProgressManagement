package main

import (
  "os"
  
  "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

func gormConnect() *gorm.DB {
  DBMS     := "mysql"
  USER     := os.Getenv("MYSQL_USER")
  PASS     := os.Getenv("MYSQL_PASS")
  PROTOCOL := "tcp(127.0.0.1:3306)"
  DBNAME   := "progress_management"

  CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
  db,err := gorm.Open(DBMS, CONNECT)

  if err != nil {
    panic(err.Error())
  }
  return db
}

func main(){
  db := gormConnect()
  
  defer db.Close()
}