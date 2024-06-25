package web

import (
	"chat/channel"
	"chat/utils"
	"fmt"
	"net/url"
	"strings"
)

type SearxngResponse struct {
	Query           string `json:"query"`
	NumberOfResults int    `json:"number_of_results"`
	Results         []struct {
		Url          string   `json:"url"`
		Title        string   `json:"title"`
		Content      string   `json:"content"`
		// Thumbnail    string   `json:"thumbnail"`
		// Engine       string   `json:"engine"`
		// ParsedUrl    []string `json:"parsed_url"`
		// Template     string   `json:"template"`
		// Engines      []string `json:"engines"`
		// Positions    []int    `json:"positions"`
		// PublishedDate string  `json:"PublishedDate"`
		// Score        float64  `json:"score"`
		// Category     string   `json:"category"`
	} `json:"results"`
}


func searxngResponse(data *SearxngResponse, queryNumber int) string {
	res := make([]string, 0)
	count := 0
	for _, item := range data.Results {
		if item.Content == "" || item.Url == "" || item.Title == "" {
			continue
		}
		if count >= queryNumber {
			break
		}
		count += 1
		res = append(res, fmt.Sprintf("%s (%s): %s", item.Title, item.Url, item.Content))
	}

	return strings.Join(res, "\n")
}

func CallSearxngAPI(query string) (*SearxngResponse, error) {
	data, err := utils.Get(fmt.Sprintf(
		"%s/search?q=%s&categories=%s&format=json&%s",
		channel.SystemInstance.GetSearchEndpoint(),
		url.QueryEscape(query),
		"page",
		channel.SystemInstance.GetSearchEngines(),
	), nil)

	if err != nil {
		return nil, err
	}

	return utils.MapToRawStruct[SearxngResponse](data)
}
