package utils

import (
	"net/http"
	"sync"
)

func GetClient() *http.Client {
	var once sync.Once
	var c *http.Client

	once.Do(func() {
		c = &http.Client{}
	})

	return c
}
