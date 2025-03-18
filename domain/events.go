package domain

// PullRequestEvent define el evento de un Pull Request de GitHub.
type PullRequestEvent struct {
	Title  string
	Number int
	User   string
	Action string
}

// ActionsEvent define el evento de GitHub Actions.
type ActionsEvent struct {
	Workflow   string
	Action     string
	Conclusion string
}