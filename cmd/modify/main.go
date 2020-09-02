package main

import (
	"flag"
	"fmt"
)



func removeBucket(bucket string) {

}

func addSubscriber(bucket, subscriber string){

}

func removeSubscriber(bucket, subscriber string){

}

func main() {
	// -D, -I
	delBucketFlag := flag.String("D","","delete bucket")
	addBucketFlag := flag.String("A","","add bucket")
	addSubscriberFlag := flag.String("a","","add subscriber")
	delSubscriberFlag := flag.String("d","","delete subscriber")

	flag.Parse()
	if delBucketFlag != nil {
		removeBucket(*delBucketFlag)
	}
	if addBucketFlag != nil {
		//addBucket(*addBucketFlag)
	}
	if addSubscriberFlag != nil {
		tail := flag.Args()
		addSubscriber(*addSubscriberFlag, tail[0])
	}
	if delSubscriberFlag != nil {
		tail := flag.Args()
		removeSubscriber(*delSubscriberFlag, tail[0])
	}

	fmt.Println("hello")
}



