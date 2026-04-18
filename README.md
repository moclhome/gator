# Gator

Gator is a command-line tool built in Go for managing RSS feeds.

The idea is that the command "gator agg <duration>" can be run in the background, selecting posts of feeds the user is following and storing them in a db. The posts can be afterwards viewed with the "gator browse" command (with an optional limit parameter).

## Features

- Add and manage RSS feed subscriptions
- Fetch and display feed articles
- Track feed reading progress
- Browse feeds from the command line (optionally read them including the description)

## Installation

You need Postgresql and Go installed to run the programme.

You can use Goose to migrate to the current database schema version.

```bash
go install github.com/bootdev/gator@latest
```

## Usage

```bash
gator addfeed <feed_name> <feed_url> # Add the feed to the database and follow it.         
gator agg  <duration>                # Start scraping all saved feeds for new posts. If found, they are added to the database. The cpommand "agg 3m" would search for posts every 3 minutes. To stop, use Ctrl-C.
gator browse [limit]                 # List the latest posts of feeds you are following. The optional "limit" limits the displayed posts. This command only displays title, publish date and url.
gator feeds                          # List all saved feeds plus information who added them
gator follow <feed_url>              # Subscribe to a feed
gator following                      # List all feeds you are following
gator help                           # List all supported command and their descriptions
gator login  <user_name>             # Set the current user to "user_name"
gator open <post_url>                # Opens the default browserand displays the post_url.
gator register <user_name>           # Adds "user_name" to the list of users and sets the current user to this user.
gator read [limit]                   # List the latest posts of feeds you are following. The optional "limit" limits the displayed posts. Compared to "browse", this command additionally  displays the description.
gator reset                          # Deletes all feeds, posts and users. ATTENTION! All data will be lost!
gator unfollow <feed_url>            # Unsubscribe a feed
gator users                          # List all users and the information which is the current one
```

## Configuration

Gator stores configuration in `~/.gatorconfig.json`. This file must contain the database connection string. The current_user_name is set later when calling "gator login". 
{"db_url":"<db-connection-string>?sslmode=disable","current_user_name":""}

## Development

```bash
go build
go test ./...
```

## License

MIT
