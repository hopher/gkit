package log

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/aliyun/aliyun-log-go-sdk/producer"
)

func Example_basic() {
	fmt.Println("Hello")
	// Output: Hello
}

func ExampleAliLogger_Log() {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = os.Getenv("Endpoint")
	producerConfig.AccessKeyID = os.Getenv("AccessKeyID")
	producerConfig.AccessKeySecret = os.Getenv("AccessKeySecret")
	producerInstance := producer.InitProducer(producerConfig)
	producerInstance.Start()           // 启动producer实例
	defer producerInstance.SafeClose() // 安全关闭

	logger := NewAliLogger("projectName", "logstorName", "127.0.0.1", "topic",
		producerInstance,
		WithAliLoggerCallBack(&Callback{}), // 可选
	)
	logger.Log("key", "value", "key2", "value2", "key3", 300)
}
