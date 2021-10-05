package main

import (
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/alecthomas/kingpin"
)

var (
	requests    = kingpin.Flag("requests", "requests").Short('n').Default("100").Int()
	concurrency = kingpin.Flag("concurrency", "concurrency").Short('c').Default("2").Int()
	endpointURL = kingpin.Flag("endpoint-url", "endpoint-url").Short('e').String()
	token       = kingpin.Flag("token", "token").Default("xxxx").Short('t').String()

	apiKey string
)

func init() {
	kingpin.Parse()

	rand.Seed(time.Now().UnixNano())
	gofakeit.Seed(time.Now().UnixNano())

	// Encode token to base64
	apiKey = base64.StdEncoding.EncodeToString([]byte(*token))
}

func main() {
	StartWorkers(*concurrency)

	for i := 0; i < *requests; i++ {
		AddJob(i)
	}

	StopWorkers()
}
