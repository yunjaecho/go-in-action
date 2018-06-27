package search

// 기본 검색기를 구현할 defaultMatcher 타입
type defaultMatcher struct {}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func (m defaultMatcher) Search(feed* Feed, searchTerm string)([]* Result, error) {
	return nil, nil
}
