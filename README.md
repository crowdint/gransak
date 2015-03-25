# Gransak

###Ransack to SQL, parser for Golang

Gransak was born for the need of a replacement for ransack. This was because we needed a way to keep the searching functionality of a previous Rails app we were migrating to Golang and we found no similar library.

So we decided to create a library to transform a ransak like string to a Sql statement an be able to use it in a normal query to a database using golang. We also found useful to generate only the 'WERE' part of the query to be able to use it with ``gorm`` (https://github.com/jinzhu/gorm) e.g.

        query, params := Gransak.ToSql(ransakQuery)
        db.Were(query, params).Find(&users)
        
or in the traditional way:
        
        db.Query(quey, params...)

##Install

Download the repository with:
``go get github.com/crowdint/gransak``

Then include the project inside your project like this:
``import "github.com/crowdint/gransak"``

##What it does

Currently gransak transforms a ransack like string into a sql 'where' statement e.g.

    package main

    import (
      . "github.com/crowdint/gransak"
      "fmt"
    )
    
    func main() {
      sql, params := Gransak.ToSql("user_name_eq", "cone")

      fmt.Printf("query-> %s, params-> %v", sql, params)
      //prints: query-> user_name = ?, params-> [cone]
    }
    
Also it can generate the complete statement if a table name is specified
e.g.

    sql, _ = Gransak.Table("users").ToSql("user_name_eq", "cone")
    //returns: SELECT * FROM users WHERE user_name = ?
    //parameters: [cone]
    
##Methods

###ToSql

Returns an SQL statement. It takes the ransak query string and the value as parameters. 

    sql, params = Gransak.ToSql("user_name_eq", "cone")
    //returns: user_name = ?
    //parameters: [cone]
 
###FromRequest

Resturns an SQL statement. It gets the query strings from a http.Request struct.

    func Handler(w http.ResponseWritter, r *http.Request){
        //request: http://someurl/params?q[user_name_eq]=cone
        sql, _ = Gransak.FromRequest(r)
        //returns: user_name = ?
        //parameters: [cone]
    }
    
We can chain several statements too.

    func Handler(w http.ResponseWritter, r *http.Request){
        //request: http://someurl/params?q[user_name_eq]=cone&q[role_cont]=admin
        sql, _ = Gransak.FromRequest(r)
        //returns: user_name = ? AND role LIKE ?
        //parameters: [cone %admin%]
    }

###FromUrlValues

Returns an SQL statement. It gets the query strings from an url.Values struct.

    func Handler(w http.ResponseWritter, r *http.Request){
        //request: http://someurl/params?q[user_name_eq]=cone
        values := r.URL.Query()
        sql, _ = Gransak.FromUrlValues(values)
        //returns: user_name = ?
        //parameters: [cone]
    }

###SetEngine

Changes the type of placeholder used e.g. "?" for MySQL or "$[n]" for
PostgreSql.

At this moment Gransak doesn't support associations

##Searching operations currently supported

In the following examples the placeholders have been substituted by the
parameter value for a better visualization.

###or

    Gransak.ToSql("first_name_or_last_name_cont", "cone")
    //returns: first_name LIKE '%cone%' OR last_name LIKE '%cone%' 
    
###and

    Gransak.ToSql("first_name_and_last_name_cont", "cone")
    //returns: first_name LIKE '%cone%' AND last_name LIKE '%cone%'

###cont (and its opposite 'not_cont')

    Gransak.ToSql("first_name_cont", "cone")
    //returns: first_name LIKE '%cone%'

###matches (and its opposite 'does_not_match')

    Gransak.ToSql("first_name_matches", "cone")
    //returns: first_name LIKE 'cone'

###lt (and its opposite 'gt')

    Gransak.ToSql("age_lt", 30)
    //returns: age < 30

###lteq (and its opposite 'gteq')

    Gransak.ToSql("age_lteq", 30)
    //returns: age <= 30

###eq (and its opposite 'not_eq')

    Gransak.ToSql("first_name_eq", "cone")
    //returns: first_name = 'cone'

    Gransak.ToSql("age_eq", 29)
    //returns: age = 29

###start (and its opposite 'not_start')

    Gransak.ToSql("first_name_start", "co")
    //returns: first_name LIKE 'co%'

###end (and its opposite 'not_end')

    Gransak.ToSql("first_name_end", "ne")
    //returns: first_name LIKE '%ne'

###true (and its opposite 'not_true')

    Gransak.ToSql("is_available", "1")
    //returns: is_available = 't'

###false (and its opposite 'not_false')

    Gransak.ToSql("is_available", "1")
    //returns: is_available = 'f'

###present (and its opposite 'blank')

    Gransak.ToSql("first_name_present", "1")
    //returns: first_name IS NOT NULL AND first_name <> ''

###null (and its opposite 'not_null')

    Gransak.ToSql("first_name_null", "1")
    //returns: first_name IS NULL

###in (and its opposite 'not_in')

    Gransak.ToSql("age_in", "27..30")
    //returns: age IN (27,28,29,30)

    Gransak.ToSql("age_in", "[27,28,29,30]")
    //returns: age IN (27,28,29,30)

    Gransak.ToSql("age_in", []int{27,28,29,30})
    //returns: age IN (27,28,29,30)

###cont_any (and its opposite 'not_cont_any')

    Gransak.ToSql("user_role_cont_any", "%w(admin developer)")
    //returns: user_role LIKE '%admin%' OR user_role LIKE '%developer%'

## Contributing

1. Create your feature branch (`git checkout -b feature/my-new-feature`)
2. Commit your changes (`git commit -am 'Add some feature'`)
3. Push to the branch (`git push origin feature/my-new-feature`)
4. Create a new Pull Request

## About the Author

[Crowd Interactive](http://www.crowdint.com) is an American web design and development company that happens to work in Colima, Mexico.
We specialize in building and growing online retail stores. We don’t work with everyone – just companies we believe in. Call us today to see if there’s a fit.
