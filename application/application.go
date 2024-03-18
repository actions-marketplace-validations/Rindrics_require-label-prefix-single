package application

import (
	"github.com/Rindrics/require-label-prefix-single/domain"
)

func New(info domain.EventInfo, client GitHubClient, config Config, logger Logger) App {
	if config.AddLabel {
		return App{
			Command: AddLabelsCommand{
				Labeler: client,
				Params: AddLabelsParams{
					Number: info.Number,
					Labels: domain.Labels{config.DefaultLabel},
				},
				OnSuccess: PostCommentCommand{
					Commenter: client,
					Params: PostCommentParams{
						RepoInfo: domain.RepoInfo{
							Owner: config.Owner,
							Repo:  config.Repository,
						},
						Number: info.Number,
						Body:   "Label added",
					},
					OnSuccess: &ExitAction{},
				},
			},
			Logger: logger,
		}
	} else {
		return App{
			Command: PostCommentCommand{
				Commenter: client,
				Params: PostCommentParams{
					RepoInfo: domain.RepoInfo{
						Owner: config.Owner,
						Repo:  config.Repository,
					},
					Number: info.Number,
					Body:   config.Comment,
				},
				OnSuccess: &ExitAction{},
			},
			Logger: logger,
		}
	}
}

func (a App) Run() error {
	a.Logger.Info("start executing command")
	err := a.Command.Execute()
	if err != nil {
		a.Logger.Error("Error executing command", err)
		return err
	}
	a.Logger.Info("command executed successfully")
	return nil
}
