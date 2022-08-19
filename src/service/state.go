package service

type StateService interface {
	Create(StateRequest) (*StateRequest, error)
	Get()
	Update()
	Delete()
	List(string, string) ([]StateRequest, error)
}
type stateService struct{}

type StateRequest struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewStateService() StateService {
	return &stateService{}
}
func (s *stateService) Create(req StateRequest) (*StateRequest, error) {

	return nil, nil
}

func (s *stateService) Get() {

	return
}

func (s *stateService) Update() {

	return
}

func (s *stateService) Delete() {

	return
}

func (s *stateService) List(page string, size string) ([]StateRequest, error) {
	stateList := []StateRequest{}
	return stateList, nil
}
