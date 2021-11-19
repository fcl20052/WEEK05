package toytrix

import (
	"errors"
	"time"
)

var LimitQueue map[string][]int64

//滑动窗口限流
func LimitFreqSingle(queName string, count int, timeSpan int64) bool {
	currTime := time.Now().Unix()
	if LimitQueue == nil {
		LimitQueue = map[string][]int64{}
	}
	if _, ok := LimitQueue[queName]; !ok {
		LimitQueue[queName] = make([]int64, 0)
	}
	//窗口未满
	if len(LimitQueue[queName]) < count {
		LimitQueue[queName] = append(LimitQueue[queName], currTime)
		return true
	}
	//窗口满了，开始判断是否滑动
	earlyTime := LimitQueue[queName][0]
	//小于限制时间，请求拒绝
	if currTime-earlyTime <= timeSpan {
		return false
	} else {
		//大于限制时间，接受请求，开始向后滑动
		LimitQueue[queName] = LimitQueue[queName][1:]
		LimitQueue[queName] = append(LimitQueue[queName], currTime)
	}
	return true

}

func Do(name string, f func() error, fallback func(error) error) error {
	//5秒内限制2个访问
	if LimitFreqSingle(name, 2, 5) {
		f()
		return nil
	} else {
		fallback(errors.New("frequently visitied"))
		return errors.New("frequently visitied")
	}
}
