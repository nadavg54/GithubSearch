package GithubSearch

import (
	"log"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

type Query string

type QueryBuilder struct {
	Language  string
	Path      string
	Extension string
	InFile    bool
	InPath    bool
	QueryStr  Query
}

var multiNamesQualifiers = map[string]string{
	"InPath": "In:Path",
	"InFile": "In:File",
}

func (qb *QueryBuilder) BuildQuery() string {

	var res string
	qbVal := reflect.ValueOf(qb).Elem()
	qbType := reflect.ValueOf(qb).Elem().Type()
	re := regexp.MustCompile("\\s+")
	for i := 0; i < qbType.NumField(); i++ {
		currType := qbType.Field(i)
		currVal := qbVal.Field(i)

		qualifierName := currType.Name
		mappedQualifierName := multiNamesQualifiers[currType.Name]
		if mappedQualifierName != "" {
			qualifierName = mappedQualifierName
		}

		typeName := currType.Type.Name()

		switch typeName {
		case "bool":
			{
				b := (currVal.Interface()).(bool)
				if b {
					res += " " + qualifierName
				}
			}
		case "string":
			{
				str := (currVal.Interface()).(string)
				if str != "" {
					if re.MatchString(str) {
						str = "\"" + str + "\""
					}
					res += " " + qualifierName + ":" + str

				}
			}
		case "Query":
			{
				strval := string(currVal.Interface().(Query))
				if re.MatchString(strval) {
					strval = "\"" + strval + "\""
				}
				res = strval + res
			}
		}
	}
	log.Println("res is " + res)
	return url.QueryEscape(strings.ToLower(res))
}

//https://github.com/search?q=user:github+extension:rb&type=Code
//https://github.com/search?utf8=%E2%9C%93&q=org:github+extension:js&type=Code
//https://github.com/search?q=repo:mozilla%2Fshumway+extension:as&type=Code
//https://github.com/search?utf8=%E2%9C%93&q=octocat+filename:readme+path:%2F&type=Code
//https://github.com/search?q=form+path:cgi-bin+language:perl&type=Code
//https://github.com/search?q=console+path:app/public%22+language:javascript&type=Code
