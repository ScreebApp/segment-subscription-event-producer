
# Benchmark for Segment subscriptions developers

This project has been built for helping companies to create Segment connectors over subscriptions (webhook).

Segment subscriptions are documented here: https://segment.com/docs/partners/subscriptions/build-webhook/

Run:

```bash
go run *.go -n 100 -c 2 -e http://localhost:3000/webhook -t xxx
```
