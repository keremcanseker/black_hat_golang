package main

import (
	"fmt"
	"net"
	"sort"
)

/*
The channel will be used
to receive work, and the WaitGroup will be used to track when a single work
item has been completed
*/
func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}
//To avoid inconsistencies, you’ll use a pool of goroutines to manage the
//concurrent work being performed. Using a for loop, you’ll create a certain number of worker goroutines as a resource pool. Then, in your main()
//“thread,” you’ll use a channel to provide work. 
func main(){
	
	ports:= make(chan int, 100)
	results:= make(chan int)
	var openports []int

	for i:=0; i<cap(ports); i++{
		go worker(ports, results)
	}

	go func(){
		for i:=1; i<=1024; i++{
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
		openports = append(openports, port)
	}
	}
	close(ports)
 	close(results)

	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}

}