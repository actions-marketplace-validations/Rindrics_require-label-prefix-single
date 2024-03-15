package app

type Action interface {
	Perform() error
}

type ExitAction struct{}

func (a *ExitAction) Perform() error {
	return nil
}
