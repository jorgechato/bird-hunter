package main

import (
	"fmt"
	"sync"

	"github.com/ahmdrz/goinsta/response"
	"github.com/jorgechato/birdhunter/birdhunter"
	"gopkg.in/ahmdrz/goinsta.v1"
)

var (
	wg      sync.WaitGroup
	mu      sync.Mutex // guards balance
	balance = 1
	liked   = []string{}
)

func (c *Client) login() {
	insta = goinsta.New(c.User, c.Password)

	if err := insta.Login(); err != nil {
		panic(err)
	}

	insta.SyncFeatures()
}

func (c *Client) getTagIds() {
	for _, tag := range c.Tags {
		wg.Add(2)
		media, _ := insta.TagFeed(tag)

		go likeMedia(media.RankedItems)
		go likeMedia(media.FeedsResponse.Items)
	}

	wg.Wait()
}

func (c *Client) getPopularIds() {
	media, _ := insta.GetPopularFeed()

	likePopularMedia(media.Items)
}

func likePopularMedia(list []response.Item) {
	for _, item := range list {
		if !item.HasLiked && item.LikeCount > likes.Minimum {
			if balance <= likes.NLikes {
				liked = append(liked, birdhunter.Slug(item.Code))
				insta.Like(item.ID)
				balance++
			}
		}
	}
}

func likeMedia(list []response.MediaItemResponse) {
	defer wg.Done()

	for _, item := range list {
		if !item.HasLiked && item.LikeCount > likes.Minimum {
			mu.Lock()
			if balance <= likes.NLikes {
				liked = append(liked, birdhunter.Slug(item.Code))
				insta.Like(item.ID)
				balance++
			}
			mu.Unlock()
		}
	}
}

func printOut() {
	fmt.Printf("\nTotal likes: %d\n\n", len(liked))
	for _, item := range liked {
		fmt.Println(item)
	}
}
