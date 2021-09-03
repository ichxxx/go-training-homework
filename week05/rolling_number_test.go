package rolling_number

import (
	"sync"
	"testing"
	"time"

	"go.uber.org/atomic"
)

func TestRollingNumberConcurrent(t *testing.T) {
	number := NewRollingNumber()
	concurrency := 10000
	var cntList [timeWindow]int64
	var totalCnt int64

	for i := 0; i < timeWindow * 2; i++ {
		now := time.Now()
		nowUnix := now.Unix()
		cnt := atomic.NewInt64(0)

		for j := 0; j < concurrency; j++ {
			go func() {
				// add := int64(rand.Intn(10))
				var add int64 = 1
				number.Add(add)
				cnt.Add(add)
			}()
		}

		time.Sleep(time.Second)

		_cnt := cnt.Load()
		totalCnt += _cnt - cntList[nowUnix % timeWindow]
		sum := number.Sum(now)
		t.Log(nowUnix, "-", _cnt, "-", totalCnt, "-", sum)

		cntList[nowUnix % timeWindow] = _cnt
	}
}

func BenchmarkRollingNumberAdd(b *testing.B) {
	number := NewRollingNumber()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		number.Add(1)
	}
}

func BenchmarkRollingNumberAddConcurrent(b *testing.B) {
	number := NewRollingNumber()
	concurrency := 100000
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(concurrency)
		for j := 0; j < concurrency; j++ {
			go func() {
				number.Add(1)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
