package arrayimpl

type ISensitiveWordFilter interface {
	DoFilter(content string) bool
}

type SexyWordFilter struct {
}

func (s *SexyWordFilter) DoFilter(content string) bool {
	legal := true
	// ...
	return legal
}

type PoliticalWordFilter struct {
}

func (p *PoliticalWordFilter) DoFilter(content string) bool {
	legal := true
	// ...
	return legal
}

type AdsWordFilter struct {
}

func (a *AdsWordFilter) DoFilter(content string) bool {
	legal := true
	// ...
	return legal
}

type SensitiveWordFilterChain struct {
	filters []ISensitiveWordFilter
}

func (s *SensitiveWordFilterChain) AddFilter(filter ISensitiveWordFilter) {
	s.filters = append(s.filters, filter)
}

func (s *SensitiveWordFilterChain) Filter(content string) bool {
	for _, filter := range s.filters {
		if !filter.DoFilter(content) {
			return false
		}
	}
	return true
}
