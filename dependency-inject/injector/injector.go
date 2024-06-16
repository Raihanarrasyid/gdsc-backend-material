package injector

type SimpleRepository struct {
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	return &SimpleService{
		SimpleRepository: repository,
	}, nil
}