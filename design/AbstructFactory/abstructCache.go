package AbstructFactory

//定义一个Cache接口，作为父类
type Cache interface {
	Get(k string) string
	Set(k, v string)
}

//实现具体的Cache：RedisCache
type RedisCache struct {
	data map[string]string
}

func (redis *RedisCache) Set(key, value string) {
	redis.data[key] = value
}

func (redis *RedisCache) Get(key string) string {
	return "redis:" + redis.data[key]
}

//实现具体的Cache：MemberCache
type MemCache struct {
	data map[string]string
}

func (mem *MemCache) Set(key, value string) {
	mem.data[key] = value
}

func (mem *MemCache) Get(key string) string {
	return "mem:" + mem.data[key]
}

//定义一个抽象Cache工厂
type CacheFactory interface {
	Create() Cache
}

//创建一个redis工厂
type RedisCacheFactory struct {
}

func (rf *RedisCacheFactory) Create() Cache {
	return rf.Create()
}

//创建一个mem工厂
type MemCacheFactory struct {
}

func (mem *MemCacheFactory) Create() Cache {
	return mem.Create()
}
