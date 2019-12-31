### REST API using Gin framework

A simple REST Api web service to test event webhooks with Hasura graphql-engine.

### Usage

* Clone the repo
* Run `make dev` and `make console`.
* Open up hasura console on  `localhost:8080`.
* Create a table `articles` with `title(text)`, `body(text)`, `count(int)` as fields.
* Navigate to `Event triggers` on hasura console and set up a event trigger. Follow the [docs](https://docs.hasura.io/1.0/graphql/manual/event-triggers/create-trigger.html).
* Insert some data in the table but add `count` as 0.
* After you insert you should see the record was created succesfully and the `count` was updated to the word count in the `body`.

___
