Gator

Gator is a command-line RSS feed aggregator written in Go.

It allows multiple users to register, add and follow RSS feeds, periodically fetch posts from those feeds, and browse stored posts directly from the command line.

This project was built as part of the Boot.dev backend development curriculum.

Requirements

* Go
* PostgreSQL
* Goose for database migrations
* SQLC for generating type-safe database queries

Setup

Clone the repository:

git clone https://github.com/goczangabor24/gator.git
cd gator

Create a PostgreSQL database for Gator.

Then create a .gatorconfig.json file in your home directory:

{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}

Replace the database URL with your own PostgreSQL connection string.

Run the database migrations and generate the SQLC code if necessary.

Build the application:

go build

Usage

Commands are executed using:

./gator <command> [arguments]

User commands

Register a new user:

./gator register <username>

Log in as an existing user:

./gator login <username>

List users:

./gator users

Feed commands

Add a new RSS feed:

./gator addfeed <name> <url>

List available feeds:

./gator feeds

Follow a feed:

./gator follow <url>

List feeds followed by the current user:

./gator following

Unfollow a feed:

./gator unfollow <url>

Aggregation

Start fetching RSS feeds at a specified interval:

./gator agg <interval>

For example:

./gator agg 30s

Gator periodically fetches RSS feeds and stores new posts in PostgreSQL.

Browse posts

Browse posts from followed feeds:

./gator browse

Tech Stack

* Go — application and CLI
* PostgreSQL — persistent data storage
* SQLC — type-safe Go code generated from SQL
* Goose — database migrations
* RSS/XML — feed parsing

Project Structure

gator/
├── internal/
│   ├── config/       # Configuration management
│   └── database/     # SQLC-generated database code
├── sql/
│   ├── queries/      # SQL queries
│   └── schema/       # Goose database migrations
├── main.go           # Application entry point
├── commands.go       # CLI command registration
├── fetchFeed.go      # RSS fetching and parsing
├── scrapeFeeds.go    # Feed aggregation
└── handler*.go       # Individual CLI command handlers
