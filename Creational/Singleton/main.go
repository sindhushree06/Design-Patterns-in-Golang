package main

import (
	"sync"
)

type config struct {
	url string
}

var instance *config
var once sync.Once

// once.do will get executed only once
func GetConfig() *config {
	once.Do(func() {
		instance = &config{
			url: "localhost:8080",
		}
	})
	return instance
}

func main() {
}
