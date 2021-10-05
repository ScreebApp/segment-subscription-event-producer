package main

import (
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	uuid "github.com/satori/go.uuid"
)

type SegmentEventType string

const (
	SegmentEventTypeTrack    SegmentEventType = "track"
	SegmentEventTypeIdentify SegmentEventType = "identify"
	SegmentEventTypeGroup    SegmentEventType = "group"
	SegmentEventTypeScreen   SegmentEventType = "screen"
	SegmentEventTypePage     SegmentEventType = "page"
	SegmentEventTypeAlias    SegmentEventType = "alias"
)

type SegmentEventBody struct {
	Version     int                    `json:"version"`
	Context     map[string]interface{} `json:"context"` // @TODO
	MessageID   string                 `json:"messageId"`
	SentAt      time.Time              `json:"sentAt"`
	ReceivedAt  time.Time              `json:"receivedAt"`
	Timestamp   time.Time              `json:"timestamp"`
	Type        SegmentEventType       `json:"type"`
	ProjectID   string                 `json:"projectId"`
	Replay      bool                   `json:"replay"`
	UserID      *string                `json:"userId"`
	AnonymousID *string                `json:"anonymousId"`
}

/**
 * Identities
 */
type SegmentEventBodyIdentify struct {
	SegmentEventBody

	Traits map[string]interface{} `json:"traits"`
}

type SegmentEventBodyGroup struct {
	SegmentEventBody

	GroupID string                 `json:"groupId"`
	Traits  map[string]interface{} `json:"traits"`
}

type SegmentEventBodyAlias struct {
	SegmentEventBody

	PreviousID string `json:"previousId"`
}

/**
 * Events
 */
type SegmentEventBodyTrack struct {
	SegmentEventBody

	Event      string                 `json:"event"`
	Properties map[string]interface{} `json:"properties"`
}

type SegmentEventBodyScreen struct {
	SegmentEventBody

	Name       string                 `json:"name"`
	Properties map[string]interface{} `json:"properties"`
}

type SegmentEventBodyPage struct {
	SegmentEventBody

	// Name       string                 `json:"name"`	// for future use
	Properties map[string]interface{} `json:"properties"`
}

/**
 * Request
 */
func buildRequestCommon(eventType SegmentEventType) SegmentEventBody {
	userID, anonymousIDs := RandIdentity()
	_user := &userID
	_anonymousID := &anonymousIDs[len(anonymousIDs)-1]

	// half of request will bring an userId
	if rand.Intn(2) == 0 {
		_user = nil
	}

	// when userID is not set, 80% of requests will bring an anonymousId
	if _user != nil && rand.Intn(5) == 0 {
		_anonymousID = nil
	}

	timestamp := time.Now()
	receivedAt := timestamp.Add(-1 * time.Duration(rand.Intn(200)) * time.Millisecond)
	sentAt := receivedAt.Add(-1 * time.Duration(rand.Intn(200)) * time.Millisecond)

	return SegmentEventBody{
		Version:     2,
		Context:     map[string]interface{}{},
		MessageID:   "m-" + uuid.NewV4().String(),
		SentAt:      sentAt,
		ReceivedAt:  receivedAt,
		Timestamp:   timestamp,
		Type:        eventType,
		ProjectID:   "benchmark-segment-subscription",
		Replay:      false,
		UserID:      _user,
		AnonymousID: _anonymousID,
	}
}

func BuildRequestIdentify() error {
	body := SegmentEventBodyIdentify{
		SegmentEventBody: buildRequestCommon(SegmentEventTypeIdentify),

		Traits: RandProperties(),
	}

	return sendWebhookRequest(body)
}

func BuildRequestGroup() error {
	body := SegmentEventBodyGroup{
		SegmentEventBody: buildRequestCommon(SegmentEventTypeGroup),

		GroupID: "g-" + uuid.NewV4().String(),
		Traits:  RandProperties(),
	}

	return sendWebhookRequest(body)
}

func BuildRequestAlias() error {
	body := SegmentEventBodyAlias{
		SegmentEventBody: buildRequestCommon(SegmentEventTypeAlias),

		PreviousID: "p-" + uuid.NewV4().String(),
	}

	return sendWebhookRequest(body)
}

func BuildRequestTrack() error {
	body := SegmentEventBodyTrack{
		SegmentEventBody: buildRequestCommon(SegmentEventTypeTrack),

		Event:      RandEventName(),
		Properties: RandProperties(),
	}

	return sendWebhookRequest(body)
}

func BuildRequestPage() error {
	fullURL, path, search, err := RandURL()
	if err != nil {
		return err
	}

	body := SegmentEventBodyPage{
		SegmentEventBody: buildRequestCommon(SegmentEventTypePage),

		Properties: map[string]interface{}{
			"path":     path,           // eg: '/toto/tata/titi'
			"referrer": gofakeit.URL(), // full url
			"search":   search,         // eg: '?from=2021-10-05&to=2021-10-05'
			"title":    RandPageTitle(),
			"url":      fullURL,
			"keywords": strings.Split(pureString(gofakeit.HipsterSentence(1+rand.Intn(5))), " "),
			// "keywords": []string{
			// 	gofakeit.Animal(),
			// 	gofakeit.Fruit(),
			// 	gofakeit.FarmAnimal(),
			// },
		},
	}

	return sendWebhookRequest(body)
}

func BuildRequestScreen() error {
	body := SegmentEventBodyScreen{
		SegmentEventBody: buildRequestCommon(SegmentEventTypeScreen),

		Name:       RandScreenName(),
		Properties: RandProperties(),
	}

	return sendWebhookRequest(body)
}
