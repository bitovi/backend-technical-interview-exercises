# Bitovi Backend Debugging Exercise

## Bug Report

The Marketing department reported that a high number of marketing emails are bouncing back due to the email address not being correct.

Your task is to identify the cause of the bug.

## Reproducing the issue

Run the ETL script the marketing department uses to get order and customer summaries:

```
go run etl.go
```

## Exercise tools

Since this exercise is done outside of your normal development environment, here are a few tools you can use to investigate the issue:

To make a request to one of the APIs:
```
curl http://localhost:3000/orders
```

To view a list of database tables:
```
psql postgresql://dbuser:dbpassword@localhost:5432/postgres -c "\dt"
```

To run a SQL statement against one of the tables:
```
psql postgresql://dbuser:dbpassword@localhost:5432/postgres -c "select * from product"
```

## Interviewer setup instructions

Note: do this before the interview starts, as it can take a minute.
```
docker compose up
```