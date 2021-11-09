package basicBlocks

import "fmt"

type Set map[string]struct{}

func Index() {
	s := make(Set)
	s["item1"] = struct{}{}
	s["item2"] = struct{}{}

	fmt.Println(s.getSetValues())
}

func (s Set) getSetValues() []string {
	var response []string
	for k := range s {
		response = append(response, k)
	}
	return response
}
