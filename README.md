# todo

A command-line todo app written in Go. It uses [Cobra](https://github.com/spf13/cobra) for the CLI, [GORM](https://gorm.io) for database access, and PostgreSQL (run with Podman) for storage.

## High-Level Design

![Todo CLI HLD](docs/hld.png)

1. The user runs a command through the Cobra CLI.
2. The CLI calls the service layer (Add, Delete, List, Toggle, Update).
3. The service reads and writes todos in PostgreSQL.
4. The result or error goes back to the CLI, which prints it.

## Project structure

```text
cmd/todo/main.go        Entry point: connects the layers together
internal/model/         Todo struct (database schema) and errors
internal/database/      Opens the connection and creates the table
internal/repository/    Database calls (the only layer that uses GORM)
internal/service/       Business rules (validation, toggle logic)
internal/handler/       Cobra commands: add, list, toggle, update, delete
```

Each layer only talks to the one below it: `handler` → `service` → `repository` → database.

## Requirements

- [Go](https://go.dev/dl/) 1.27 or newer
- [Podman](https://podman.io/) with `podman compose` (on macOS: `brew install podman podman-compose`)

## Setup

### 1. Start Podman (macOS/Windows only)

```bash
podman machine start
```

### 2. Create your `.env` file

```bash
cp .env.example .env
```

Open `.env` and set your own `POSTGRES_PASSWORD`. Then update the password in `DATABASE_URL` to match. The app reads `DATABASE_URL` from this file automatically.

> If your password has special characters (such as `/` or `@`), URL-encode it in `DATABASE_URL` (`/` becomes `%2F`, `@` becomes `%40`). A password with only letters and numbers avoids this.

### 3. Create the data folder

The database files are stored on your Desktop, so you can delete them by moving the folder to Trash.

```bash
mkdir -p ~/Desktop/volumes/todo/data
```

### 4. Start the database

```bash
podman compose up -d
podman compose ps
```

Wait until the status shows `healthy`.

### 5. Build the app

```bash
go build -o bin/todo ./cmd/todo
```

The `todos` table is created automatically the first time you run the app.

## Usage

```bash
./bin/todo add "buy milk"
./bin/todo list
./bin/todo toggle 1
./bin/todo update 1 "buy oat milk"
./bin/todo delete 1
```

Example output of `list`:

```text
[x] 1   buy oat milk
[ ] 2   call mom
```

| Command | What it does |
| --- | --- |
| `add <title>` | Add a todo |
| `list` | Show all todos (`[x]` means done) |
| `toggle <id>` | Mark done, or not done if it already is |
| `update <id> <title>` | Change the title |
| `delete <id>` (or `rm`) | Delete a todo |

Run `./bin/todo --help` for details.

## Tests

```bash
go test ./... -race
```

The tests use a fake repository, so they do not need the database.

## Database commands

| Command | What it does |
| --- | --- |
| `podman compose up -d` | Create and start the database in the background |
| `podman compose ps` | Show the status |
| `podman compose logs -f` | Show the logs (press `Ctrl+C` to exit) |
| `podman compose stop` | Pause the database |
| `podman compose start` | Resume a paused database |
| `podman compose down` | Remove the container (your data is kept) |

## Where is my data?

In `~/Desktop/volumes/todo/data`, outside the container. `podman compose down` does not delete it.

To delete everything, run `podman compose down`, then move `~/Desktop/volumes/todo` to Trash.
