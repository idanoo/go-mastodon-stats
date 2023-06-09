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
		ClientID:     MASTODON_CLIENT_ID,
		ClientSecret: MASTODON_CLIENT_SECRET,
	})

	err := c.Authenticate(context.Background(), MASTODON_CLIENT_USERNAME, MASTODON_CLIENT_PASSWORD)
	if err != nil {
		log.Fatal(err)
	}

	// Build status
	startOfDay := getStartofDay()
	msg := fmt.Sprintf("Weekly user counts @ %d-%d-%d :laserkiwi:\n\n", startOfDay.Year(), startOfDay.Month(), startOfDay.Day())
	for _, m := range metrics {
		msg = msg + getPrintableString(m, true) + "\n"
	}
	msg = msg + "\n\n" + "#WeeklyStats"

	toot := &mastodon.Toot{
		Visibility: "unlisted",
		Status:     msg,
	}

	_, err = c.PostStatus(context.Background(), toot)
	if err != nil {
		log.Print(err)
	}
}
