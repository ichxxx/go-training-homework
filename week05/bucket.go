package rolling_number

import (
	goadder "github.com/linxGnu/go-adder"
)

type bucketArray [timeWindow]*bucket

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

func(b *bucket) add(n int64) {
	b.value.Add(n)
}

func(b *bucket) sum() int64 {
	return b.value.Sum()
}

func(q *bucketArray) get(pos int64) (*bucket, bool) {
	if b := q[pos % timeWindow]; b != nil {
		return b, true
	}
	return nil, false
}

func(q *bucketArray) put(pos int64, bucket *bucket) {
	if bucket == nil {
		return
	}

	// 指针类型的替换是原子操作，不会发生并发冲突
	// 但有可能会发生数据覆盖的问题
	// 不过实际测试下来，该情况很少发生
	// 即使发生了，数据丢失的数量也在可接受范围内
	// 为了提高性能，这里选择不加锁
	q[pos % timeWindow] = bucket
}