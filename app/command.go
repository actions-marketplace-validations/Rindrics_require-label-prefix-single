package app

type Command interface {
	Execute() error
}

type CommentInfo struct {
	owner, repo string
	number      int
	body        string
}

type Commenter interface {
	PostComment(CommentInfo) error
}

type PostCommentCommand struct {
	CommentInfo CommentInfo
	Commenter   Commenter
	onSuccess   Action
}

func (c *PostCommentCommand) Execute() error {
	err := c.Commenter.PostComment(c.CommentInfo)
	if err != nil {
		return err
	}

	return c.onSuccess.Perform()
}

func (c *PostCommentCommand) Perform() error {
	return c.Execute()
}

type LabelInfo struct {
	owner, repo string
	number      int
	labels      []string
}

type Labeler interface {
	AddLabels(info LabelInfo) error
}

type AddLabelsCommand struct {
	LabelInfo LabelInfo
	Labeler   Labeler
	onSuccess PostCommentCommand
}

func (c *AddLabelsCommand) Execute() error {
	if err := c.Labeler.AddLabels(c.LabelInfo); err != nil {
		return err
	}

	return c.onSuccess.Execute()
}
