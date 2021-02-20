package GithubSearch

import (
	"net/url"
	"strings"
	"testing"
)

func TestQueryBuilder_buildQuery(t *testing.T) {
	builder := QueryBuilder{
		Language:  "java script",
		Path:      "app/public",
		Extension: "",
		InFile:    true,
		InPath:    false,
		QueryStr:  "my query",
	}

	want := strings.ToLower(url.QueryEscape("\"my query\" lang:\"java script\" path:app/public " +
		"" +
		"in:file"))
	got := strings.ToLower(builder.BuildQuery())

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
