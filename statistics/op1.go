package statistics

import "fmt"

//msrstnAcctoLastDcsnDnstyResponseItem 은 측정소별 최종확정 농도를 조회했을 때 서버에서
//return 해주는 데이터 집합이다.
type msrstnAcctoLastDcsnDnstyResponseItem struct {
	DataTime    string  `xml:"dataTime"` //오염도 측정 년도-월-일
	So2Average  float64 `xml:"so2Avg"`   //아황산가스 농도 (단위 : ppm)
	CoAverage   float64 `xml:"coAvg"`    //일산화탄소 농도 (단위 : ppm)
	O3Average   float64 `xml:"o3Avg"`    //오존 농도 (단위 : ppm)
	No2Average  float64 `xml:"no2Avg"`   //이산화질소 농도 (단위 : ppm)
	Pm10Average int64   `xml:"pm10Avg"`  //미세먼지 농도 (단위 : ㎍/㎥)
}

//msrstnAcctoLastDcsnDnstyRequest 은 측정소별 최종확정 농도를 조회할 때 사용하는 데이터 집합이다.
type msrstnAcctoLastDcsnDnstyRequest struct {
	StationName     string //측정소명
	SearchCondition string //검색 조건 (년도별 : YEAR, 월별:  MONTH, 일별 : DAILY)
	PageNo          int64  //페이지 번호
	NumOfRows       int64  //한 페이지 결과 수
}

//Op1ResponseItem 은 msrstnAcctoLastDcsnDnstyResponseItem 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op1ResponseItem msrstnAcctoLastDcsnDnstyResponseItem

//Op1Request 은 msrstnAcctoLastDcsnDnstyRequest 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op1Request msrstnAcctoLastDcsnDnstyRequest

//Op1Operator 은 근접측정소 목록 OpenAPI를 실행하기 위한 명령을 만들도록 도와주는 인터페이스이다.
type Op1Operator struct {
	Req Op1Request
}

//Request는 msrstnAcctoLastDcsnDnstyRequest 객체 주소를 리턴한다.
//interface{}타입으로 리턴되므로 다른 모듈에서 데이터 사용시 타입지정이 필요하다.
func (o *Op1Operator) Request() interface{} {
	return &o.Req
}

//RequestString은 근접측정소 목록 조회의 Operation 1번에 대한 Url을 입력된
//데이터를 바탕으로 만들어서 리턴한다.
func (o Op1Operator) RequestString() string {
	return fmt.Sprintf("%s/%s?stationName=%s&searchCondition=%s&pageNo=%d&numOfRows=%d",
		serviceName, operationMap[1],
		o.Req.StationName, o.Req.SearchCondition, o.Req.PageNo, o.Req.NumOfRows)
}

//Item은 Op1ResponseItem 데이터를 생성하여 주소를 리턴한다.
func (o Op1Operator) Item() interface{} {
	return &Op1ResponseItem{}
}
