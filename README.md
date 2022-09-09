# Overview
In this project, I'm using Debezium and Apache Kafka to observe the changes in a PostgreSQL database.
After a record is inserted or updated, a message is published on a topic, then handled by one of the Go application's consumers.

For example, imagine we have a table `public.billionaires`:

| id | name         | net_worth |
|----|--------------|----------:|
| 1  | Elon Musk    | 262       |
| 2  | Jeff Bezos   | 157       |
| 3  | Gautam Adani | 143       |

If we insert a record with the ID of `4`, a message would be published on the topic `postgres.public.billionaires`:
```json
{
  "before": null,
  "after": {
    "id": 4,
    "name": "Bill Gates",
    "net_worth": 116
  }
}
```

However, instead of inserting a new item, we update an existing record:
```json
{
  "before": {
    "id": 4,
    "name": "Jeff Bezos",
    "net_worth": 157
  },
  "after": {
    "id": 4,
    "name": "Jeff Bezos",
    "net_worth": 180
  }
}
```

# Sample Database Structure
```sql
CREATE DATABASE debezium_test;

-- Connect to the database "debezium_test"

CREATE TABLE billionaires (
  'id'        SERIAL PRIMARY KEY,
  'name'      VARCHAR(255),
  'net_worth' NUMERIC
);
```

# How to set up
You'll need Go, Docker and Docker Compose set up on your computer to run this project.
```sh
$ docker compose up
$ go run cmd/main.go
```

---

Made with ❤️ by Dominick Brasileiro 💎

Feel free [to reach out](https://www.linkedin.com/in/dominickbrasileiro/)!
