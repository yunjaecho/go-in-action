package search

import (
	"os"
	"encoding/json"
)

const dataFile = "/home/comp1/project/golang/go-in-action/chapter2/sample/data/data.json"



// 피드를 처리할 정보를 표현하는 구조체
type Feed struct {
	Name string `json:"site"`
	URI string `json:"link"`
	Type string `json:"type"`
}

// RetriveFeeds 함수는 피드 데이터 파일을 읽어 구조체로 변환한다.
func RetriveFeeds()([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	// defer 함수를 이용해서 이 함수가 리턴될때
	// 앞서 열어둔 파일이 닫히도록 한다.
	defer file.Close()

	// 슬라이스로 변환한다.
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// 호출 함수가 오류를 처리할 수 없으므로 오류 처리는 하지 않는다.
	return feeds, err



}