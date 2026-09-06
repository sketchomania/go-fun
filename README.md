# Go Fun

A collection of small Go programs created while exploring web servers,
databases, file conversion, and memory layout. Each directory is an independent
example, so there is no single command that runs the entire repository.

These projects are intended for learning and experimentation. Review and adapt
their configuration before exposing them to a network or using them with
production data.

## Projects

- [`bee-crud`](./bee-crud/) — Beego v2 and MySQL ORM experiments.
- [`beego`](./beego/) — a minimal Beego v2 server.
- [`converter-md`](./converter-md/) — converts `test.md` to `test.html`.
- [`converter-pfd`](./converter-pfd/) — converts `test.txt` to
  `generated.pdf`. The directory name is retained for compatibility.
- [`mysql-go`](./mysql-go/) — prepared-statement CRUD examples for MySQL.
- [`simple-http-server`](./simple-http-server/) — serves files over HTTP on
  port `9000`.
- [`local_file_server.go`](./local_file_server.go) — serves selected local
  directories on port `9999`; update the example directory paths first.
- [`struct-field-optimised.go`](./struct-field-optimised.go) — demonstrates how
  field ordering affects struct size.

## Requirements

- Go 1.26 or newer for the Beego examples
- MySQL for `bee-crud` and `mysql-go`

The converter and HTTP server examples do not require a database.

## Running an example

Clone the repository, enter an example directory, then use `go run`:

```bash
git clone https://github.com/<your-username>/go-fun.git
cd go-fun/converter-md
go run .
```

Other useful commands:

```bash
# Text to PDF
cd converter-pfd
go run .

# Simple HTTP server
cd simple-http-server
go run server.go

# Beego server
cd beego
go run .
```

For the database examples, create the databases and tables expected by the
source code and replace the local demonstration credentials before running
them.

## Testing

Run tests separately in each directory that contains a `go.mod` file:

```bash
cd bee-crud
go test ./...
```

The remaining examples can be checked with `go build` or `go run` from their
respective directories.

## Dependency security

Dependencies are tracked independently by each Go module. Dependabot alerts are
enabled for the public repository, and the affected Beego, Markdown, Go
cryptography, networking, text, Prometheus, and protobuf dependencies have been
upgraded.

For an additional local check, install and run Go's vulnerability scanner in a
module directory:

```bash
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
```
