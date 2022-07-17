package client

import (
	"errors"
	"fampay/utils"
	"flag"
	"net/http"
	"time"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
	query      = flag.String("query", "cricket", "Search term")
	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
)

type YoutubeClient struct {
	HttpClient    http.Client
	DeveloperKey  string
	QuotaExceeded bool
}

var clients = []YoutubeClient{}

func InitYTClients(developerKeys []string) {
	for _, key := range developerKeys {
		client := &http.Client{
			Transport: &transport.APIKey{Key: key},
		}
		youtubeClient := YoutubeClient{
			HttpClient:    *client,
			DeveloperKey:  key,
			QuotaExceeded: false,
		}
		clients = append(clients, youtubeClient)
	}
}

func getYoutybeClient() (int, YoutubeClient, error) {
	for i, client := range clients {
		if !client.QuotaExceeded {
			return i, client, nil
		}
	}
	return 0, YoutubeClient{}, errors.New("No Client Found")
}

func GetYoutubeVideos(publishedAfter time.Time) []*youtube.SearchResult {
	logger := utils.Logger.Sugar()
	flag.Parse()

	clientIndex, YTClinet, err := getYoutybeClient()
	if err != nil {
		logger.Errorf(err.Error())
		return []*youtube.SearchResult{}
	}
	service, err := youtube.New(&YTClinet.HttpClient)
	if err != nil {
		logger.Errorf("Error creating new YouTube client: %v", err)
	}
	// publishedAfter := time.Now().Add(-time.Minute * 10)
	// Make the API call to YouTube.
	searchListCall := service.Search.List([]string{"id,snippet"}).
		Q(*query).
		MaxResults(*maxResults).
		Type("video").Order("date").PublishedAfter(publishedAfter.Format(time.RFC3339))
	response, err := searchListCall.Do()
	if err != nil {
		logger.Errorf("Error calling youtube client: %v", err)
	}
	if response.HTTPStatusCode == http.StatusForbidden {
		logger.Errorf("quota Exceeded")
		YTClinet.QuotaExceeded = true
		clients[clientIndex] = YTClinet
	}
	logger.Infof("Youtube Response : %v", response)
	return response.Items
}
