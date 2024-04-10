# Shorty
## A very simple URL shortener in-memory with [Golang](https://golang.org)

```bash
go build -o shorty .

```

### Usage
```bash
curl -X POST "http://localhost:8080/shorten?url=http://kernel.org"
{"longURL":"http://kernel.org","shortURL":"tIx7kV"}%
curl http://localhost:8080/tIx7kV
<a href="http://kernel.org">Moved Permanently</a>.
```

