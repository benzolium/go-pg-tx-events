# go-pg-tx-events
Golang transactional events with postgresql notify

Based on:
https://blog.insiderattack.net/atomic-microservices-transactions-with-mongodb-transactional-outbox-1c96e0522e7c

# Usage

- create table to store events
- create trigger to notify of table's updates
- write function to emit to ampq (or anything else) on postgresql's notify
- re-write your app's publisher to write to database

Sql that will emit message to amqp:
```sql
begin;

-- some app's logic
insert into events (object_id, object_routing_key, data)
values (1, 'user.created', '{"id": 1, "name": "John"}'::json);

commit;
```

Sql that will not emit message to amqp:
```sql
begin;

-- some app's logic
insert into events (object_id, object_routing_key, data)
values (1, 'user.created', '{"id": 1, "name": "John"}'::json);

rollback;
```

# Pros & Cons

Pros:
- No more callbacks with anonymous functions

Cons:
- If notify is made before consumes starts, event will be lost
