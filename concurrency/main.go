package main

import (
	"fmt"
	"sync"
	"time"
)

func infiniteCount(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Second * 1)
	}
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func countWithChannel(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

func main() {
	lesson1()
	// ...
	// lesson9()
}

// this eg will count dog forever and never get to the cat
func lesson1() {
	infiniteCount("dog")
	infiniteCount("cat")
}

/*
this eg using goroutine it will run count("dog") in the background
then continue to the next line count("cat")
*/
func lesson2() {
	go infiniteCount("dog")
	infiniteCount("cat")
}

/*
what if we run both count func as goroutine? should be the same as lesson2?
-> we don't get anything, what happened? -> when the main func finishes the program exits
*/
func lesson3() {
	go infiniteCount("dog")
	go infiniteCount("cat")
}

/*
what we can do is to use a wait group -> it just the counter
wg.Add(1) to increment the counter to say we have 1 goroutine to wait for and
wg.Done() to decrement the counter when the goroutine finishes
wg.Wait() will block until the counter is zero
*/
func lesson4() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("dog")
		wg.Done()
	}()

	wg.Wait()
}

/*
channel -> way for goroutine to communicate with each other
*/
func lesson5() {
	c := make(chan string)
	go countWithChannel("dog", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

/*
we might expect this to work and just output "hello world" but we actually get deadlock!
this bacause the send will block until something is ready to receive
*/
func lesson6() {
	c := make(chan string)
	c <- "hello world"
	msg := <-c
	fmt.Println(msg)
}

/*
we can make a buffered channel by giving a capacity when we make the channel
you can fill up a buffered channel without a corresponding receiver and it won't block
until the channel is full
*/
func lesson7() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}

/*
select -> allow us to receive from whichever channel is ready
*/
func lesson8() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			c1 <- "Every 500ms"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- "Every two seconds"
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

/*
worker pools
*/
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
func lesson9() {
	// In order to use our pool of workers we need to send them work and collect their results. We make 2 channels for this.
	const numJobs = 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This starts up 3 workers, initially blocked because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 10 jobs and then close that channel to indicate that’s all the work we have.
	for j := 1; j <= numJobs; j++ {
		fmt.Println("send job: ", j)
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work. This also ensures that the worker goroutines have finished. An alternative way to wait for multiple goroutines is to use a WaitGroup.
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
