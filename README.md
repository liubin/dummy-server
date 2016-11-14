Dummy(mock) server used for API client test by response with json files.


# How it works

If you have an API `/api/v1/users`, save your response file `/api/v1/users.json` under `./responses` directory.

And when this API is called, the file's content will be sent to caller.

You can use this program as a dummy(mock) server for client test.


# How to use:

## Start server

```
./dummy-server -port 8082
```

Default `port` is `8081` .

## Save API response JSON file to responses

For example:

```
$ cat responses/api/v1/users.json 
{"user": "12345"}
```

## Call from client

```
$ curl -s localhost:8082/api/v1/users | jq .
{
  "user": "12345"
}

```

