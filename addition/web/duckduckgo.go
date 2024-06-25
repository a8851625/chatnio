package web

import (
	"chat/channel"
	"chat/utils"
	"fmt"
	"net/url"
	"strings"
)

type DDGResponse struct {
	Results []struct {
		Body  string `json:"body"`
		Href  string `json:"href"`
		Title string `json:"title"`
	} `json:"results"`
}

func duckDuckGoResponse(data *DDGResponse) string {
	res := make([]string, 0)
	for _, item := range data.Results {
		if item.Body == "" || item.Href == "" || item.Title == "" {
			continue
		}

		res = append(res, fmt.Sprintf("%s (%s): %s", item.Title, item.Href, item.Body))
	}

	return strings.Join(res, "\n")
}

func CallDuckDuckGoAPI(query string) (*DDGResponse, error) {
	data, err := utils.Get(fmt.Sprintf(
		"%s/search?q=%s&max_results=%d",
		channel.SystemInstance.GetSearchEndpoint(),
		url.QueryEscape(query),
		channel.SystemInstance.GetSearchQuery(),
	), nil)

	if err != nil {
		return nil, err
	}

	return utils.MapToRawStruct[DDGResponse](data)
}
