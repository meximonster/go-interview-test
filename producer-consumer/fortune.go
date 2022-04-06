package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func Produce(stream *TweetStream) []*Tweet {
	var tweets []*Tweet
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			break
		}
		tweets = append(tweets, tweet)
	}
	return tweets
}

func Consume(tweets []*Tweet) []*Tweet {
	var result []*Tweet
	for _, t := range tweets {
		if t.IsTalkingAboutFortune() {
			fmt.Printf("%s is talking about fortune!\n", t.Username)
			result = append(result, t)
		} else {
			fmt.Printf("%s is NOT talking about fortune.\n", t.Username)
		}
	}
	return result
}

func main() {
	start := time.Now()
	stream := InitStream()
	tweets := Produce(stream)
	result := Consume(tweets)
	fmt.Printf("finished after %v. %d people are talking about fortune.", time.Since(start), len(result))
}

// try not to edit the code below.

type Tweet struct {
	Username string
	Content  string
}

func (t *Tweet) IsTalkingAboutFortune() bool {
	// let's assume that this function is a very sophisticated machine learning procedure, which takes 0.3 sec
	time.Sleep(300 * time.Millisecond)
	if strings.Contains(strings.ToLower(t.Content), "fortune") {
		return true
	}
	return false
}

type TweetStream struct {
	pos    int
	tweets []Tweet
}

func (t *TweetStream) Next() (*Tweet, error) {
	// assuming that the tweet needs 0.2 seconds to be retrieved
	time.Sleep(200 * time.Millisecond)
	if t.pos >= len(t.tweets) {
		return &Tweet{}, ErrEOF
	}

	tweet := t.tweets[t.pos]
	t.pos++

	return &tweet, nil
}

var ErrEOF = errors.New("end of file")

func InitStream() *TweetStream {
	return &TweetStream{
		pos: 0,
		tweets: []Tweet{
			{
				Username: "Dylan Thomas",
				Content:  "When one burns one's bridges, what a very nice fire it makes.",
			}, {
				Username: "William Shakespeare",
				Content:  "Oh, I am fortune's fool!",
			}, {
				Username: "Samuel Beckett",
				Content: "All I know is what the words know, and dead things, and that makes a handsome little " +
					"sum, with a beginning and a middle and an end, as in the well-built phrase and the long " +
					"sonata of the dead.",
			}, {
				Username: "Mahatma Ghandi",
				Content:  "Seek not greater wealth, but simpler pleasure; not higher fortune, but deeper felicity.",
			}, {
				Username: "Mario Puzo",
				Content:  "Behind every successful fortune there is a crime.",
			}, {
				Username: "Mark Twain",
				Content:  "The secret source of humor is not joy but sorrow; there is no humor in Heaven.",
			}, {
				Username: "Christopher Marlowe",
				Content:  "I'll burn my books.",
			}, {
				Username: "Samuel Johnson",
				Content: "Your manuscript is both good and original, but the part that is good is not original " +
					"and the part that is original is not good...",
			}, {
				Username: "Miguel de Cervantes Saavedra",
				Content:  "Diligence is the mother of good fortune",
			}, {
				Username: "J. R. R. Tolkien",
				Content:  "\"Elves and Dragons!\" I says to him.  \"Cabbages and potatoes are better for you and me.\"",
			},
		},
	}
}
