package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		//		"http://bixer.world", // Guaranteed failure case
	}

	c := make(chan string) // channels are typed. Only data of defined type can be transmitted

	for _, link := range links {
		go checkLink(link, c) /* "go" keyword in front of a function call makes it a routine
		routines can execute concurrently (ie CPU only works on it until blocking call or finished)
		But the main routine keeps executing to the end and may finish before child routines.
		This is where channels come in. Channels are how Go routines communicate with one another.
		*/
	}

	/*
		To send value into a channel: channelName <- 5
		To wait for a value to be sent from a channel: myNumber <- channelName
		Wait for value to be sent to channel then log immediately: fmt.Println(<- channelName)
	*/
	/*
		for i := 0; i < len(links); i++ { // Iterates times the number of messages expected over the channel
			fmt.Println(<-c) // This is blocking code
		}
	*/
	/*
		for { // Infinite loop
			go checkLink(<-c, c) // Any time a value goes into the channel, a new checkLink routine is spawned
			// checkLink takes a string and a channel. Go sees that the first value is of type string
		}
	*/
	for l := range c { // equivalent to infinite loop above
		//		go checkLink(l, c)          // replace <-c with l which is the range of the channel
		go func(link string) { // Function Literal: same as Lambda function in Python/Ruby or Anonymous Function
			time.Sleep(5 * time.Second) /* sleep for 5 seconds. This shouldn't be in the main routine.
			But it also might not be appropriate to put in the subroutine function, because we want to know
			the current uptime-- not X seconds ago.
			*/
			checkLink(link, c)
		}(l) // Needs trailing parens to actually call the function. Pass in "l" as a variable.
		// String is copied to "link". The value is updated each time rather than only once.
	}
}

func checkLink(link string, c chan string) { // must declare chan "c" is type string here as well
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
