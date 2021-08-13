package rolling_number

import (
	"sync"

	goadder "github.com/linxGnu/go-adder"
)

type bucketArray struct {
	array [timeWindow]*bucket
	mutex sync.RWMutex
}

// bucket 内部维护一个基于 LongAdder 实现的原子计数器，
// 相较于普通的原子计数器，LongAdder 采用了以空间换时间的方法，
// 使得其吞吐量大大增加
type bucket struct {
	value goadder.LongAdder
	time int64 // const
}

func newBucket(time int64) *bucket {
	adder := goadder.NewLongAdder(goadder.JDKAdderType)
	return &bucket{
		value: adder,
		time:  time,
	}
}

func newBucketArray() *bucketArray {
	return &bucketArray{}
}

func(b *bucket) add(n int64) {
	b.value.Add(n)
}

func(b *bucket) sum() int64 {
	return b.value.Sum()
}

func(b *bucket) isValid(now int64) bool {
	if b.time >= now - timeWindow {
		return true
	}
	return false
}

func(q *bucketArray) get(pos int64) (*bucket, bool) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	if b := q.array[pos % timeWindow]; b != nil {
		return b, true
	}
	return nil, false
}

func(q *bucketArray) put(pos int64, newBucket *bucket) {
	if newBucket == nil {
		return
	}

	idx := pos % timeWindow
	// 双重校验
	if oldBucket := q.array[idx]; oldBucket == nil || oldBucket.time < newBucket.time {
		q.mutex.Lock()
		if oldBucket := q.array[idx]; oldBucket == nil || oldBucket.time < newBucket.time {
			q.array[idx] = newBucket
		}
		q.mutex.Unlock()
	}
}

func(q *bucketArray) forEach(fn func(b *bucket)) {
	for _, b := range q.array {
		fn(b)
	}
}