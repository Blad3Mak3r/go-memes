package gomemes

import (
	"fmt"
	"testing"
)

func TestGetRandomMemeFromSubreddit(t *testing.T) {
	meme, err := GetRandomMemeFromSubreddit("memes")
	if err != nil {
		t.Error(err)
	} else if err == nil && meme == nil {
		t.Errorf("no error and no meme...")
	} else if meme == nil {
		t.Errorf("no meme")
	} else {
		fmt.Printf("%+v\n", meme)
		t.Logf("success")
	}
}

func TestGetRandomMeme(t *testing.T) {
	meme, err := GetRandomMeme()
	if err != nil {
		t.Error(err)
	} else if err == nil && meme == nil {
		t.Errorf("no error and no meme...")
	} else if meme == nil {
		t.Errorf("no meme")
	} else {
		fmt.Printf("%+v\n", meme)
		t.Logf("success")
	}
}
