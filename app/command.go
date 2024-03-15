package app

func (c *PostCommentCommand) Execute() error {
	err := c.Commenter.PostComment(c.EventInfo)
	if err != nil {
		return err
	}

	return c.onSuccess.Perform()
}

func (c *PostCommentCommand) Perform() error {
	return c.Execute()
}

func (c *AddLabelsCommand) Execute() error {
	if err := c.Labeler.AddLabels(c.EventInfo, c.Labels); err != nil {
		return err
	}

	return c.onSuccess.Execute()
}
