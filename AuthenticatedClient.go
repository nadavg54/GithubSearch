package GithubSearch

import (
	"io/ioutil"
	"net/http"
)

type AuthenticatedClient struct {
	Client http.Client
	Token  string
}

func (ac *AuthenticatedClient) Get(url string) (string, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token "+ac.Token)
	response, err := ac.Client.Do(req)
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
