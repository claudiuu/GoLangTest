package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"
)

// Request is used for making requests to services behind a load balancer.
type Request struct {
	Payload interface{}
	RspChan chan Response
}

// Response is the value returned by services behind a load balancer.
type Response interface{}

// LoadBalancer is used for balancing load between multiple instances of a service.
type LoadBalancer interface {
	Request(payload interface{}) chan Response
	RegisterInstance(chan Request)
}

// MyLoadBalancer is the load balancer you should modify!
type MyLoadBalancer struct {
	instances map[int]chan Request
	total     int
	last      int
}

// Request is currently a dummy implementation. Please implement it!
func (lb *MyLoadBalancer) Request(payload interface{}) chan Response {

	lb.determineNextInstance()
	fmt.Println("Instance", lb.last, "is handling the request")
	fmt.Println(lb.instances)

	rsp := make(chan Response)
	lb.instances[lb.last] <- Request{Payload: payload, RspChan: rsp}
	return rsp
}

func (lb *MyLoadBalancer) determineNextInstance() {
	if len(lb.instances) > 0 {
		keys := reflect.ValueOf(lb.instances).MapKeys()
		lb.last = rand.Intn(len(keys))
		fmt.Println(keys)
		// for {
		// 	lb.last = rand.Intn(len(keys))
		// 	fmt.Println("trying instance ", lb.last)
		// 	if _, ok := <-lb.instances[lb.last]; ok {
		// 		return
		// 	}
		// }

	}
	fmt.Println(lb.instances)
	fmt.Println("Seting last as ", lb.last)
}

// RegisterInstance is currently a dummy implementation. Please implement it!
func (lb *MyLoadBalancer) RegisterInstance(ch chan Request) {
	if len(lb.instances) == 0 {
		lb.instances = make(map[int]chan Request)
	}

	// register a new instance at the end
	lb.instances[lb.total] = ch
	fmt.Println("Registered instance", lb.total)
	// go listenForUnregister(ch, lb.total, lb)
	lb.total++
	return
}

// func listenForUnregister(ch chan Request, pos int, lb *MyLoadBalancer) {
// 	select {
// 	case req := <-ch:
// 		if req.Payload != nil && req.Payload.(bool) {
// 			fmt.Println("unregister instance", pos)
// 			delete(lb.instances, pos)
// 		}
// 	}
// }

/******************************************************************************
 *  STANDARD TIME SERVICE IMPLEMENTATION -- MODIFY IF YOU LIKE                *
 ******************************************************************************/

// TimeService is a single instance of a time service.
type TimeService struct {
	Dead            chan struct{}
	ReqChan         chan Request
	Unregister      chan struct{}
	AvgResponseTime float64
}

// Run will make the TimeService start listening to the two channels Dead and ReqChan.
func (ts *TimeService) Run() {
	for {
		select {
		case <-ts.Dead:
			ts.ReqChan <- Request{true, nil}
			// close(ts.ReqChan)
			return
		case req := <-ts.ReqChan:
			processingTime := time.Duration(ts.AvgResponseTime+1.0-rand.Float64()) * time.Second
			time.Sleep(processingTime)
			req.RspChan <- time.Now()
		}
	}
}

/******************************************************************************
 *  CLI -- YOU SHOULD NOT NEED TO MODIFY ANYTHING BELOW                       *
 ******************************************************************************/

// main runs an interactive console for spawning, killing and asking for the
// time.
func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	bio := bufio.NewReader(os.Stdin)
	var lb LoadBalancer = &MyLoadBalancer{}

	manager := &TimeServiceManager{}

	for {
		fmt.Printf("> ")
		cmd, err := bio.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading command: ", err)
			continue
		}
		switch strings.TrimSpace(cmd) {
		case "kill":
			manager.Kill()
		case "spawn":
			ts := manager.Spawn()
			lb.RegisterInstance(ts.ReqChan)
			go ts.Run()
		case "time":
			select {
			case rsp := <-lb.Request(nil):
				fmt.Println(rsp)
			case <-time.After(5 * time.Second):
				fmt.Println("Timeout")
			}
		default:
			fmt.Printf("Unknown command: %s Available commands: time, spawn, kill\n", cmd)
		}
	}
}

// TimeServiceManager is responsible for spawning and killing.
type TimeServiceManager struct {
	Instances []TimeService
}

// Kill makes a random TimeService instance unresponsive.
func (m *TimeServiceManager) Kill() {
	if len(m.Instances) > 0 {
		fmt.Printf("There are %d instances. ", len(m.Instances))
		n := rand.Intn(len(m.Instances))
		fmt.Printf("Killing Instance %d\n", n)
		close(m.Instances[n].Dead)
		m.Instances = append(m.Instances[:n], m.Instances[n+1:]...)
	} else {
		fmt.Println("No instance to kill")
	}
}

// Spawn creates a new TimeService instance.
func (m *TimeServiceManager) Spawn() TimeService {
	ts := TimeService{
		Dead:            make(chan struct{}, 0),
		ReqChan:         make(chan Request, 10),
		AvgResponseTime: rand.Float64() * 3,
	}
	m.Instances = append(m.Instances, ts)
	return ts
}
