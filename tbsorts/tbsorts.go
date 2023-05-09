package tbsorts

import (
	"fmt"
	"sort"
	"time"
)

var scores []int64

// func main() {
// 	sorts()
// 	//bit()
// }

func bit() {
	fmt.Println(len(fmt.Sprintf("%d", 1<<9)))
	fmt.Println(len(fmt.Sprintf("%d", 1<<41)))
	fmt.Println(len(fmt.Sprintf("%d", 1<<23)))
	fmt.Println(len("1000000"))
}

func Sorts(score int64) {
	now := time.Now().UnixMilli()
	for i := 1; i < 10; i++ {
		scores = append(scores, ToScore41(int64(i+100000), now))
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})
	for _, score := range scores {
		fmt.Println(load(score))
		//fmt.Println(score)
	}
}

func ToScore41(point int64, periodEndTimestamp int64) (score int64) {
	score = (score | point) << 41
	score = score | periodEndTimestamp
	return
}
func ToScore9(point int64, periodEndTimestamp int64) (score int64) {
	score = (score | point) << 9
	score = score | periodEndTimestamp
	return
}

func load(score int64) int64 {
	return score >> 41
}
