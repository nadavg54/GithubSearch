package GithubSearch

import (
	"log"
)

type ContentFetcher struct {
	Client *AuthenticatedClient
}

func (cf *ContentFetcher) fetch(urls <-chan string, output chan<- string) {
	for url := range urls {
		content, err := cf.Client.Get(url)
		if err != nil {
			log.Println("can't fetch content from " + url + " with error " + err.Error())
		}
		output <- content
	}
}
