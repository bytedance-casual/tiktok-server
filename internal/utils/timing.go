package utils

import (
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

//定时任务

// SetAsynchronousTask 定时任务执行，ms为执行时刻的毫秒时间戳
func SetAsynchronousTask(ms int64, task func()) {
	execTime := time.UnixMilli(ms)
	nowTime := time.Now()
	timer := time.NewTimer(execTime.Sub(nowTime))
	go func() {
		<-timer.C
		task()
	}()
}

var privateCron *cron.Cron

// SetTimingTask 设置周期性任务执行
func SetTimingTask(task func(), cronExpression string) { //每天凌晨两点进行任务
	log.Println("[Time Task Setup]", cronExpression)
	if privateCron == nil {
		privateCron = cron.New()
	}
	//"0 2 * * *"
	_, err := privateCron.AddFunc(cronExpression, task)
	if err != nil {
		log.Println(err)
		return
	}
	privateCron.Start()
}

// TodayZeroUnix 获取当天的8点对应的时间戳 TODO 注意处理时区问题
func TodayZeroUnix() int64 {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
	return startTime
}