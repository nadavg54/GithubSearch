package GithubSearch

import (
	"net/url"
	"strings"
	"testing"
)

func TestQueryBuilder_buildQuery(t *testing.T) {
	builder := QueryBuilder{
		Lang:      "javascript",
		Path:      "app/public",
		Extension: "",
		InFile:    true,
		InPath:    false,
	}

	want := "+lang:javascript+path:app/public+in:file"
	got, _ := url.QueryUnescape(builder.buildQuery())
	got = strings.ToLower(got)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

//https://github.com/search?q=user:github+extension:rb&type=Code
//https://github.com/search?utf8=%E2%9C%93&q=org:github+extension:js&type=Code
//https://github.com/search?q=repo:mozilla%2Fshumway+extension:as&type=Code
//https://github.com/search?utf8=%E2%9C%93&q=octocat+filename:readme+path:%2F&type=Code
//https://github.com/search?q=form+path:cgi-bin+language:perl&type=Code
//https://github.com/search?q=console+path:app/public%22+language:javascript&type=Code
