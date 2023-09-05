package imagesearch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var client = http.Client{}

type Image struct {
	Height         int    `json:"height"`
	Image          string `json:"image"`
	ImageToken     string `json:"image_token"`
	Source         string `json:"source"`
	Thumbnail      string `json:"thumbnail"`
	ThumbnailToken string `json:"thumbnail_token"`
	Title          string `json:"title"`
	Url            string `json:"url"`
	Width          int    `json:"width"`
}

const URL = "https://duckduckgo.com/i.js?&o=json&vqd=4-212188280713482655434498307572246496973&f=,,,,,&q=%s&p=%d"

type DDGImages struct {
	Results []Image `json:"results"`
}

func Urls(q string, p uint16) (urls []string, err error) {
	u := fmt.Sprintf(URL, url.QueryEscape(q), p)
	get, err := client.Get(u)
	if err != nil {
		return
	}
	defer get.Body.Close()
	all, err := io.ReadAll(get.Body)
	if err != nil {
		return
	}
	result := new(DDGImages)
	json.Unmarshal(all, result)
	urls = make([]string, len(result.Results))
	for i, s := range result.Results {
		urls[i] = s.Image
	}
	return
}
