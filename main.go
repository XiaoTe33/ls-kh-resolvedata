package main

import "ls-kh-resolvedata/service"

func main() {
	service.Serve()
	forever := make(chan struct{})
	<-forever
}
