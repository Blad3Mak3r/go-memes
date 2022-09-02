package gomemes

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type postChildren struct {
	Kind string `json:"kind"`
	Data struct {
		Subreddit      string `json:"subreddit"`
		AuthorFullname string `json:"author_fullname"`
		Title          string `json:"title"`
		Downs          int    `json:"downs"`
		Ups            int    `json:"ups"`
	} `json:"data"`
}

type post struct {
	Kind struct {
		Data struct {
			Children []postChildren `json:"children"`
		} `json:"data"`
	} `json:"kind"`
}

func GetRandomMemeFromSubreddit(subreddit string) (*Meme, error) {

}

func GetRandomMeme() (*Meme, error) {
	return getMeme(Subreddits[rand.Intn(len(Subreddits))])
}

func getMeme(subreddit string) (*Meme, error) {
	url := buildUrl(subreddit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var meme Meme
	if err := json.NewDecoder(req.Body).Decode(&meme); err != nil {
		return nil, err
	}

}

func buildUrl(subreddit string) string {
	return fmt.Sprintf("https://www.reddit.com/r/%s/hot/.json?count=100", subreddit)
}
