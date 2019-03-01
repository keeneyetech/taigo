# Taiga API client written in Go

This library doesn't cover all the API, so contributions are welcome.

### Generate auth token

Run the script `gen_token.sh` to get e fresh auth token from your credentials.

### Usage

Like many other Go API clients, this library follows the same pattern as
`go-github`:

```go
// Create a client
client := taigo.NewClient("YOUR-TAIGA-API-URL", "YOUR-AUTH-TOKEN")

// Get a project for his slug name
project, _, err := client.Project.GetBySlug("user-project")

// List userstories of the project
opts := taigo.UserStoryListOptions{ ProjectID: &project.ID }
userstories, _, err := client.UserStory.List(opts)
```
