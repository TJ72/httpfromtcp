# httpfromtcp

Building an HTTP/1.1 server from raw TCP in Go — following ThePrimeagen's
["From TCP to HTTP" full course](https://www.youtube.com/watch?v=FknTw9bJsXM).

> **Learning log, not a portfolio piece.** The code follows the course; the
> notes and anything under *Extensions* are mine.

## Progress

Course chapters (will adjust as I go):

- [ ] TCP listener
- [ ] UDP sender detour
- [ ] Request line parsing
- [ ] Header parsing
- [ ] Body parsing
- [ ] Response writing
- [ ] HTTP server
- [ ] Chunked encoding
- [ ] Trailers
- [ ] Binary data / streaming

## Extensions (after the course — the part that's actually mine)

- [ ] Connection keep-alive
- [ ] Benchmark against `net/http` and write up the numbers

## Notes

Running notes as I go: what surprised me, where my mental model of HTTP was
wrong, what the stdlib does differently.

*(nothing yet)*
