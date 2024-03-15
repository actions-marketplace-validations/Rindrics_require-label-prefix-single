package infra

import (
	"context"

	"github.com/Rindrics/require-label-prefix-on-closed/app"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	client *github.Client
}

func newGitHubClient(token string) *GitHubClient {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return &GitHubClient{client: github.NewClient(tc)}
}

type Client interface {
	postComment(owner, repo string, number int, body string) error
	addLabels(owner, repo string, number int, labels []string) error
}

func (g *GitHubClient) PostComment(info app.CommentInfo) error {
	comment := &github.IssueComment{Body: &info.Body}
	_, _, err := g.client.Issues.CreateComment(context.Background(), info.Owner, info.Repo, info.Number, comment)
	return err
}

func (g *GitHubClient) AddLabels(info app.LabelInfo) error {
	_, _, err := g.client.Issues.AddLabelsToIssue(context.Background(), info.Owner, info.Repo, info.Number, info.Labels)

	return err
}
