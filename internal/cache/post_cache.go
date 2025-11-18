package cache

import (
	"encoding/json"
	"fmt"
	"post-management/pkg/dto"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type PostCache struct {
	cache *memcache.Client
}

func NewPostCache(cache *memcache.Client) *PostCache {
	return &PostCache{
		cache: cache,
	}
}
func (c *PostCache) SetById(postId int64, value *dto.PostResponse) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = c.cache.Set(&memcache.Item{
		Key:        fmt.Sprintf("post:%d", postId),
		Value:      data,
		Expiration: int32(time.Now().Add(15 * time.Minute).Unix()),
	})
	if err != nil {
		return err
	}
	return nil
}
func (c *PostCache) GetById(postId int64) (*dto.PostResponse, error) {
	item, err := c.cache.Get(fmt.Sprintf("post:%d", postId))
	if err != nil {
		return nil, err
	}
	resp := new(dto.PostResponse)
	err = json.Unmarshal(item.Value, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (c *PostCache) DeleteById(postId int64) error {
	err := c.cache.Delete(fmt.Sprintf("post:%d", postId))
	if err != nil {
		return err
	}
	return nil
}
