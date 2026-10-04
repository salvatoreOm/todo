# Todo CLI (Go)

A command-line todo app written in Go, built as a learning project with production-style practices.

## Reference

- https://github.com/patni1992/CLI-Todo-App-In-Go: the reference implementation (flat layout: `main.go`, `command.go`, `todo.go`, `storage.go`; companion tutorial at codingwithpatrik.dev). We use it for feature scope only. Our structure is more layered (see below).

## Conventions

- Go 1.27, standard library first; add a dependency only with a clear reason.
- Layers depend inward only: `cli -> service -> store`; `todo` (domain) imports nothing from the project.
- Interfaces are defined by the consumer (service defines `Store`), not the implementer.
- Wrap errors with `%w`; use sentinel errors (`todo.ErrNotFound`) and check with `errors.Is`.
- Pass `context.Context` as the first argument to service/store methods.
- Inject the clock and the store so the service can be unit-tested.
- Tests: table-driven, run with `go test ./... -race`.
- Format and vet: `gofmt`, `go vet ./...` before finishing a change.
