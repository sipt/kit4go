package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
)

func main() {
	var v *int
	v = nil
	var c interface{}
	c = v
	fmt.Println(c != nil)

	arr := [2]int{}
	fmt.Printf("%p\n", &arr)
	tmp := arr[:0]
	fmt.Printf("%p\n", tmp)
	tmp = append(tmp, 1)
	fmt.Printf("%p\n", tmp)
	tmp = append(tmp, 1, 1)
	fmt.Printf("%p\n", tmp)
	tmp = append(tmp, 1, 1, 1, 1, 1)
	fmt.Printf("%p\n", tmp)

	_ = new(chan bool)

	m := make(map[string]*A)
	m["a"] = &A{[]string{"haha"}}
	m["a"].List = append(m["a"].List, "xixi")
	fmt.Println(m["a"].List)

	x := new([]int)
	fmt.Printf("%p\n", *x)
	*x = append(*x, 1, 2, 3, 4)
	fmt.Printf("%p\n", *x)
	*x = append(*x, 1, 2, 3, 4)
	fmt.Printf("%p\n", *x)
	fmt.Println(cap(*x), len(*x))

	// cap
	{
		arr := make([]int, 10)
		arr = append(arr, 10)
		fmt.Printf("ptr:%p, cap:%d\n", arr, cap(arr))
		arr = arr[:10]
		fmt.Printf("ptr:%p, cap:%d\n", arr, cap(arr))
	}

	//nats test
	{
		server1 := "nats://10.10.11.23:4222"
		server2 := "nats://10.10.11.23:4223"
		key := "hello"
		conn1, err := nats.Connect(server1)
		if err != nil {
			panic(err)
		}
		conn1.Subscribe(key, func(msg *nats.Msg) {
			fmt.Printf("conn1: %s\n", string(msg.Data))
		})

		conn2, err := nats.Connect(server2)
		if err != nil {
			panic(err)
		}
		conn2.Subscribe(key, func(msg *nats.Msg) {
			fmt.Printf("conn2: %s\n", string(msg.Data))
		})

		conn3, err := nats.Connect(server2)
		if err != nil {
			panic(err)
		}
		conn3.Subscribe(">", func(msg *nats.Msg) {
			fmt.Printf("conn3: %s\n", string(msg.Data))
		})

		conn4, err := nats.Connect(server2)
		if err != nil {
			panic(err)
		}
		conn4.Subscribe(key, func(msg *nats.Msg) {
			fmt.Printf("conn3: %s\n", string(msg.Data))
		})

		conn1.Publish(key, []byte("hello world![conn1]"))
		conn2.Publish(key, []byte("hello world![conn2]"))
		conn1.Publish("a.b.c.d", []byte("hello world![conn3]"))
		end := make(chan bool)
		<-end
	}
}

type A struct {
	List []string
}

func test() *int {
	x := new(int)
	*x = 9
	return x
}

//func test1() {
//	arr := []int{1, 2, 3, 4, 5}
//	fmt.Println(len(arr), cap(arr))
//	arr = append(arr, 1, 2)
//	fmt.Println(len(arr), cap(arr))
//	arr2 := arr[:len(arr)/2]
//	arr2[1] = 999
//	fmt.Println("arr:", arr, "\n", "arr2:", arr2)
//	arr2 = append(arr2, 1000)
//	fmt.Println("arr:", arr, "\n", "arr2:", arr2)
//}
