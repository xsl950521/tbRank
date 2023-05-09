package timetask

import (
	"time"
)

// DailyCron 每日指定时间 执行任意个无参任务
// 若今天已经超过了执行时间则等到第二天的指定时间再执行任务
func DailyCron(dateTime time.Time, tasks ...func()) {
	for {
		now := time.Now()
		next := time.Date(now.Year(), now.Month(), now.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), dateTime.Nanosecond(), dateTime.Location())
		// 检查是否超过当日的时间
		if next.Sub(now) < 0 {
			next = now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), dateTime.Nanosecond(), dateTime.Location())
		}
		// 阻塞到执行时间
		t := time.NewTimer(next.Sub(now))
		<-t.C
		// 执行的任务内容
		for _, task := range tasks {
			go task()
		}
	}
}
