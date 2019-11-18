# channel-closing

The Channel Closing Principle
在使用Go channel的时候，一个适用的原则是不要从接收端关闭channel，也不要在多个并发发送端中关闭channel。换句话说，如果sender(发送者)只是唯一的sender或者是channel最后一个活跃的sender，那么你应该在sender的goroutine关闭channel，从而通知receiver(s)(接收者们)已经没有值可以读了。维持这条原则将保证永远不会发生向一个已经关闭的channel发送值或者关闭一个已经关闭的channel

[article](https://studygolang.com/articles/9478)
