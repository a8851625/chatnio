package web

import (
	"chat/channel"
	"chat/globals"
	"chat/utils"
	"fmt"
	"net/url"
)

func GetBingUrl(q string) string {
	return "https://bing.com/search?q=" + url.QueryEscape(q)
}

func RequestWithUA(url string) string {
	data, err := utils.GetRaw(url, map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/116.0",
		"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
	})

	if err != nil {
		return ""
	}

	return data
}

func SearchReverse(q string) string {
	// deprecated
	uri := GetBingUrl(q)
	if res := CallPilotAPI(uri); res != nil {
		return utils.Marshal(res.Results)
	}
	data := RequestWithUA(uri)
	return ParseBing(data)
}

func SearchWebResult(q string) string {

	t := channel.SystemInstance.GetSearchType()
	globals.Debug(fmt.Sprintf("[web] search type : %s (query: %s)", t, q))

	var content = ""
	if t== "searxng" {
		res, err := CallSearxngAPI(q)
		if err != nil {
			globals.Warn(fmt.Sprintf("[web] failed to get search result: %s (query: %s)", err.Error(), q))
			return ""
		}
		c := channel.SystemInstance.GetSearchQuery()
		content = searxngResponse(res, c)
	}else{
		res, err := CallDuckDuckGoAPI(q)
		if err != nil {
			globals.Warn(fmt.Sprintf("[web] failed to get search result: %s (query: %s)", err.Error(), q))
			return ""
		}
		content = duckDuckGoResponse(res)
	}

	globals.Debug(fmt.Sprintf("[web] search result: %s (query: %s)", utils.Extract(content, 50, "..."), q))
	return content
}
