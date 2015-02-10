# Gransak

###Ransack to SQL, parser for Golang

##Install

Download the repository with:
``go get github.com/crowdint/gransak/filter``

Then include the project inside your project like this:
``import "github.com/crowdint/gransak/filter"``

##What it does

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
    
Also it can generate the complete statement if a table name is specified
e.g.

    sql = Gransak.Table("users").ToSql("user_name_eq", "cone")
    //returns: SELECT * FROM users WHERE user_name = 'cone'
    
##Methods

###ToSql

Returns an SQL statement. It takes the ransak query string and the value as parameters 

    sql = Gransak.ToSql("user_name_eq", "cone")
    //returns: user_name = 'cone'
 
###FromRequest

Resturns an SQL statement. It gets the query strings from a http.Request struct

    func Handler(w http.ResponseWritter, r *http.Request){
        //request: http://someurl/params?q[user_name_eq]=cone
        sql = Gransak.FromRequest(r)
        //returns: user_name = 'cone'
    }
    
We can chain several statements too

    func Handler(w http.ResponseWritter, r *http.Request){
        //request: http://someurl/params?q[user_name_eq]=cone&q[role_cont]=admin
        sql = Gransak.FromRequest(r)
        //returns: user_name = 'cone' AND role LIKE %admin%
    }

###FromUrlValues

Returns an SQL statement. It gets the query strings from an url.Values struct

    func Handler(w http.ResponseWritter, r *http.Request){
        //request: http://someurl/params?q[user_name_eq]=cone
        values := r.URL.Query()
        sql = Gransak.FromUrlValues(values)
        //returns: user_name = 'cone'
    }

At this moment Gransak doesn't support associations

##Searching operations currently supported

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

###cont_any (and its opposite 'not_cont_any'

    Gransak.ToSql("user_role_cont_any", "%w(admin developer)")
    //returns: user_role LIKE '%admin%' OR user_role LIKE '%developer%'

###we can even to things like this

    Gransak.ToSql("name_cont_and_role_eq", "%w(cone developer)")
    //returns: name LIKE '%cone%' and role = 'developer'

## Contributing

1. Create your feature branch (`git checkout -b feature/my-new-feature`)
2. Commit your changes (`git commit -am 'Add some feature'`)
3. Push to the branch (`git push origin feature/my-new-feature`)
4. Create a new Pull Request
