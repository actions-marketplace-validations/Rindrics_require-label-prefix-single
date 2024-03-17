package application

import (
	"github.com/Rindrics/require-label-prefix/domain"
)

type Command interface {
	Execute() error
}

type App struct {
	Command Command
	Logger  Logger
}

type Logger interface {
	Debug(string, ...any)
	Info(string, ...any)
	Error(string, ...any)
}

type PostCommentParams struct {
	RepoInfo domain.RepoInfo
	Number   int
	Body     string
}

type Commenter interface {
	PostComment(PostCommentParams) error
}

type PostCommentCommand struct {
	Params    PostCommentParams
	Commenter Commenter
	OnSuccess Action
}

type Labeler interface {
	AddLabels(AddLabelsParams) error
}

type AddLabelsParams struct {
	RepoInfo domain.RepoInfo
	Number   int
	Labels   domain.Labels
}

type AddLabelsCommand struct {
	Params    AddLabelsParams
	Labeler   Labeler
	OnSuccess PostCommentCommand
}

type Action interface {
	Perform() error
}

type ExitAction struct{}

type GitHubClient interface {
	PostComment(params PostCommentParams) error
	AddLabels(params AddLabelsParams) error
}
