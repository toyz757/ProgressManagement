package main

import (
  "os"
  "time"
  
  "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)


type ticket struct {
  id int
  title string `gorm:"default:'galeone'"`
  responsible string
  deadline time.Time
}

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

  // 構造体のインスタンス化
  ticketEx := ticket{id: 100, title: "title1", responsible: "toy", deadline: time.Now()}

  println("A地点到達")
  // INSERTを実行
  db.Create(&ticketEx)
  println("ticket id = "+ticketEx.title)

  println("B地点到達")
  time.Sleep(1 * time.Second) //おそらく、早すぎてDBのエラーが出ないことがあるための暫定対策
}