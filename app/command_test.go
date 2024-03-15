package app

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAction struct {
	mock.Mock
}

func (m *MockAction) Perform() error {
	args := m.Called()
	return args.Error(0)
}

type MockCommenter struct {
	mock.Mock
	EventInfo
}

func (m *MockCommenter) PostComment(c EventInfo) error {
	args := m.Called(c)
	return args.Error(0)
}

type MockLabeler struct {
	mock.Mock
	EventInfo
	Labels
}

func (m *MockLabeler) AddLabels(e EventInfo, l Labels) error {
	args := m.Called(e, l)
	return args.Error(0)
}

type MockCommand struct {
	mock.Mock
	onSuccess Action
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
	mockAction := new(MockAction)

	mockCommenter.On("PostComment", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockAction.On("Perform").Return(nil)

	command := PostCommentCommand{
		Commenter: mockCommenter,
		onSuccess: mockAction,
	}
	err := command.Execute()

	mockCommenter.AssertExpectations(t)
	mockAction.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestPostCommentCommandError(t *testing.T) {
	mockCommenter := new(MockCommenter)
	mockAction := new(MockAction)

	expectedError := errors.New("failed to post comment")
	mockCommenter.On("PostComment", mockCommenter.EventInfo).Return(expectedError)
	mockAction.On("Perform").Return(nil)

	command := PostCommentCommand{
		Commenter: mockCommenter,
		onSuccess: mockAction,
	}
	err := command.Execute()

	mockCommenter.AssertExpectations(t)
	mockAction.AssertNotCalled(t, "Perform")
	assert.ErrorIs(t, err, expectedError)
}

func TestAddLabelsCommandWithPostCommentOnSuccess(t *testing.T) {
	mockLabeler := new(MockLabeler)
	mockCommenter := new(MockCommenter)
	mockAction := new(MockAction)

	mockLabeler.On("AddLabels", mockLabeler.EventInfo, mockLabeler.Labels).Return(nil)
	mockCommenter.On("PostComment", mockCommenter.EventInfo).Return(nil)
	mockAction.On("Perform").Return(nil)

	postCommentCommand := PostCommentCommand{
		Commenter: mockCommenter,
		onSuccess: mockAction,
	}

	addLabelCommand := AddLabelsCommand{
		EventInfo: mockLabeler.EventInfo,
		Labels:    mockLabeler.Labels,
		Labeler:   mockLabeler,
		onSuccess: postCommentCommand,
	}

	err := addLabelCommand.Execute()

	mockLabeler.AssertExpectations(t)
	mockCommenter.AssertExpectations(t)
	mockAction.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestAddLabelsCommandWithPostCommentOnError(t *testing.T) {
	mockLabeler := new(MockLabeler)
	mockCommenter := new(MockCommenter)
	mockAction := new(MockAction)

	expectedError := errors.New("failed to add label")
	mockLabeler.On("AddLabels", mockLabeler.EventInfo, mockLabeler.Labels).Return(expectedError)
	mockCommenter.On("PostComment", mockCommenter.EventInfo).Return(nil)
	mockAction.On("Perform").Return(nil)

	postCommentCommand := PostCommentCommand{
		Commenter: mockCommenter,
		onSuccess: mockAction,
	}

	addLabelCommand := AddLabelsCommand{
		EventInfo: mockLabeler.EventInfo,
		Labels:    mockLabeler.Labels,
		Labeler:   mockLabeler,
		onSuccess: postCommentCommand,
	}

	err := addLabelCommand.Execute()

	mockLabeler.AssertExpectations(t)
	mockCommenter.AssertNotCalled(t, "PostComment")
	mockAction.AssertNotCalled(t, "Perform")
	assert.ErrorIs(t, err, expectedError)
}
