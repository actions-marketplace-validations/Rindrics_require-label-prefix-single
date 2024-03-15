package infra

import (
	"testing"
)

type mockCommenter struct{}

func (c *mockCommenter) postComment(owner, repo string, number int, body string) error {
	return nil
}

func TestCommenter(t *testing.T) {
	mockCommenter := &mockCommenter{}
	owner := "octocat"
	repo := "sandbox"
	number := 1
	body := "test comment"

	err := mockCommenter.postComment(owner, repo, number, body)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
