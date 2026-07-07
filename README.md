# httpfromtcp

Building an HTTP/1.1 server from raw TCP in Go — following ThePrimeagen's
["From TCP to HTTP" full course](https://www.youtube.com/watch?v=FknTw9bJsXM).

> **Learning log, not a portfolio piece.** The code follows the course; the
> notes and anything under *Extensions* are mine.

## Progress

Course chapters (will adjust as I go):

- [x] TCP listener
- [x] UDP sender detour
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

### Roles and direction: who listens, who connects

The thing I kept tripping on. The two ends of a connection are *not*
symmetric:

- **Server / receiver** — *passive*. Binds a port and waits (`net.Listen` +
  `Accept`, or `nc -l`). Does nothing until someone shows up.
- **Client / sender** — *active*. Reaches out to a `host:port` (`net.Dial` /
  `DialUDP`, or plain `nc host port`).

`nc`'s role flips with its flags, which is exactly what confused me:
`nc localhost 42069` is a **client** (connects out); `nc -u -l 42069` is a
**listener** (`-l`) over UDP (`-u`). Same tool, opposite end of the wire.
Data always flows sender → receiver. HTTP is just this with names: the client
sends a request, the server sends back a response.

### A connection is just an `io.Reader` (a byte stream)

The spine of the whole course. `os.File` and `net.Conn` both satisfy
`io.ReadCloser`, so `getLinesChannel` didn't change *one line* when I swapped
a file for a TCP connection — same code, different source.

The consequence that matters: a stream arrives in **arbitrary chunks**, not
tidy messages. Reading 8 bytes at a time forces this — one line can span
several reads, or several lines can arrive in one read. My job is to buffer
the partial pieces and reassemble the meaning myself. Parsing HTTP later is
the same move, just with richer boundaries (`\r\n`, request-line, headers,
body) instead of a plain `\n`.

### TCP vs UDP, felt firsthand

- TCP needs a handshake *before* any data → connecting to a dead port fails
  immediately with `connection refused`.
- UDP just yeets the packet → `conn.Write` to a port with no listener
  succeeds with no error (the OS *may* surface `connection refused` on a
  later write via ICMP, but there's no connection to establish).
- Why I split error handling in the sender: **read** errors are terminal
  (EOF = input is over → stop), **write** errors are transient (one packet
  failed → log and keep going).
