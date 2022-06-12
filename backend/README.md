# ValushaJS API

- Golang & Gin

## POST /sign-in

Sign in with a Google account to leave comments.

## GET /comments

Get a flat list of top-level comments, sorted by date.

```
$ curl http://localhost:8000/comments
```

Note that API returns only the first page, if there's more that one. In this case `nextPageId` will be added in response.

```
$ curl http://localhost:8000/comments?pageId=...
```

^ to get the "next page".

## GET /comments/ID

Get a flat list of comments under the comment with `ID`.

## POST /comments

Add a new comment. You may provide `parentId` in the request body. The comment is considered as top-level unless `parentId` is provided.

```
$ curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ..." \
    -d '{"comment": "hello, world!", "parentId": "..."}' \
    http://localhost:8000/comments
```