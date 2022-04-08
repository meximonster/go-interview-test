package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

var workers = 10
var result []*Tweet

func Produce(stream *TweetStream, c chan *Tweet) {
	defer close(c)
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			break
		}
		c <- tweet
	}
}

func Consume(c chan *Tweet, wg *sync.WaitGroup) {
	var mtx sync.Mutex
	defer wg.Done()
	for t := range c {
		if t.IsTalkingAboutFortune() {
			fmt.Printf("%s is talking about fortune!\n", t.Username)
			mtx.Lock()
			result = append(result, t)
			mtx.Unlock()
		} else {
			fmt.Printf("%s is NOT talking about fortune.\n", t.Username)
		}
	}
}

func main() {
	start := time.Now()
	stream := InitStream()
	c := make(chan *Tweet)
	wg := sync.WaitGroup{}
	go Produce(stream, c)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go Consume(c, &wg)
	}
	wg.Wait()
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
	return strings.Contains(strings.ToLower(t.Content), "fortune")
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
