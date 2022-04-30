package bithumb

type MsgDesc struct {
	Status string `json:"status"`
	Msg    string `json:"resmsg"`
}

var errors = map[int32][]MsgDesc{
	5100: {
		MsgDesc{"Bad Request(SSL)", "https 호출 URL이 아님"},
		MsgDesc{"Bad Request(Bad Method)", "POST 형식으로 호출하지 않음"},
		MsgDesc{"Bad Request(Auth Data)", "잘못된 요청 내용"},
		MsgDesc{"Bad Request(Request Time:reqTime{호출시간}/nowTime{서버시간})", "API 서버 시간과 API 호출 시간이 20초 이상 차이남(시간 기준: KST)"},
	},
	5200: {
		MsgDesc{"Not Member", "회원가입이 되어있지 않음"},
	},
	5300: {
		MsgDesc{"Invalid Apikey", "올바르지 않은 API Key, Secret Key로 호출"},
		MsgDesc{"Method Not Allowed.(Access IP)", "접속 허용 IP가 아님"},
	},
	5302: {
		MsgDesc{"Method Not Allowed.(BTC Adress)", "BTC 출금 허용 주소가 아님"},
		MsgDesc{"Method Not Allowed.(Access)", "API 활성화 시 설정한 항목 외 호출했을 경우"},
	},
	5400: {
		MsgDesc{"Database Fail", "데이터베이스 에러일 경우"},
	},
	5500: {
		MsgDesc{"Invalid Parameter", "잘못된 인자 값으로 호출"},
	},
	5600: {
		MsgDesc{"too many connections", "과도한 접속시도로 인한 접속제한 상태특정 IP 차단 (일반적인 유저에게는 나타날 수 없음)"},
		MsgDesc{"Please try again", "과도한 거래시도로 인한 접속제한"},
		MsgDesc{"Not Allow IP", "PRIME이 아닌 유저가 PRIME/public API를 요청"},
	},
	5900: {
		MsgDesc{"Unknown Error", "알 수 없는 에러"},
	},
}
