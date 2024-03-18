package infra

import (
	"context"
	"log/slog"

	"github.com/Rindrics/require-label-prefix-single/application"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	client *github.Client
	logger *slog.Logger
}

func NewGitHubClient(token string, logger *slog.Logger) *GitHubClient {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return &GitHubClient{client: github.NewClient(tc), logger: logger}
}

func (g *GitHubClient) PostComment(p application.PostCommentParams) error {
	g.logger.Debug("posting comment", "body", p)
	body := "comment body"
	comment := &github.IssueComment{Body: &body}
	_, _, err := g.client.Issues.CreateComment(context.Background(), p.RepoInfo.Owner, p.RepoInfo.Repo, p.Number, comment)
	return err
}

func (g *GitHubClient) AddLabels(p application.AddLabelsParams) error {
	g.logger.Debug("adding labels", "labels", p)
	_, _, err := g.client.Issues.AddLabelsToIssue(context.Background(), p.RepoInfo.Owner, p.RepoInfo.Repo, p.Number, p.Labels)

	return err
}
