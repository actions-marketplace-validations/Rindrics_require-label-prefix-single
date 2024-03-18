package application

import "github.com/stretchr/testify/mock"

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
