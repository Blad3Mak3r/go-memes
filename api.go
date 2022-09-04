package gomemes

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

var (
	Version = "0.0.1"

	validExtensions = []string{
		".png",
		".jpg",
		".jpeg",
		".gif",
		".webp",
	}

	httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MaxVersion: tls.VersionTLS12,
			},
		},
	}
)

const (
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59"
)

type Posts struct {
	Kind string `json:"kind"`
	Data struct {
		Children []struct {
			Kind string   `json:"kind"`
			Data PostData `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type PostData struct {
	ID             string  `json:"id"`
	Subreddit      string  `json:"subreddit"`
	AuthorFullname string  `json:"author_fullname"`
	Title          string  `json:"title"`
	Downs          int     `json:"downs"`
	Ups            int     `json:"ups"`
	Score          int     `json:"score"`
	Over18         bool    `json:"over_18"`
	Permaling      string  `json:"permalink"`
	Url            string  `json:"url"`
	Comments       int     `json:"num_comments"`
	CreatedUTC     float64 `json:"created_utc"`
}

// GetRandomMemeFromSubreddit retrieve a random meme from the given subreddit
func GetRandomMemeFromSubreddit(subreddit string) (*Meme, error) {
	return getMeme(subreddit)
}

// GetRandomMeme retrieve a random meme from a random subreddit
func GetRandomMeme() (*Meme, error) {
	return getMeme(Subreddits[rand.Intn(len(Subreddits))])
}

func getMeme(subreddit string) (*Meme, error) {
	url := buildUrl(subreddit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("user-agent", userAgent)
	req.Header.Set("accept", "application/json")

	fmt.Printf("user-agent is %s\n", req.UserAgent())

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("request returned status: %s", res.Status)
	}

	defer res.Body.Close()

	var posts Posts
	if err := json.NewDecoder(res.Body).Decode(&posts); err != nil {
		return nil, err
	}

	validPosts := filterPosts(posts)

	if len(validPosts) < 1 {
		return nil, fmt.Errorf("not found memes in %s", subreddit)
	}

	meme := parsePostDataToMeme(validPosts[rand.Intn(len(validPosts))])

	return meme, nil
}

func buildUrl(subreddit string) string {
	return fmt.Sprintf("https://www.reddit.com/r/%s/hot/.json?count=100", subreddit)
}

func isPostWithImage(data PostData) bool {
	for _, ext := range validExtensions {
		if strings.HasSuffix(data.Url, ext) {
			return true
		}
	}
	return false
}

func filterPosts(posts Posts) []PostData {
	list := []PostData{}

	for _, post := range posts.Data.Children {
		if isPostWithImage(post.Data) {
			list = append(list, post.Data)
		}
	}

	return list
}

func parsePostDataToMeme(data PostData) *Meme {
	return &Meme{
		ID:         data.ID,
		Subreddit:  data.Subreddit,
		Title:      data.Title,
		Author:     data.AuthorFullname,
		Image:      data.Url,
		Score:      data.Score,
		Ups:        data.Ups,
		Downs:      data.Downs,
		Comments:   data.Comments,
		NSFW:       data.Over18,
		CreatedUTC: data.CreatedUTC,
	}
}
