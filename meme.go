package gomemes

// Meme
type Meme struct {
	ID         string  `json:"id"`
	Subreddit  string  `json:"subreddit"`
	Title      string  `json:"title"`
	Author     string  `json:"author"`
	Image      string  `json:"image"`
	Ups        int     `json:"ups"`
	Downs      int     `json:"downs"`
	Score      int     `json:"score"`
	Comments   int     `json:"comments"`
	NSFW       bool    `json:"nsfw"`
	CreatedUTC float64 `json:"createdUtc"`
}

var Subreddits = []string{
	"funny",
	"memes",
	"meirl",
	"valorantmemes",
	"dankmemes",
	"wholesomememes",
	"okbuddyretard",
	"comedymemes",
	"pewdiepiesubmissions",
	"lastimages",
	"historymemes",
	"raimimemes",
	"Overwatch_Memes",
}
