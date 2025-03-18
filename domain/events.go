package domain


type PullRequestEvent struct {
	Title  string
	Number int
	User   string
	Action string
}


type ActionsEvent struct {
	Workflow   string
	Action     string
	Conclusion string
}
