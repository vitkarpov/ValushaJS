# ValushaJS API

- Golang & Gin

## Development

```sh
$ make postgres
$ make createdb
$ make migrate
```

## POST /sign-in

Sign in with a Google account to leave comments. A token is returned in response, which should be added to `Authorization: Bearer` header in further POST requests.

## GET /comments

Get a flat list of top-level comments, sorted by date.

```
$ curl http://localhost:8080/comments
```

API returns the first page by default. If there's more than one page, `nextPageId` is added in response so it's possible to request the next page.

> Note: there's only one sorting type so far, which is by date but we may support some more in the future, so instead of reffering to the next page by number it's by page id. That is for data integrity, e.g. if a new comment is added on top the next page doesn't have a duplicate from the previous page.

```
$ curl http://localhost:8080/comments?pageId=...
```

^ to get the "next page".

## GET /comments/ID

Get a flat list of comments under the comment with `ID`.

## POST /comments

Add a new comment. You may provide `parentId` in the request body. The comment is considered as top-level unless `parentId` is provided.

```
$ curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ..." \
    -d '{"comment": "hello, world!", "parentId": "..."}' \
    http://localhost:8080/comments
```