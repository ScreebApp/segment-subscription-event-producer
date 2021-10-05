
# Segment event producer for subscription developers

This project has been built for helping companies to create Segment connectors using subscriptions API (webhook). It allows to send a high volume of identity/group/alias/track/page/screen calls.

Segment subscriptions are documented here: [https://segment.com/docs/partners/subscriptions/build-webhook/](https://segment.com/docs/partners/subscriptions/build-webhook/).

For load testing, just increase the `concurrency` and `requests` arguments.

<p align="center">
    <img src="doc/fish.gif" width="384" height="480">
</p>

## Overview of event definition

### Payload

- API schema v2
- 42k different `identities` (1 userId + 5 anonymousIds)
- 50 property names
- 200 event names
- 100 screen names
- 100% random URLs
- 50% calls don't have userId
- 20% calls don't have anonymousId
- replay always false
- `sentAt` < `receivedAt` < `timestamp`

During execution, properties sent into `identify`, `group`, `track`, `page` and `screen` calls will always have the same data type (number, string, datetime, boolean, object, array).

### Ratio of calls per type

- 20% `identify`
- 5% `group`
- 5% `alias`
- 15% `track`
- 30% `page`
- 25% `screen`

Events will be triggered in a random order, and will be emitted from random users.

### Auth

This tool only support basic token authentication.

## Run

```bash
go get github.com/ScreebApp/benchmark-segment-subscription

go run github.com/ScreebApp/benchmark-segment-subscription \
	--requests 10000 \
	--concurrency 100 \
	--endpoint-url http://localhost:3000/webhook \
	--token xxxxxx
```

## Contrib

```bash
go run *.go -n 100 -c 2 -e http://localhost:3000/webhook -t xxx
```

## Todo

- Add `context` payload (see "common fields" section in doc)
- Collect and print response status code, at the end of execution
- Collect and display response time, at the end of execution
- Disable some events type from CLI
- Display progress bar
