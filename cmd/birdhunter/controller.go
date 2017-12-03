package main

import (
	"fmt"
	"sync"

	"github.com/ahmdrz/goinsta/response"
	"gopkg.in/ahmdrz/goinsta.v1"
)

var (
	wg      sync.WaitGroup
	mu      sync.Mutex // guards balance
	balance int
)

func (c *Client) login() {
	insta = goinsta.New(c.User, c.Password)

	if err := insta.Login(); err != nil {
		panic(err)
	}

	insta.SyncFeatures()
}

func (c *Client) getTagIds() {
	client.login()
	balance = 0

	wg.Add(len(c.Tags))

	for _, tag := range c.Tags {
		media, _ := insta.TagFeed(tag)

		go likeMedia(media.RankedItems)
		go likeMedia(media.FeedsResponse.Items)
	}

	wg.Wait()
}

func likeMedia(list []response.MediaItemResponse) {
	defer wg.Done()

	for _, item := range list {
		if !item.HasLiked && item.LikeCount > likes.Minimum {
			mu.Lock()
			balance++
			if balance <= likes.NLikes {
				fmt.Println(item.Code)
				insta.Like(item.ID)
			} else {
				break
			}
			mu.Unlock()
		}
	}
}
