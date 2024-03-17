package application

func (c PostCommentCommand) Execute() error {
	err := c.Commenter.PostComment(c.Params)
	if err != nil {
		return err
	}

	return c.OnSuccess.Perform()
}

func (c PostCommentCommand) Perform() error {
	return c.Execute()
}

func (c AddLabelsCommand) Execute() error {
	if err := c.Labeler.AddLabels(c.Params); err != nil {
		return err
	}

	return c.OnSuccess.Execute()
}
