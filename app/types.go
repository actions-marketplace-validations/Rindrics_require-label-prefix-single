package app

type Command interface {
	Execute() error
}

type CommentInfo struct {
	Owner, Repo string
	Number      int
	Body        string
}

type Commenter interface {
	PostComment(CommentInfo) error
}

type PostCommentCommand struct {
	CommentInfo CommentInfo
	Commenter   Commenter
	onSuccess   Action
}

type LabelInfo struct {
	Owner, Repo string
	Number      int
	Labels      []string
}

type Labeler interface {
	AddLabels(info LabelInfo) error
}

type AddLabelsCommand struct {
	LabelInfo LabelInfo
	Labeler   Labeler
	onSuccess PostCommentCommand
}
