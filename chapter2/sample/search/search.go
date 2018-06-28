package search

import (
	"log"	// log 패키지는 표준 출력, 표준오류 또는 사용자정의 장치를 통해 로깅 메세지를 출력
	"sync"	// 고루틴 사이의 동기화
)

// 겸색을 처리할 검색기의 매핑 정보를 저장할
// make 변수의 초기화
// map 은 make 함수를 이용해 GO 런타임에 생성을 요청해야 하는 참조타입
// map[string]Matcher == java Map<String, Matcher>
var matchers = make(map[string]Matcher)

// 검색 로직을 수행한 Run함수
func Run(searchTerm string) {
	// 검색할 피드의 목록을 조회한다.
	// := 변수의 선언과 초기화 동시처리
	feeds, err := RetriveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 버퍼가 없는 채널을 생성하여 화면에 표시할 검색 결과를 전달받는다.
	results := make(chan *Result)

	// 모든 피드를 처리할 때까지 기달릴 대기 그룹(Wait group)을 설정한다.
	var waitGroup sync.WaitGroup

	// 개별 피드를 처리하는 동안 대기해야 할
	// 고루틴의 개수를 설정한다.
	waitGroup.Add(len(feeds))

	// 각기 다른 종류의 피드를 처리할 고루틴을 실행한다.
	for _, feed := range feeds {
		// 검색을 위해 검색기를 조회한다.
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 검색을 실행하기 위해 고루틴을 실행한다.
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		} (matcher, feed)
	}

	// 모든 작업이 완료되었는지를 모니터링할 고루틴을 실행한다.
	go func() {
		// 모든 작업이 처리될 때까지 기다린다.
		waitGroup.Wait()
		// Display 함수에게 프로그램을 종료할 수 있음을
		// 알리기 위해 채널을 닫는다.
		close(results)
	}()

	// 검색 결과를 화면에 표시한다.
	// 마지막 결과를 표시한 뒤 리턴한다.
	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "검색기가 이미 등록되었습니다.")
	}

	log.Println("등록완료", feedType, "검색기")
	matchers[feedType] = matcher
}