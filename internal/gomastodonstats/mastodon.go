package gomastodonstats

import (
	"context"
	"fmt"
	"log"

	"github.com/mattn/go-mastodon"
)

func postToMastodon(metrics []metric) {
	if MASTODON_INSTANCE_URL == "" {
		log.Printf("Skipping posting to mastodon. Missing configuration")
		return
	}

	c := mastodon.NewClient(&mastodon.Config{
		Server:       fmt.Sprintf("https://%s", MASTODON_INSTANCE_URL),
		ClientID:     "go-mastodon-stats",
		ClientSecret: "randomgeneratedstringthatimnotsureisneeded",
	})
	err := c.Authenticate(context.Background(), MASTODON_USERNAME, MASTODON_PASSWORD)
	if err != nil {
		log.Println("Invalid mastodon credentials")
		return
	}

	// Build status
	startOfDay := getStartofDay()
	msg := fmt.Sprintf("Stats for %s\n\n", startOfDay.String())
	for _, m := range metrics {
		msg = msg + getPrintableString(m) + "\n"
	}
	msg = msg + "\n\n" + "#Stats"

	toot := mastodon.Toot{
		Status: msg,
	}

	_, err = c.PostStatus(context.Background(), &toot)
	if err != nil {
		log.Print(err)
	}
}
