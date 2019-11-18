package main

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// M个receiver，N个sender，它们当中任意一个通过通知一个moderator（仲裁者）关闭额外的
// signal channel来说“让我们结束游戏吧”.这是最复杂的场景了。我们不能让任意的receivers
// 和senders关闭data channel，也不能让任何一个receivers通过关闭一个额外的
// signal channel来通知所有的senders和receivers退出游戏。这么做的话会打破
// channel closing principle。但是，我们可以引入一个moderator来关闭一个额外的
// signal channel。这个例子的一个技巧是怎么通知moderator去关闭额外的signal channel

func case4() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the moderator goroutine shown below.
	// Its reveivers are all senders and receivers of dataCh.
	toStop := make(chan string, 1)
	// the channel toStop is used to notify the moderator
	// to close the additional signal channel (stopCh).
	// Its senders are any senders and receivers of dataCh.
	// Its reveiver is the moderator goroutine shown below.

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		close(stopCh) // part of the trick used to notify the moderator
		// to close the additional signal channel.
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(MaxRandomNumber)
				if 0 == value {
					// here, a trick is used to notify the moderator
					// to close the additional signal channel.
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// the first select here is to try to exit the goroutine
				// as early as possible.This select blocks with one
				// receive operation case and one default branches will
				// be specially optimized as a try-receive operation by
				// the standard Go compiler.
				select {
				case <-stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first branch in the
				// second select may be still not selected for some
				// loops (and for ever in theory) if the send to
				// dataCh is also non-blocking.
				// This is why the first select block above is needed.
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// same as senders, the first select here is to
				// try to exit the goroutine as early as possible.
				// This select blocks with one send operation case and
				// one default branches will be specially optimized as
				// a try-send operation by the standard Go compiler.
				select {
				case <-stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first branch in the
				// second select may be still not selected for some
				// loops (and for ever in theory) if the receive from
				// dataCh is also non-blocking.
				// This is why the first select block is needed.
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == MaxRandomNumber-1 {
						// the same trick is used to notify the moderator
						// to close the additional signal channel.
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}
