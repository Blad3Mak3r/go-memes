package gomemes

import "testing"

func TestGetRandomMemeFromSubreddit(t *testing.T) {
	meme, err := getMeme("memes")
	if err != nil {
		t.Error(err)
	} else if err == nil && meme == nil {
		t.Errorf("no error and no meme...")
	} else if meme == nil {
		t.Errorf("no meme")
	} else {
		t.Logf("success")
	}
}
