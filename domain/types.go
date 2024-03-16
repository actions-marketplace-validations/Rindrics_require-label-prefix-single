package domain

type Labels []string

type EventInfo struct {
	Number int
	Labels Labels
}

type RepoInfo struct {
	Owner, Repo string
}
