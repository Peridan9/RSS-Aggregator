# RSS Aggregator 📡

## Overview 📖
The RSS Aggregator is a command-line tool written in Go that fetches and aggregates RSS feeds, allowing users to easily manage and read their subscribed content. This project was developed as part of a guided learning experience and serves as an example of a well-documented Go application.

## ✨ Features
- Fetch and aggregate RSS feeds
- User authentication and management
- Add, list, follow, and unfollow RSS feeds
- Browse new articles from followed feeds
- Stores feed data using PostgreSQL

## 🛠️ Technologies Used
- **Go** – Main programming language
- **PostgreSQL** – Database for storing feed data
- **Goose** - Database migration tool for building tables
- **SQLC** - Generates Go code for database queries
- **Go Modules** – Dependency management
- **CLI** – Command-line interface for user interaction
The RSS Aggregator is a command-line tool written in Go that fetches and aggregates RSS feeds, allowing users to easily manage and read their subscribed content. This project was developed as part of a guided learning experience and serves as an example of a well-documented Go application.

## Prerequisites 🛠️
Before running this program, ensure you have the following installed on your system:

- [Go](https://go.dev/doc/install) (version 1.18 or later recommended)
- [PostgreSQL](https://www.postgresql.org/download/) (for storing feed data)

## Installation 🚀
To install the `gator` CLI tool, run the following command:

```sh
go install github.com/Peridan9/RSS-Aggregator@latest
```

This will install the binary in your `$GOPATH/bin`, making it accessible as `gator`.

## Configuration ⚙️
Before running the aggregator, you need to configure the database connection. Create a `.env` file in the project root with the following variables:

```
DB_HOST=your_database_host
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_NAME=your_database_name
DB_PORT=your_database_port
```

## Usage 🎮
### Running in Development Mode
To run the application in development mode, use:

```sh
go run .
```

### Running in Production Mode
Once built, the application can be run as a standalone binary:

```sh
gator
```

## Commands 📜
The `gator` CLI provides various commands for managing RSS feeds and user interactions:

### Authentication & User Management
- **Login:**
  ```sh
  gator login
  ```
- **Register a new user:**
  ```sh
  gator register
  ```
- **Reset the users table:**
  ```sh
  gator reset
  ```
- **List all users:**
  ```sh
  gator users
  ```

### Feed Management
- **Add a new feed:**
  ```sh
  gator addfeed https://example.com/rss
  ```
- **List all feeds:**
  ```sh
  gator feeds
  ```
- **Aggregate feed data:**
  ```sh
  gator agg
  ```

### Follow & Unfollow Feeds
- **Follow a feed:**
  ```sh
  gator follow https://example.com/rss
  ```
- **List followed feeds:**
  ```sh
  gator following
  ```
- **Unfollow a feed:**
  ```sh
  gator unfollow https://example.com/rss
  ```

### Browsing Feeds
- **Browse new articles from followed feeds:**
  ```sh
  gator browse
  ```

## Deployment 📦
Once installed and configured, the `gator` binary can be used on any system without requiring the Go toolchain. To build a production binary, run:

```sh
go build -o gator
```

This creates an executable `gator` that can be run independently.

