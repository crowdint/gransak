package gransak

//import (
//"net/http"
//"net/url"
//"regexp"
//"strconv"
//"strings"
//)

//func parseRequest(r *http.Request) string {
//params := r.URL.Query()

//return getGransakQuery(&params)
//}

//func parseUrlValues(params url.Values) string {
//return getGransakQuery(&params)
//}

//func getGransakQuery(params *url.Values) string {
//r := regexp.MustCompile(`^q\[[\w]+\]$`)
//var temp, sql string
//statements := []string{}

//for key, value := range *params {

//if r.MatchString(key) {
//temp = strings.Replace(key, "q[", "", 1)
//temp = strings.Replace(temp, "]", "", 1)

//sql, _ := getSqlString(temp, value[0])

//if strings.Trim(sql, " ") != "" {
//statements = append(statements, sql)
//}
//}
//}

//return strings.Join(statements, " AND ")
//}

//func getSqlString(query, value string) (string, []interface{}) {

//if intVal, err := strconv.ParseInt(value, 0, 64); err == nil {
//return Gransak.ToSql(query, intVal)
//}

//if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
//return Gransak.ToSql(query, floatVal)
//}

//return Gransak.ToSql(query, value)
//}
