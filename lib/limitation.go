package lib

type Limitation struct {
	PeriodVal  int    // 시간
	PeriodUnit string // MS, SEC, MIN, HOUR, DAY
	Count      int
}
