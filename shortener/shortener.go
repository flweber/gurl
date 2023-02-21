package shortener

import (
	"github.com/flweber/gurl/db"
	"github.com/go-redis/redis/v9"
	"math/rand"
	"net/url"
)

type Shortener struct {
	rdb         *db.Client
	letterBytes string
}

func New(rdb *db.Client) Shortener {
	return Shortener{rdb: rdb, letterBytes: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"}
}

func (s *Shortener) GetUrl(shortUrl string) (*string, error) {
	longUrl, err := s.rdb.Client.Get(s.rdb.Ctx, shortUrl).Result()
	return &longUrl, err
}

func (s *Shortener) Shorten(urlString string) (*string, error) {
	longUrl, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, err
	}
	var shortUrl string
	for {
		short := s.randStringBytes(10)
		_, err := s.GetUrl(short)
		if err != nil && err != redis.Nil {
			return nil, err
		}
		if err == redis.Nil {
			shortUrl = short
			break
		}
	}
	err = s.rdb.Client.Set(s.rdb.Ctx, shortUrl, longUrl.String(), 0).Err()
	if err != nil {
		return nil, err
	}

	return &shortUrl, nil
}

func (s *Shortener) randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = s.letterBytes[rand.Intn(len(s.letterBytes))]
	}
	return string(b)
}
