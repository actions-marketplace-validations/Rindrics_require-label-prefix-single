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

func (g *GitHubClient) PostComment(p app.PostCommentParams) error {
	body := "comment body"
	comment := &github.IssueComment{Body: &body}
	_, _, err := g.client.Issues.CreateComment(context.Background(), p.RepoInfo.Owner, p.RepoInfo.Repo, p.Number, comment)
	return err
}

func (g *GitHubClient) AddLabels(p app.AddLabelsParams, labels app.Labels) error {
	_, _, err := g.client.Issues.AddLabelsToIssue(context.Background(), p.RepoInfo.Owner, p.RepoInfo.Repo, p.Number, labels)

	return err
}
