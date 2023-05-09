package define

type PlayerRankInfo struct {
	Numid    int
	Rank     int
	Score    int
	NickName string
	IsMyInfo bool
}

type PlayerRankInfoSql struct {
	Numid    int
	Rank     int
	Score    float64
	NickName string
	IsMyInfo bool
}
