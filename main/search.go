package main

import (
	"encoding/json"
	_ "flag"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {

	//var typeflag = flag.String("type", "Code", "resource to search: Topics,Issues,PullRequests,Discussions,Code,Commits,Users,Packages,Wikis")
	//var keywork = flag.String("text","","text to find")
	//var langauge = flag.String("lang","java","search only this language")

	token := os.Getenv("GOTOKEN")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/search/code?q=\"github+api\"+language:go", nil)
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", "token "+token)
	res, err := client.Do(req)

	if err != nil {
		println("error is " + err.Error())
		os.Exit(1)
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		print("status code is " + strconv.Itoa(res.StatusCode))
		os.Exit(1)
	}

	all, _ := ioutil.ReadAll(res.Body)
	var val map[string]interface{}
	err = json.Unmarshal(all, &val)
	if err != nil {
		println("err is " + err.Error())
		os.Exit(1)
	}

	var items1 []interface{}
	items1 = val["items"].([]interface{})

	for _, v := range items1 {

		mpVal := v.(map[string]interface{})
		println(mpVal["sha"].(string))
	}

	//curl -H  'Accept: application/vnd.github.v3.text-match+json' 'https://api.github.com/search/code?utf8=%E2%9C%93&q=octocat+filename:readme+path:%2F&access_token=a5fc00bdbc386349362ddac9b874e465d9516fb0'

	//https://github.com/search?q=user:github+extension:rb&type=Code
	//https://github.com/search?utf8=%E2%9C%93&q=org:github+extension:js&type=Code
	//https://github.com/search?q=repo:mozilla%2Fshumway+extension:as&type=Code
	//https://github.com/search?utf8=%E2%9C%93&q=octocat+filename:readme+path:%2F&type=Code
	//https://github.com/search?q=form+path:cgi-bin+language:perl&type=Code
	//https://github.com/search?q=console+path:app/public%22+language:javascript&type=Code

}
