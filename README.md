# Gransak

###Golang implementation of Ransack

####Install

Download the repository with:
``go get github.com/crowdint/gransak``

Then include the project inside your project like this:
``import "github.com/crowdint/gransak"``

####What it does

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

#####OR

    Gransak.ToSql("first_name_or_last_name_cont", "cone")
    //returns: first_name LIKE '%cone%' OR last_name LIKE '%cone%' 
    
#####AND

    Gransak.ToSql("first_name_and_last_name_cont", "cone")
    //returns: first_name LIKE '%cone%' AND last_name LIKE '%cone%'
