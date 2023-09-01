package search

import (
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/gomodule/redigo/redis"
	"thichlab-backend-docs/constant"
	"thichlab-backend-docs/dto"
	"thichlab-backend-docs/infrastructure/logger"
	"thichlab-backend-docs/infrastructure/util"
)

type Client struct {
	Client *redisearch.Client
}

func (c *Client) InitializeConnection(addr, name, password string) {
	var err error
	pool := &redis.Pool{Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", addr, redis.DialPassword(password))
	}}
	c.Client = redisearch.NewClientFromPool(pool, name)

	// Create a schema
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("body")).
		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
		AddField(redisearch.NewNumericField("date"))

	// Check if index does not exist in index list then create new index
	listIndex, _ := c.Client.List()
	if !util.CheckStringInSlice(listIndex, name) {
		if err = c.Client.CreateIndex(sc); err != nil {
			logger.Error("[REDIS-SEARCH] Create the index - ERROR: %s", err)
		}
	}

}

func (c *Client) CreateSearchDocument(id, title, content, slug string) error {
	idDoc := fmt.Sprintf("%s:%s", constant.CacheSearchPost, id)
	doc := redisearch.NewDocument(idDoc, 1.0)
	doc.Set("title", title).
		Set("title_format", util.ToSearchFormat(title)).
		Set("content", util.ToSearchFormat(content)).
		Set("slug", slug)

	// Index the document. The API accepts multiple documents at a time
	if err := c.Client.IndexOptions(redisearch.DefaultIndexingOptions, doc); err != nil {
		logger.Error("[REDIS-SEARCH] Index Document - ERROR: %s", err)
	}

	return nil
}

func (c *Client) SearchDocs(query string, pageIndex, pageSize int) (*dto.PostsByDocSearchRes, error) {
	docs, total, err := c.Client.Search(redisearch.NewQuery(query).
		Limit(pageIndex*pageSize-pageSize, pageSize).
		SetReturnFields("title", "slug"))

	if err != nil {
		logger.Error("[REDIS-SEARCH] SearchDocs - ERROR: %s", err)
		return nil, err
	}

	var result []dto.PostsByDocSearch
	for i, _ := range docs {
		result = append(result, dto.PostsByDocSearch{
			Title: (docs[i].Properties["title"]).(string),
			Slug:  (docs[i].Properties["slug"]).(string),
		})
	}

	resultRes := dto.PostsByDocSearchRes{
		Total: total,
		Data:  result,
	}

	return &resultRes, nil
}

func (c *Client) DelDoc(id string) error {
	return c.Client.DeleteDocument(id)

}
