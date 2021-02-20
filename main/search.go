package main

import (
	"GithubSearch"
	"encoding/base64"
	"encoding/json"
	"flag"
	_ "flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {

	path := flag.String("path", "", "search files inside this path")
	infile := flag.Bool("infile", false, "search inside file content")
	inpath := flag.Bool("inpath", false, "search inside the path string")
	lang := flag.String("lang", "", "search only this language")
	query := flag.String("query", "", "main query string")
	extension := flag.String("extension", "", "search only inside files with this extension")

	flag.Parse()

	if *query == "" {
		print("--query is a mandatory param")
		os.Exit(1)
	}

	q := GithubSearch.QueryBuilder{
		Language:  *lang,
		Path:      *path,
		Extension: *extension,
		InFile:    *infile,
		InPath:    *inpath,
		QueryStr:  GithubSearch.Query(*query),
	}

	token := os.Getenv("GOTOKEN")
	client := &http.Client{}
	log.Println("request url is " + "" + q.BuildQuery())
	req, _ := http.NewRequest("GET", "https://api.github.com/search/code?q="+q.BuildQuery(), nil)
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", "token "+token)
	res, err := client.Do(req)

	if err != nil {
		println("failed to communicate with the server: " + err.Error())
		os.Exit(1)
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		log.Println("request failed with status code " + strconv.Itoa(res.StatusCode))
		body, _ := ioutil.ReadAll(res.Body)
		log.Print(string(body))
		os.Exit(1)
	}

	all, _ := ioutil.ReadAll(res.Body)
	var val map[string]interface{}
	err = json.Unmarshal(all, &val)
	if err != nil {
		log.Println("err is " + err.Error())
		os.Exit(1)
	}

	var items []interface{}
	items = val["items"].([]interface{})

	//var buff map[string]interface{}
	for _, v := range items {

		item := v.(map[string]interface{})

		textMatches := item["text_matches"].([]interface{})
		for _, v := range textMatches {
			textMatch := v.(map[string]interface{})
			fragment := textMatch["fragment"].(string)
			objecturl := textMatch["object_url"].(string)
			log.Println("fragment:" + fragment)
			log.Println("objecturl:" + objecturl)
			log.Println("----------------------------------")
		}

	}

	//printAllContent(items, token, client)

	//curl -H  'Accept: application/vnd.github.v3.text-match+json' 'https://api.github.com/search/code?utf8=%E2%9C%93&q=octocat+filename:readme+path:%2F&access_token=a5fc00bdbc386349362ddac9b874e465d9516fb0'

	//https://github.com/search?q=user:github+extension:rb&type=Code

	//https://github.com/search?utf8=%E2%9C%93&q=org:github+extension:js&type=Code
	//https://github.com/search?q=repo:mozilla%2Fshumway+extension:as&type=Code
	//https://github.com/search?utf8=%E2%9C%93&q=octocat+filename:readme+path:%2F&type=Code
	//https://github.com/search?q=form+path:cgi-bin+language:perl&type=Code
	//https://github.com/search?q=console+path:app/public%22+language:javascript&type=Code
	//{
	//		totocal_coount : 23423,
	//			incomplete_results: 2325,
	//			items:[
	//				{
	//					name: "string",
	//					sha: dsfsdf,
	//					url:sdasd,
	//					html_url:sdfsd,
	//					repository:sdfsd,
	//
	//
	//					score:sd,
	//					text_matches:[
	//						{
	//							object_url:sfd,
	//							object_type:sdfsdf,
	//							property:edfm
	//							framegment:dfsdf
	//
	//						}
	//					]
	//				}
	//
	//
	//			]
	//
	//
	//}
}

func printAllContent(items []interface{}, token string, client *http.Client) {
	var buff map[string]interface{}
	for _, v := range items {

		mpVal := v.(map[string]interface{})

		req, _ := http.NewRequest("GET", (mpVal["url"]).(string), nil)
		req.Header.Set("Authorization", "token "+token)
		resp, _ := client.Do(req)
		respBody, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(respBody, &buff)
		decodeString, _ := base64.StdEncoding.DecodeString(buff["content"].(string))
		log.Println(string(decodeString))
	}
}
