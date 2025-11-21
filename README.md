# httpecho

This is a simple HTTP webserver that can be used for various testing.

## Usage

Start the program and an HTTP server will be listening on port 8080.

This is easy to start with docker:

```
docker run -it -p 8080:8080 ianneub/httpecho
```

## Endpoints

`/` - Will echo the HTTP request headers back to the client

`/long` - Will sleep for 300 seconds by default and return OK to the client. You can set the `s` query string to the number of seconds you want to wait before returning.
