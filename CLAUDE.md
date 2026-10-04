# Todo CLI (Go)

A command-line todo app written in Go, built as a learning project with production-style practices.

## Reference

- https://github.com/patni1992/CLI-Todo-App-In-Go: the reference implementation (flat layout, companion tutorial at codingwithpatrik.dev). We use it for feature scope only. Our structure is more layered.

## Stack

Go 1.27, Cobra (CLI), GORM + PostgreSQL (run with Podman via `compose.yaml`), godotenv (loads `.env`).

## Layout

```
cmd/todo/main.go          wiring only: db -> repository -> service -> handler
internal/model/           Todo struct (GORM schema) and sentinel errors
internal/database/        Connect + AutoMigrate
internal/repository/      GORM calls; the only layer that knows GORM
internal/service/         business rules (title validation, toggle logic)
internal/handler/         Cobra commands: add, list, toggle, update, delete
```

## Conventions

- Layers depend downward only: `handler -> service -> repository -> database`.
- Each layer is used through an interface (`TodoService`, `TodoRepository`) so it can be faked in tests.
- Wrap errors with `%w`; use sentinel errors (`model.ErrNotFound`, `model.ErrEmptyTitle`) and check with `errors.Is`.
- Pass `context.Context` as the first argument to service and repository methods.
- Commands print with `cmd.OutOrStdout()`, not `fmt.Println`, so output can be tested.
- Tests: table-driven, with a fake repository in service tests. Run `go test ./... -race`.
- Before finishing a change: `gofmt -l .` and `go vet ./...`.
