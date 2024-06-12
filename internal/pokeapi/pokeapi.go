package pokeapi

import (
	"net/http"
	"time"

	"github.com/ggrangel/pokedex/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheEvictionFrequency time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheEvictionFrequency),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
