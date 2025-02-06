# GitHub User Activity Fetcher

A CLI tool to fetch and display recent GitHub activity for a given user.

## Features
- Fetch recent activity from GitHub for a specified user.
- Display events such as commits, issues, and stars.
- Handle API errors gracefully.

## Installation
```sh
git clone https://github.com/yourusername/github-activity.git
cd github-activity
go build -o github-activity
```

## Usage

### Fetch User Activity
```sh
./github-activity <username>
```

### Example Output
```
- Pushed 3 commits to repo-name
- Opened an issue in repo-name
- Starred repo-name
```

## Error Handling
- Handles invalid GitHub usernames.
- Handles API failures and rate limits.



