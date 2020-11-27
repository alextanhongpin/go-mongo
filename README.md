# go-mongo

Testing out MongoDB with golang's driver.


## Quickstart

```bash
# Spins up dockerized mongo.
$ make up

# Start the application.
$ make start

# Spins down docker containers.
$ make down
```

## Thoughts

- How to handle migrations?
- How to handle creation of indices?
- How to handle the transactions between different stores (a.k.a repositories)?
- Are there Change-data-capture (CDC) functionality ?
- Is this still a better choice over Postgres?
