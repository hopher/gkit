package log

import (
	"fmt"

	"github.com/aliyun/aliyun-log-go-sdk/producer"
)

type Callback struct {
}

func (callback *Callback) Success(result *producer.Result) {
	attemptList := result.GetReservedAttempts() // 遍历获得所有的发送记录
	for _, attempt := range attemptList {
		fmt.Println(attempt)
	}
}

func (callback *Callback) Fail(result *producer.Result) {
	fmt.Println(result.IsSuccessful())        // 获得发送日志是否成功
	fmt.Println(result.GetErrorCode())        // 获得最后一次发送失败错误码
	fmt.Println(result.GetErrorMessage())     // 获得最后一次发送失败信息
	fmt.Println(result.GetReservedAttempts()) // 获得producerBatch 每次尝试被发送的信息
	fmt.Println(result.GetRequestId())        // 获得最后一次发送失败请求Id
	fmt.Println(result.GetTimeStampMs())      // 获得最后一次发送失败请求时间
}
