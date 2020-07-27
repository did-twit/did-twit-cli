package tweet

import (
	"github.com/did-twitter/did-twitter-cli/internal/lib/tweet"
)

type TweetAPI interface {
	PostTweet(tweet tweet.Tweet) (string, error)
	GetTweet(id string) error
}
