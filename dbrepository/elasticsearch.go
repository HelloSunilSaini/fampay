package dbrepository

import (
	"context"
	"encoding/json"
	"fampay/domain"
	"fampay/utils"
	"strings"

	"github.com/olivere/elastic/v7"
	"google.golang.org/api/youtube/v3"
)

type ElasticsearchRepo struct {
	Client    *elastic.Client
	IndexName string
}

var logger = utils.Logger.Sugar()

func (e *ElasticsearchRepo) AddOrUpdateVediosBulk(vedios []*youtube.SearchResult) error {
	bulkRequest := e.Client.Bulk()
	for _, vedio := range vedios {
		request := elastic.NewBulkIndexRequest().Index(e.IndexName).Doc(vedio).Id(vedio.Id.VideoId)
		bulkRequest.Add(request)
	}
	_, err := bulkRequest.Do(context.TODO())
	if err != nil {
		logger.Errorf("Error in persisting ES %s document: %v", e.IndexName, err)
	}
	return nil
}

func (e *ElasticsearchRepo) GetVedioDetailsBySearchTerm(searchTerm string, offset, size int) (*domain.VediosResponse, error) {
	query := e.getSearchQuery(searchTerm)
	sortBy := elastic.NewFieldSort("snippet.publishedAt").Desc()
	fetchSourceContext := elastic.NewFetchSourceContext(true)
	ss := elastic.NewSearchSource().Query(query).From(offset).FetchSourceContext(fetchSourceContext)
	ss = ss.SortBy(sortBy)
	ss = ss.Size(size)

	elasticRequest := e.Client.Search(e.IndexName).SearchSource(ss)
	searchResult, err := elasticRequest.Do(context.TODO())
	if err != nil {
		logger.Errorf("Error in fetching Products from ES %v", err)
		return nil, err
	}
	return e.transformESResultToVediosSearchResponse(searchResult)
}

func (e *ElasticsearchRepo) getSearchQuery(searchTerm string) elastic.Query {
	lowerSearchTerm := strings.ToLower(searchTerm)
	return elastic.NewBoolQuery().Should(
		elastic.NewMatchQuery("snippet.channelTitle", lowerSearchTerm),
		elastic.NewMatchQuery("snippet.description", lowerSearchTerm),
		elastic.NewMatchQuery("snippet.title", lowerSearchTerm),
	)
}

func (e *ElasticsearchRepo) transformESResultToVediosSearchResponse(searchResult *elastic.SearchResult) (*domain.VediosResponse, error) {
	resp := domain.VediosResponse{
		TotalCount: searchResult.Hits.TotalHits.Value,
		Vedios:     []youtube.SearchResult{},
	}
	logger.Debugf("Actaul ES Count:%v %v %v", searchResult.Hits.TotalHits.Value, searchResult.Hits.TotalHits.Relation, len(searchResult.Hits.Hits))
	for _, hit := range searchResult.Hits.Hits {
		esVedio := &youtube.SearchResult{}
		err := json.Unmarshal(hit.Source, &esVedio)
		if err != nil {
			logger.Errorf("Error in deserializing search result:{}, err:{}", hit.Source, err)
			continue
		}
		resp.Vedios = append(resp.Vedios, *esVedio)
	}
	return &resp, nil
}
