package infra

import (
	"context"

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

func (g *GitHubClient) postComment(owner, repo string, number int, body string) error {
	comment := &github.IssueComment{Body: &body}

	_, _, err := g.client.Issues.CreateComment(context.Background(), owner, repo, number, comment)
	return err
}

func (g *GitHubClient) addLabels(owner, repo string, number int, labels []string) error {
	_, _, err := g.client.Issues.AddLabelsToIssue(context.Background(), owner, repo, number, labels)

	return err
}
