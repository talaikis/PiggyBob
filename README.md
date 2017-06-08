# [PiggyBob](https://piggybob.com)

Simple Personal Finance Tracker.

In dev

DON'T LOOK, IT DOESN'T WORK AND IS IN DEVELOPMENT.

## Technologies

* [Golang](https://github.com/golang/go) 1.8
* [Goth](https://github.com/markbates/goth)
* [pq](https://github.com/lib/pq)
* ...

## Run

First, configure .env, create (Postgres) database, then:

```text
git clone https://github.com/xenu256/PiggyBob
cd PiggyBob
go get
go run db_utils.go //this creates db tables
go run main.go
```

## History

* 2007-2009 it was a simple personal finance and time management program for Windows.
* 2017. This idea turned web. First of thought to implement in Python, but then I have decided to learn Go and this project was right on the todo list.
