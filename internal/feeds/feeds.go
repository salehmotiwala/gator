package feeds

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	feed := &RSSFeed{}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return feed, err
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return feed, err
	}

	defer res.Body.Close()

	decoder := xml.NewDecoder(res.Body)
	err = decoder.Decode(feed)

	if err != nil {
		return feed, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := range feed.Channel.Item {
		item := &feed.Channel.Item[i]
		title := item.Title
		desc := item.Description

		item.Title = html.UnescapeString(title)
		item.Description = html.UnescapeString(desc)

	}

	fmt.Println("%v", feed)

	return feed, nil
}
