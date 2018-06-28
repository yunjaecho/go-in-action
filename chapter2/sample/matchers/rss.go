package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"../search"
	"log"
	"regexp"
)

type (
	// item 구조체는 RSS 문서 내의 item 태크에
	// 정의 된 필드들에 대응하는 필드들을 선언한다.
	item struct {
		XMLName 	xml.Name	`xml:"item"`
		PubDate		string 		`xml:"pubDate"`
		Title		string		`xml:"title"`
		Description	string		`xml:"description"`
		Link		string		`xml:"link"`
		GUID		string		`xml:"guid"`
		GeorssPoint	string		`xml:"georss:point"`
	}

	// image 구조체는 RSS 문서 내의 image 태그에
	// 정의된 필드들에 대응하는 필드들을 선언한다.
	image struct {
		XMLName 	xml.Name	`xml:"image"`
		URL			string 		`xml:"url"`
		Title		string		`xml:"title"`
		Link		string		`xml:"link"`
	}

	// channel 구조체는 RSS문서 내의 channel 태그에
	// 정의된 필드들에 대응하는 필드들을 선언하다.
	channel struct {
		XMLName 		xml.Name	`xml:"channel"`
		Title			string		`xml:"title"`
		Description		string		`xml:"description"`
		Link			string		`xml:"link"`
		PubDate			string 		`xml:"pubDate"`
		LastBuildDate	string		`xml:"lastBuildDate"`
		TTL				string		`xml:"ttl"`
		Language		string		`xml:"language"`
		ManagingEditor	string		`xml:"managingEditor"`
		WebMaster		string		`xml:"webMaster"`
		Image			image		`xml:"image"`
		Item			[]item		`xml:"item"`
	}

	// rssDoucment 구조체는 RSS 문서에 정의된 필드들에 대응하는 필드을 정의한다.
	rssDocument struct {
		XMLName 	xml.Name	`xml:"item"`
		Channel 	channel		`xml:"channel"`
	}
)

// it should be declare with rassMacher that implement Matcher
// Matcher 인터페이스를 구현하는 rassMatcher 타입을 선언한다.
// 이 코드는 앞서 defaultMatcher 타입을 선언했을 때의 코드와 거의 동일하게 보인다.
// 이 경우 역시 Matcher 인터페이스를 구현할 뿐 관리해야 할 상태가 없기 때문에
// 빈 구조체를 사용해도 무방하다.
// 다음으로 해야 할 일은 init 함수를 통해 검색기를 등록하는 작업이다.
type rssMatcher struct {}

// init 함수를 통해 프로그램에 검색기를 등록한다.
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

/**
Search-function search that keyword from particular document
 */
func (m rssMatcher) Search(feed* search.Feed, searchTerm string)([]*search.Result, error) {
	var results[]* search.Result

	log.Printf("Feed Type[%s] Site[%s] Adress[%s] search\n", feed.Type, feed.Name, feed.URI)

	// Query the data to be searched
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		// keyword search from title
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		// If It find searched word That the result save
		if matched {
			results = append(results, &search.Result{
				Field:	"Title",
				Content: channelItem.Title,
			})
		}

		// keyword search from description
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		// If It find searched word That the result save
		if matched {
			results = append(results, &search.Result{
				Field:	"Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}

/**
	The result decoding after that RSS feed request of  HTTP GET
 */
func (m rssMatcher) retrieve(feed* search.Feed)(*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("It is not defind that search RSS Feed")
	}

	// RSS document search from Web
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// Response stream will close when function is return
	defer resp.Body.Close()

	// Status code inspect that it is 200
	// Response check if It is right or wrong
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error: %d\n", resp.StatusCode)
	}

	// Decode that Structure of RSS feed document
	// This function is not treat error because That caller function decide if fail or success
	var doucment rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&doucment)
	return &doucment, err
}


