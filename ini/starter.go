package ini

type IStarterInterface interface {
	SetUp()
}

type starter struct {
	elements []IStarterInterface
}

func (s *starter) SetUp() {

	if len(s.elements) == 0 {
		return
	}

	for i := 0; i < len(s.elements); i++ {
		s.elements[i].SetUp()
	}
}

func (s *starter) Register(starter IStarterInterface) {
	s.elements = append(s.elements, starter)
}

func NewStarter() *starter {
	return &starter{}
}
