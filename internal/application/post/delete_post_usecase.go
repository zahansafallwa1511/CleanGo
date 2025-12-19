package post

type DeletePostUseCase struct {
	// TODO: Add repository dependency
	// repo IPostRepository
}

func NewDeletePostUseCase() *DeletePostUseCase {
	return &DeletePostUseCase{}
}

func (uc *DeletePostUseCase) Execute(id string) error {
	if id == "0" {
		return ErrPostNotFound
	}
	// TODO: Call repository
	return nil
}
