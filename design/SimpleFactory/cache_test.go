package SimpleFactory

import (
	"fmt"
	"testing"
)

func TestCacheFactory_Create(t *testing.T) {
	cacheFactory := &CacheFactory{}
	redisCache, err := cacheFactory.Create(redis)
	if err != nil {
		t.Error()
	}
	redisCache.Set("fox", "I love you")

	fmt.Println(redisCache.Get("fox"))
}
