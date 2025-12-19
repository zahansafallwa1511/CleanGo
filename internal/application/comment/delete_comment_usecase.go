package comment

type DeleteCommentUseCase struct {
	// TODO: Add repository dependency
}

func NewDeleteCommentUseCase() *DeleteCommentUseCase {
	return &DeleteCommentUseCase{}
}

func (uc *DeleteCommentUseCase) Execute(id string) error {
	if id == "0" {
		return ErrCommentNotFound
	}
	// TODO: Call repository
	return nil
}
