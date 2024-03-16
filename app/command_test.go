package app

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockExitAction struct {
	mock.Mock
}

func (m *MockExitAction) Perform() error {
	args := m.Called()
	return args.Error(0)
}

type MockCommenter struct {
	mock.Mock
	PostCommentParams
}

func (m *MockCommenter) PostComment(p PostCommentParams) error {
	args := m.Called(p)
	return args.Error(0)
}

type MockLabeler struct {
	mock.Mock
	AddLabelsParams
}

func (m *MockLabeler) AddLabels(p AddLabelsParams) error {
	args := m.Called(p)
	return args.Error(0)
}

type MockCommand struct {
	mock.Mock
	OnSuccess MockExitAction
}

func (m *MockCommand) Execute() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockCommand) Perform() error {
	return m.Execute()
}

func TestPostCommentCommand(t *testing.T) {
	mockCommenter := new(MockCommenter)
	mockExitAction := new(MockExitAction)

	mockCommenter.On("PostComment", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockExitAction.On("Perform").Return(nil)

	command := PostCommentCommand{
		Commenter: mockCommenter,
		OnSuccess: mockExitAction,
	}
	err := command.Execute()

	mockCommenter.AssertExpectations(t)
	mockExitAction.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestPostCommentCommandError(t *testing.T) {
	mockCommenter := new(MockCommenter)
	mockExitAction := new(MockExitAction)

	expectedError := errors.New("failed to post comment")
	mockCommenter.On("PostComment", mockCommenter.PostCommentParams).Return(expectedError)
	mockExitAction.On("Perform").Return(nil)

	command := PostCommentCommand{
		Commenter: mockCommenter,
		OnSuccess: mockExitAction,
	}
	err := command.Execute()

	mockCommenter.AssertExpectations(t)
	mockExitAction.AssertNotCalled(t, "Perform")
	assert.ErrorIs(t, err, expectedError)
}

func TestAddLabelsCommandWithPostCommentOnSuccess(t *testing.T) {
	mockLabeler := new(MockLabeler)
	mockCommenter := new(MockCommenter)
	mockExitAction := new(MockExitAction)

	mockLabeler.On("AddLabels", mockLabeler.AddLabelsParams).Return(nil)
	mockCommenter.On("PostComment", mockCommenter.PostCommentParams).Return(nil)
	mockExitAction.On("Perform").Return(nil)

	postCommentCommand := PostCommentCommand{
		Commenter: mockCommenter,
		OnSuccess: mockExitAction,
	}

	addLabelCommand := AddLabelsCommand{
		Params:    mockLabeler.AddLabelsParams,
		Labeler:   mockLabeler,
		OnSuccess: postCommentCommand,
	}

	err := addLabelCommand.Execute()

	mockLabeler.AssertExpectations(t)
	mockCommenter.AssertExpectations(t)
	mockExitAction.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestAddLabelsCommandWithPostCommentOnError(t *testing.T) {
	mockLabeler := new(MockLabeler)
	mockCommenter := new(MockCommenter)
	mockExitAction := new(MockExitAction)

	expectedError := errors.New("failed to add label")
	mockLabeler.On("AddLabels", mockLabeler.AddLabelsParams).Return(expectedError)
	mockCommenter.On("PostComment", mockCommenter.PostCommentParams).Return(nil)
	mockExitAction.On("Perform").Return(nil)

	postCommentCommand := PostCommentCommand{
		Commenter: mockCommenter,
		OnSuccess: mockExitAction,
	}

	addLabelCommand := AddLabelsCommand{
		Params:    mockLabeler.AddLabelsParams,
		Labeler:   mockLabeler,
		OnSuccess: postCommentCommand,
	}

	err := addLabelCommand.Execute()

	mockLabeler.AssertExpectations(t)
	mockCommenter.AssertNotCalled(t, "PostComment")
	mockExitAction.AssertNotCalled(t, "Perform")
	assert.ErrorIs(t, err, expectedError)
}
