package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const githubAPI = "https://api.github.com/users/%s/events"

func FetchActivity(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}

	url := fmt.Sprintf(githubAPI, username)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch activity: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return errors.New("user not found")
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected response from GitHub API: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	var events []map[string]interface{}
	if err := json.Unmarshal(body, &events); err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	if len(events) == 0 {
		fmt.Println("No recent activity found for this user.")
		return nil
	}

	fmt.Println("\n\nRecent activity:\n")
	for _, event := range events {
		if action, description, date, url, message := parseEvent(event); action != "" {
			fmt.Printf("- %s %s\n", action, description)
			fmt.Printf("  Date: %s\n", date)
			fmt.Printf("  URL: %s\n", url)
			if message != "" {
				fmt.Printf("  Commit Message: %s\n\n", message)
			}
		}
	}

	return nil
}

func parseEvent(event map[string]interface{}) (action, description, date, url, message string) {
	eventType, ok := event["type"].(string)
	if !ok {
		return "", "", "", "", ""
	}

	repo, ok := event["repo"].(map[string]interface{})
	repoName := ""
	if ok {
		repoName, _ = repo["name"].(string)
	}

	// Extract the date
	createdAt, _ := event["created_at"].(string)
	if createdAt != "" {
		parsedDate, err := time.Parse(time.RFC3339, createdAt)
		if err == nil {
			date = parsedDate.Format("2006-01-02 15:04:05")
		}
	}

	// Construct the URL
	url = fmt.Sprintf("https://github.com/%s", repoName)

	// Extract the commit message for PushEvent
	if eventType == "PushEvent" {
		payload, _ := event["payload"].(map[string]interface{})
		commits, _ := payload["commits"].([]interface{})
		if len(commits) > 0 {
			commit, _ := commits[0].(map[string]interface{})
			message, _ = commit["message"].(string)
		}
	}

	switch eventType {
	case "PushEvent":
		action = "Pushed to"
		description = repoName
	case "IssuesEvent":
		payload, _ := event["payload"].(map[string]interface{})
		action = "Opened a new issue in"
		description = repoName
		if issue, ok := payload["issue"].(map[string]interface{}); ok {
			title, _ := issue["title"].(string)
			description = fmt.Sprintf("%s: %s", repoName, title)
			url, _ = issue["html_url"].(string)
		}
	case "WatchEvent":
		action = "Starred"
		description = repoName
	default:
		action = ""
		description = ""
	}

	return action, description, date, url, message
}
