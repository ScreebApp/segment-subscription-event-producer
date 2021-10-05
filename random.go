package main

import (
	"math/rand"
	"net/url"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	uuid "github.com/satori/go.uuid"
)

type PropertyType string

const (
	PropertyTypeString  PropertyType = "string"
	PropertyTypeNumber  PropertyType = "number"
	PropertyTypeDate    PropertyType = "date"
	PropertyTypeBoolean PropertyType = "boolean"
	PropertyTypeObject  PropertyType = "object"
	PropertyTypeArray   PropertyType = "array"
	PropertyTypeNil     PropertyType = "nil"
)

var propertyTypes = []PropertyType{
	PropertyTypeString,
	PropertyTypeString,
	PropertyTypeString,
	PropertyTypeNumber,
	PropertyTypeNumber,
	PropertyTypeDate,
	PropertyTypeBoolean,
	PropertyTypeObject,
	PropertyTypeArray,
	PropertyTypeNil,
}

const (
	IDENTITY_COUNT    = 42000
	PROPERTY_COUNT    = 50
	EVENT_NAME_COUNT  = 200
	SCREEN_NAME_COUNT = 100
)

var identities []struct {
	UserID      string
	AnonymousID []string
}
var properties []struct {
	Name string
	Type PropertyType
}
var eventNames []string
var screenNames []string

func init() {
	// init identities userId and anonymousIds
	for i := 0; i < IDENTITY_COUNT; i++ {
		identity := struct {
			UserID      string
			AnonymousID []string
		}{
			UserID: "u-" + uuid.NewV4().String(),
		}

		for i := 0; i < 1+rand.Intn(4); i++ { // 1 to 5 AnonymousID per identity
			identity.AnonymousID = append(identity.AnonymousID, "a-"+uuid.NewV4().String())
		}

		identities = append(identities, identity)
	}

	// init properties name and type
	for i := 0; i < PROPERTY_COUNT; i++ {
		name := strings.ReplaceAll(pureString(gofakeit.HipsterSentence(1+rand.Intn(5))), " ", "_")
		properties = append(properties, struct {
			Name string
			Type PropertyType
		}{
			Name: name,
			Type: propertyTypes[rand.Intn(len(propertyTypes))],
		})
	}

	// init events name
	for i := 0; i < EVENT_NAME_COUNT; i++ {
		eventNames = append(eventNames, pureString(gofakeit.HipsterSentence(1+rand.Intn(5))))
	}

	// init screens name
	for i := 0; i < SCREEN_NAME_COUNT; i++ {
		screenNames = append(screenNames, pureString(gofakeit.HipsterSentence(1+rand.Intn(5))))
	}
}

func pureString(str string) string {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, ",", "")
	str = strings.ReplaceAll(str, "'", "")
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ReplaceAll(str, "_", "")
	str = strings.ReplaceAll(str, "-", "")
	return str
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandPageTitle() string {
	// part := []string{}

	// for i := 0; i < 2+rand.Intn(5); i++ {
	// 	part = append(part, RandStringRunes(2+rand.Intn(6)))
	// }

	// return strings.Join(part, " ") + " | " + RandStringRunes(2+rand.Intn(6))

	return gofakeit.Dessert() + " | " + gofakeit.Company()
}

func RandURL() (string, string, string, error) {
	u, err := url.Parse(gofakeit.URL())
	if err != nil {
		return "", "", "", err
	}

	search := u.RawQuery
	if len(search) > 0 {
		search = "?" + search
	}

	return u.String(), u.Path, search, nil
}

func RandIdentity() (string, []string) {
	identity := identities[rand.Intn(IDENTITY_COUNT)]
	return identity.UserID, identity.AnonymousID
}

func RandEventName() string {
	return eventNames[rand.Intn(EVENT_NAME_COUNT)]
}

func RandScreenName() string {
	return strings.ReplaceAll(eventNames[rand.Intn(SCREEN_NAME_COUNT)], " ", "")
}

func RandProperty() (string, interface{}) {
	prop := properties[rand.Intn(PROPERTY_COUNT)]

	switch prop.Type {
	case PropertyTypeString:
		return prop.Name, gofakeit.HipsterSentence(1 + rand.Intn(9))
	case PropertyTypeNumber:
		return prop.Name, gofakeit.Int16()
	case PropertyTypeDate:
		return prop.Name, gofakeit.Date()
	case PropertyTypeBoolean:
		return prop.Name, gofakeit.Bool()
	case PropertyTypeObject:
		return prop.Name, gofakeit.Address()
	case PropertyTypeArray:
		return prop.Name, []string{gofakeit.Fruit(), gofakeit.FarmAnimal()}
	case PropertyTypeNil:
		return prop.Name, nil
	}

	return prop.Name, nil
}

func RandProperties() map[string]interface{} {
	output := map[string]interface{}{}

	for i := 0; i < 3+rand.Intn(25); i++ {
		name, value := RandProperty()
		output[name] = value
	}

	return output
}
