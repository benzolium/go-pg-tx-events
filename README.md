# go-pg-tx-events
Golang transactional events with postgresql notify

Based on:
https://blog.insiderattack.net/atomic-microservices-transactions-with-mongodb-transactional-outbox-1c96e0522e7c

Pros:
- No more callbacks with anonymous functions

Cons:
- If notify is made before consumes starts, event will be lost
