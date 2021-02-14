package GithubSearch

import (
	"net/url"
	"reflect"
)

type QueryBuilder struct {
	Lang      string
	Path      string
	Extension string
	InFile    bool
	InPath    bool
}

var multiNamesQualifiers = map[string]string{
	"InPath": "In:Path",
	"InFile": "In:File",
}

func (qb *QueryBuilder) buildQuery() string {

	var res string
	qbVal := reflect.ValueOf(qb).Elem()
	qbType := reflect.ValueOf(qb).Elem().Type()

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
					res += "+" + qualifierName
				}
			}
		case "string":
			{
				str := (currVal.Interface()).(string)
				if str != "" {
					res += "+" + qualifierName + ":" + str
				}
			}
		}
	}
	return url.QueryEscape(res)
}

//https://github.com/search?q=user:github+extension:rb&type=Code
//https://github.com/search?utf8=%E2%9C%93&q=org:github+extension:js&type=Code
//https://github.com/search?q=repo:mozilla%2Fshumway+extension:as&type=Code
//https://github.com/search?utf8=%E2%9C%93&q=octocat+filename:readme+path:%2F&type=Code
//https://github.com/search?q=form+path:cgi-bin+language:perl&type=Code
//https://github.com/search?q=console+path:app/public%22+language:javascript&type=Code
