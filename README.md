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
Recent activity:

- Pushed to saurabhdhingra/github-activity
Date: 2025-02-06 10:38:36
URL: https://github.com/saurabhdhingra/github-activity
Commit Message: Changes to meta data being displayed.

- Pushed to saurabhdhingra/CSES-Solutions
Date: 2025-02-04 13:50:04
URL: https://github.com/saurabhdhingra/CSES-Solutions
Commit Message: Segregating sections
```

## Error Handling
- Handles invalid GitHub usernames.
- Handles API failures and rate limits.

## Acknowledgement
https://roadmap.sh/projects/github-user-activity



