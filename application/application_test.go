package application

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestRun(t *testing.T) {
	t.Run("add label and then post comment", func(t *testing.T) {
		mockedLabeler := new(MockLabeler)
		mockedLabeler.On("AddLabels", mock.Anything).Return(nil)
		mockedCommenter := new(MockCommenter)
		mockedCommenter.On("PostComment", mock.Anything).Return(nil)
		mockedLogger := &MockLogger{}
		mockedLogger.On("Info", "start executing command", mock.Anything).Once()
		mockedLogger.On("Info", "command executed successfully", mock.Anything).Once()

		app := App{
			Command: AddLabelsCommand{
				Labeler: mockedLabeler,
				Params:  AddLabelsParams{},
				OnSuccess: PostCommentCommand{
					Commenter: mockedCommenter,
					OnSuccess: &ExitAction{},
				},
			},
			Logger: mockedLogger,
		}

		app.Run()

		mockedLabeler.AssertExpectations(t)
		mockedCommenter.AssertExpectations(t)
		mockedLogger.AssertExpectations(t)
	})
}
