package go_learning

import "testing"

func TestM3Message(t *testing.T) {
	//开发生产者/消费者程序，按接口开发实现，消息只需要支持字符串即可
	//生产者每1秒钟生产一条消息，消息格式为：product message time:当前时间
	//消费者每3秒钟消费3条消息，并把消费的消息打印出来，消息格式为：consume message time:当前时间 message1或者message2或者message3:生产消息
	//要求：满足以下测试用例，消费者按消息格式每3秒钟输出3条消息
	//consume message time:2019-10-12 10:56:08.3479874 +0800 CST m=+30.027009601 message1:product message time:2019-10-12 10:56:06.3469824 +0800 CST m=+28.026002201
	//consume message time:2019-10-12 10:56:08.3479874 +0800 CST m=+30.027009601 message2:product message time:2019-10-12 10:56:07.4419769 +0800 CST m=+29.120998001
	//consume message time:2019-10-12 10:56:08.3479874 +0800 CST m=+30.027009601 message3:product message time:2019-10-12 10:56:08.3479874 +0800 CST m=+30.027009601
	messagechan := make(chan string, 10)
	producer := NewStringProducer(messagechan)
	go producer.Product()
	consumer := NewStringConsumer(messagechan)
	consumer.Consume()

}
