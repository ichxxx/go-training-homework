package rolling_number

import (
	"time"
)

const timeWindow = 10

// RollingNumber 内部维护着一个计数桶列表，
// 每个计数桶的时间长度为1秒，共有10个桶，滑动统计10秒内的计数
type RollingNumber struct {
	buckets *bucketArray
}

func NewRollingNumber() *RollingNumber {
	return &RollingNumber{buckets: &bucketArray{}}
}

func(rn *RollingNumber) getCurrentBucket() *bucket {
	now := time.Now().Unix()
	b, ok := rn.buckets.get(now)

	// 当前时间和桶的时间不一致时，
	// 说明该桶已经过期（即超过10秒）
	// 用新桶替换掉
	if !ok || b.time != now {
		b = newBucket(now)
		rn.buckets.put(now, b)
	}
	return b
}

func(rn *RollingNumber) getBucket(t int64) (*bucket, bool) {
	now := time.Now().Unix()

	if t >= now - timeWindow {
		b, ok := rn.buckets.get(t)
		if !ok || b.time != now {
			b = newBucket(now)
			rn.buckets.put(now, b)
		}
		return b, true
	}
	return nil, false
}

// Add 向当前时间桶增加计数 i
func (rn *RollingNumber) Add(i int64) {
	if i == 0 {
		return
	}

	rn.getCurrentBucket().add(i)
}

// AddWithTime 向特定时间桶增加计数 i
func (rn *RollingNumber) AddWithTime(time, i int64) {
	if i == 0 {
		return
	}

	b, ok := rn.getBucket(time)
	if ok {
		b.add(i)
	}
}

// Sum 对当前所有桶的计数求和
func(rn *RollingNumber) Sum(now time.Time) (sum int64) {
	for _, b := range rn.buckets {
		if b != nil && b.time >= now.Unix() - timeWindow {
			sum += b.sum()
		}
	}

	return
}

// Max 返回当前所有桶中的最大计数
func(rn *RollingNumber) Max(now time.Time) (max int64) {
	for _, b := range rn.buckets {
		if b != nil && b.time >= now.Unix() - timeWindow {
			if val := b.sum(); val > max {
				max = val
			}
		}
	}

	return
}

// Avg 对当前所有桶求平均计数
func(rn *RollingNumber) Avg(now time.Time) int64 {
	return rn.Sum(now) / timeWindow
}

// Reset 重置所有桶的计数
func(rn *RollingNumber) Reset() {
	*rn = *NewRollingNumber()
}

// UpdateMax 更新当前时间桶中的最大计数.
func(rn *RollingNumber) UpdateMax(max int64) {
	b := rn.getCurrentBucket()
	if max > b.sum() {
		b.value.Store(max)
	}
}