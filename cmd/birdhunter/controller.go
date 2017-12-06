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
)

func login() {
	log.Info("New login to Instagram")

	insta = goinsta.New(
		client.Username,
		client.Password,
	)

	if err := insta.Login(); err != nil {
		log.Panic(err)
	}

	insta.SyncFeatures()
}

func getTagIds() {
	balance = 1
	for _, tag := range client.Tags {
		wg.Add(2)
		media, _ := insta.TagFeed(tag)

		go likeMedia(media.RankedItems)
		go likeMedia(media.FeedsResponse.Items)
	}

	wg.Wait()
}

func getPopularIds() {
	media, _ := insta.GetPopularFeed()

	likePopularMedia(media.Items)
}

func likePopularMedia(list []response.Item) {
	for _, item := range list {
		if !item.HasLiked && item.LikeCount > likes.InPhoto {
			if balance <= likes.Number {
				insta.Like(item.ID)
				balance++
			}
		}
	}
}

func likeMedia(list []response.MediaItemResponse) {
	defer wg.Done()

	for _, item := range list {
		user, _ := insta.GetUserByID(item.User.ID)
		following := user.User.FollowingCount

		if !item.HasLiked && item.LikeCount > likes.InPhoto && following < 1000 {
			mu.Lock()
			if balance <= likes.Number {
				insta.Like(item.ID)
				balance++

				log.Info("Liked: ", birdhunter.Slug(item.Code))
			}
			mu.Unlock()
		}
	}
}

func updateConfig(url string) {
	res, _ := birdhunter.Client(
		birdhunter.HTTPClient{
			Url:         url,
			Method:      "POST",
			ContentType: birdhunter.Text_plain,
		},
		nil,
	)

	fmt.Print(res.Status)
}
