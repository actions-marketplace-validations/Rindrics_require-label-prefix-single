package app

type Command interface {
	Execute() error
}

type RepoInfo struct {
	Owner, Repo string
}

type Labels []string

type EventInfo struct {
	Number int
	Labels Labels
}

type PostCommentParams struct {
	RepoInfo RepoInfo
	Number   int
	Body     string
}

type Commenter interface {
	PostComment(PostCommentParams) error
}

type PostCommentCommand struct {
	Params    PostCommentParams
	Commenter Commenter
	onSuccess Action
}

type Labeler interface {
	AddLabels(AddLabelsParams) error
}

type AddLabelsParams struct {
	RepoInfo RepoInfo
	Number   int
	Labels   Labels
}

type AddLabelsCommand struct {
	Params    AddLabelsParams
	Labeler   Labeler
	onSuccess PostCommentCommand
}
