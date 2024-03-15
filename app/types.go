package app

type Command interface {
	Execute() error
}

type Labels []string

type EventInfo struct {
	Owner, Repo string
	Number      int
	Body        string
}

type Commenter interface {
	PostComment(EventInfo) error
}

type PostCommentCommand struct {
	EventInfo EventInfo
	Commenter Commenter
	onSuccess Action
}

type Labeler interface {
	AddLabels(EventInfo, Labels) error
}

type AddLabelsCommand struct {
	EventInfo EventInfo
	Labels    Labels
	Labeler   Labeler
	onSuccess PostCommentCommand
}
