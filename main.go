package main

import (
	"fampay/apihandlers"
	"fampay/client"
	"fampay/dbrepository"
	"fampay/utils"
	"flag"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/olivere/elastic/v7"
)

func init() {
	http.DefaultClient.Timeout = time.Minute * 10
}

const ElasticSearchIndex = "youtubevedios"

var (
	elasticSearchURL      = flag.String("elasticsearchurl", os.Getenv("ELASTICSEARCHURL"), "elastic search url")
	elasticSearchUserName = flag.String("elasticsearchusername", os.Getenv("ES_USERNAME"), "elastic search username")
	elasticSearchPassword = flag.String("elasticsearchpassword", os.Getenv("ES_PASSWORD"), "elastic search password")
	developerkeys         = flag.String("developerkeys", os.Getenv("DEVLOPER_KEYS"), "elastic search password")
	logger                = utils.Logger.Sugar()
)

func main() {
	// initial delay to wait for elastic search container setup completion
	time.Sleep(time.Second * 10)
	if *elasticSearchURL == "" {
		panic("elasticsearch url not set")
	}
	if *developerkeys == "" {
		panic("DEVLOPER_KEYS not provided")
	}
	elasticSearchClient, err := elastic.NewClient(elastic.SetBasicAuth(*elasticSearchUserName, *elasticSearchPassword), elastic.SetURL(*elasticSearchURL), elastic.SetSniff(false))
	if err != nil {
		logger.Errorf(err.Error())
		panic("Not able to create client")
	}
	elasticSearchRepo := &dbrepository.ElasticsearchRepo{
		Client:    elasticSearchClient,
		IndexName: ElasticSearchIndex,
	}
	go BackGoundWorker(elasticSearchRepo)
	pingHandler := &apihandlers.PingHandler{}
	youtubeHandler := &apihandlers.YoutubeHandler{}
	vediosHandler := &apihandlers.VediosHandler{DBRepo: elasticSearchRepo}

	logger.Info("Setting up resources.")
	h := mux.NewRouter()

	h.Handle("/fampay/ping/", pingHandler)
	h.Handle("/fampay/youtube/", youtubeHandler)
	h.Handle("/fampay/vedios/", vediosHandler)
	logger.Info("Resource Setup Done.")

	addr := ":7776"
	s := &http.Server{
		Addr:         addr,
		Handler:      h,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	s.ListenAndServe()
}

func BackGoundWorker(repo dbrepository.IRepo) {
	devKeys := strings.Split(*developerkeys, ",")
	client.InitYTClients(devKeys)
	for {
		publishedAfter := time.Now().Add(-time.Minute * 10)
		vedios := client.GetYoutubeVideos(publishedAfter)
		repo.AddOrUpdateVediosBulk(vedios)
		time.Sleep(time.Second * 10)
	}
}
