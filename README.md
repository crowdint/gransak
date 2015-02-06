# Gransak

###Golang implementation of Ransack

####Install

Download the repository with:
``go get github.com/crowdint/gransak``

Then include the project inside your project like this:
``import "github.com/crowdint/gransak"``

Currently gransak transforms a ransack like string into a sql 'where' statement e.g.
    package main

    import (
      . "github.com/crowdint/gransak/filter"
      "fmt"
    )
    
    func main() {
      sql := Gransak.ToSql("user_name_eq", "cone")

      fmt.Println(sql)
      //prints: user_name = 'cone'
    }
    
####Operations currently supported
