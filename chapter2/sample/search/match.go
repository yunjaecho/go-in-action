package search

import "log"

type Result struct {
	Field string
	Content string
}

type Matcher interface {
	Search(feed* Feed, searchTerm string)([]* Result, error)
}

func Match(matcher Matcher, feed* Feed, searchTerm string, results chan<- *Result) {
	searchResult, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 검색 결과를 체널에 기록한다.
	for _, result := range searchResult {
		results <- result
	}
}

// 함수는 개별 고루틴이 전달할
// 검색 결과를 콘솔 창에 출력한다.
func Display(results chan* Result) {
	// 채널은 검색 결과가 기록될 따까지 접근이 차단된다.
	//채널이 닫히면 닫히면 루프가 종료된다.
	for result := range results {
		log.Println("%s:\n%s\n\n", result.Field, result.Content)
	}
}