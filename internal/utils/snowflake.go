package utils

import (
	"log"
	"strconv"
	"sync"
	"time"
)

const (
	epoch             = int64(1577808000000)                           // 设置起始时间(时间戳/毫秒)：2020-01-01 00:00:00，有效期69年
	timestampBits     = uint(41)                                       // 时间戳占用位数
	datacenterIdBits  = uint(3)                                        // 数据中心id所占位数
	workerIdBits      = uint(7)                                        // 机器id所占位数
	sequenceBits      = uint(12)                                       // 序列所占的位数
	timestampMax      = int64(-1 ^ (-1 << timestampBits))              // 时间戳最大值
	datacenterIdMax   = int64(-1 ^ (-1 << datacenterIdBits))           // 支持的最大数据中心id数量
	workerIdMax       = int64(-1 ^ (-1 << workerIdBits))               // 支持的最大机器id数量
	sequenceMask      = int64(-1 ^ (-1 << sequenceBits))               // 支持的最大序列id数量
	workerIdShift     = sequenceBits                                   // 机器id左移位数
	datacenterIdShift = sequenceBits + workerIdBits                    // 数据中心id左移位数
	timestampShift    = sequenceBits + workerIdBits + datacenterIdBits // 时间戳左移位数
)

type Snowflake struct {
	lock         sync.Mutex // 锁
	timestamp    int64      // 时间戳 ，毫秒
	workerId     int64      // 工作节点
	datacenterId int64      // 数据中心机房id
	sequence     int64      // 序列号
}

func NewIdGenerator(datacenterId int64, workerId int64) *Snowflake {
	if datacenterId > datacenterIdMax {
		panic("datacenterId overflows.")
	}
	if workerId > workerIdMax {
		panic("workerId overflows.")
	}
	return &Snowflake{
		sync.Mutex{},
		time.Now().UnixMilli(),
		workerId,
		datacenterId,
		0,
	}
}

func (this *Snowflake) NextId() string {
	this.lock.Lock()
	now := time.Now().UnixMilli()
	if this.timestamp == now {
		this.sequence = (this.sequence + 1) & sequenceMask
		// 如果序列号重置，需要等待下个时间戳
		if this.sequence == 0 {
			for now <= this.timestamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		// 不同时间戳下重置序列号
		this.sequence = 0
	}
	t := now - epoch
	if t > timestampMax {
		this.lock.Unlock()
		log.Fatalln("Timestamp overflow!")
		return ""
	}
	this.timestamp = now
	id := (t)<<timestampShift | (this.datacenterId << datacenterIdShift) | (this.workerId << workerIdShift) | (this.sequence)
	this.lock.Unlock()
	return strconv.FormatInt(id, 16)
}