### Web service made with Gin

A Go web service to test webhooks on postgres insert event on Hasura Graphql Engine.

### Usage

* Clone the repo
* Run `make dev` and `make console`.
* Open up hasura console on  `localhost:8080`.
* Create a table `articles` with `title(text)`, `body(text)`, `count(int)` as fields.
* Navigate to `Event triggers` on hasura console and set up a event trigger. Follow the [docs](https://docs.hasura.io/1.0/graphql/manual/event-triggers/create-trigger.html).
* Insert some data in the table but add `count` as 0.
* After you insert you should see the record was created succesfully and the `count` was updated to the word count in the `body`.

___

### What is Hasura GraphQL Engine?
Hasura GraphQL Engine lets you make powerful queries with built-in filtering, pagination, pattern search, bulk insert, update, delete mutations & subscriptions. You can also Trigger webhooks or serverless functions on Postgres insert/update/delete events. Comes with fine-grained dynamic access control that integrates with your auth system. This one-click setup also includes an empty Postgres database and automatic HTTPS from Let’s Encrypt using Caddy webserver.