package gomastodonstats

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mattn/go-mastodon"
)

func doTheRegisterThing() {
	srv, err := mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
		Server:       fmt.Sprintf("https://%s", MASTODON_INSTANCE_URL),
		ClientName:   MASTODON_CLIENT_NAME,
		Scopes:       "read write follow",
		RedirectURIs: "urn:ietf:wg:oauth:2.0:oob",
		Website:      "https://github.com/mattn/go-mastodon",
	})
	if err != nil {
		log.Fatal("Couldn't register the app. Error: %v\n\nExiting...\n", err)
	}

	log.Println(
		"Pls open this link: %s\n\n" + srv.AuthURI +
			"When completeAdd this to your config under MASTODON_TOKEN and rerun the app")
	os.Exit(0)
}

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
	msg := fmt.Sprintf("User counts for %s :laserkiwi:\n\n", startOfDay.String())
	for _, m := range metrics {
		msg = msg + getPrintableString(m) + "\n"
	}
	msg = msg + "\n\n" + "#DailyStats"

	toot := &mastodon.Toot{
		Status: msg,
	}

	_, err = c.PostStatus(context.Background(), toot)
	if err != nil {
		log.Print(err)
	}
}
