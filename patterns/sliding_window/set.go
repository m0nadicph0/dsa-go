package sliding_window

type Set struct {
	hashMap map[string]bool
}

func NewSet() *Set {
	return &Set{hashMap: make(map[string]bool)}
}

func (s *Set) Add(v string) {
	s.hashMap[v] = true
}
func (s *Set) Delete(v string) {
	delete(s.hashMap, v)
}

func (s *Set) Exists(v string) bool {
	_, ok := s.hashMap[v]
	return ok
}

func (s *Set) All() []string {
	result := make([]string, 0)
	for k, _ := range s.hashMap {
		result = append(result, k)
	}
	return result
}
