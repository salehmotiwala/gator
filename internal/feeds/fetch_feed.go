package feeds

import (
	"context"
	"encoding/xml"
	"html"
	"net/http"
)

func FetchFeed(feedURL string) (*RSSFeed, error) {
	feed := &RSSFeed{}

	req, err := http.NewRequestWithContext(context.Background(), "GET", feedURL, nil)
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

	return feed, nil
}
