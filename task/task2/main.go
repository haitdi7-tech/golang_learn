package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//fmt.Println("Hello, world!")
	//GoTask()
	//time.Sleep(time.Second)

	//schedulerTest()

	//interfaceTest()

	//personTest()

	//ChannelNoCacheTest()

	//ChannelCacheTest()

	//LockTest()

	Atomtest()
}

func personTest() {
	e := Employee{EmployeeID: "e222",
		Person: Person{Name: "Alice", Age: 20},
	}
	e.PrintInfo()
}

func interfaceTest() {
	r := Rectangle{name: "r1"}
	c := Circle{name: "c1"}

	pr := &r
	pc := &c
	r.Area()
	r.Perimeter()
	c.Area()
	c.Perimeter()

	pr.Area()
	pr.Perimeter()
	pc.Area()
	pc.Perimeter()
}

// 任务调度
func schedulerTest() {
	s := Scheduler{}

	tasks := []Task{
		{
			TaskName: "t1",
			Func: func() {
				time.Sleep(time.Second)
			},
		},
		{
			TaskName: "t2",
			Func: func() {
				time.Sleep(2 * time.Second)
			},
		},
		{
			TaskName: "t3",
			Func: func() {
				time.Sleep(4 * time.Second)
			},
		},
		{
			TaskName: "t4",
			Func: func() {
				time.Sleep(3 * time.Second)
			},
		},
	}

	s.Tasks = tasks

	// for _, v := range tasks {
	// 	s.AddTask(v)
	// }

	s.Run()

	time.Sleep(5 * time.Second)

	for _, res := range s.Results {

		fmt.Printf(
			"任务：%s | 开始：%s | 结束：%s | 耗时：%v ",
			res.TaskName,
			res.StartTime.Format("15:04:05.000"),
			res.EndTime.Format("15:04:05.000"),
			res.Duration,
		)
	}
}

// 指针加10
func add(num *int) {
	*num += 10
}

// 切片乘2
func mutiply(nums []int) {
	for i := 0; i < len(nums); i++ {
		nums[i] *= 2
	}
}

// 协程调用goroutine
func GoTask() {

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(100 * time.Millisecond)
			if i%2 == 1 {
				println("打印奇数", i)
			} else {
				continue
			}
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(100 * time.Millisecond)
			if i%2 == 0 {
				println("打印偶数", i)
			} else {
				continue
			}
		}
	}()
}

// 任务结构体
type Task struct {
	TaskName string
	Func     func()
}

// 结果结构体
type TaskResult struct {
	TaskName  string
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
}

// 任务调度器
type Scheduler struct {
	Tasks   []Task
	Results []TaskResult
}

// 添加任务
func (scheduler *Scheduler) AddTask(t Task) {
	scheduler.Tasks = append(scheduler.Tasks, t)
}

func (s *Scheduler) Run() {

	for _, t := range s.Tasks {

		go func(t Task) {
			result := TaskResult{
				TaskName:  t.TaskName,
				StartTime: time.Now(),
			}
			t.Func()
			result.EndTime = time.Now()
			result.Duration = result.EndTime.Sub(result.StartTime)
			s.Results = append(s.Results, result)
		}(t)
	}
}

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	name string
}

type Circle struct {
	name string
}

func (r *Rectangle) Area() {
	fmt.Println(r.name, "Area")
}

func (r *Rectangle) Perimeter() {
	fmt.Println(r.name, "Perimeter")
}

func (r *Circle) Area() {
	fmt.Println(r.name, "Area")
}

func (r *Circle) Perimeter() {
	fmt.Println(r.name, "Perimeter")
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID string
	Person
}

func (e *Employee) PrintInfo() {
	fmt.Printf("EmployeeID: %s | Name: %s | Age: %d", e.EmployeeID, e.Name, e.Age)
}

// 无缓冲
func ChannelNoCacheTest() {
	cn := make(chan int)

	go Send(cn)

	go Receive(cn)

	//time.Sleep(3 * time.Second)
	timeout := time.After(5 * time.Second)

	for {
		select {
		case v, ok := <-cn:
			if !ok {
				fmt.Println("通道已经关闭")
				return
			}
			fmt.Printf("主goroutine接收到数据：%d ", v)

		case <-timeout:
			fmt.Println("超时")
			return
		default:
			fmt.Println("等待数据。。。")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// 有缓存的通道
func ChannelCacheTest() {
	cn := make(chan int, 5)

	go Product(cn)

	go Consume(cn)

	//time.Sleep(3 * time.Second)
	timeout := time.After(5 * time.Second)

	for {
		select {
		case v, ok := <-cn:
			if !ok {
				fmt.Println("通道已经关闭")
				return
			}
			fmt.Printf("主goroutine接收到数据：%d ", v)

		case <-timeout:
			fmt.Println("超时")
			return
		default:
			fmt.Println("等待数据。。。")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func Send(cn chan<- int) {
	for i := 1; i <= 10; i++ {
		cn <- i
		fmt.Printf("发送: %d\n", i)
	}
	close(cn)
}

// 生产
func Product(cn chan<- int) {
	for i := 1; i <= 100; i++ {
		cn <- i
		fmt.Printf("生产发送: %d\n", i)
	}
	close(cn)
}

// 消费
func Consume(cn <-chan int) {
	for rcv := range cn {
		fmt.Printf("消费数据：%d\n ", rcv)
		//fmt.Println()
		//time.Sleep(150 * time.Millisecond)
	}
}

// 接收
func Receive(cn <-chan int) {
	for rcv := range cn {
		fmt.Printf("收到数据：%d ", rcv)
		fmt.Println()
		time.Sleep(150 * time.Millisecond)
	}
}

// 同步锁
func LockTest() {
	var sm sync.Mutex
	var sg sync.WaitGroup
	counter := 0

	for i := 0; i < 10; i++ {
		sg.Add(1)
		go func(sg *sync.WaitGroup) {
			for j := 0; j < 1000; j++ {
				sm.Lock()
				counter += 1
				sm.Unlock()
			}
			sg.Done()
		}(&sg)
	}

	sg.Wait()
	fmt.Println("所有协程操作已完成")
	//time.Sleep(10 * time.Second)

	fmt.Printf("计数器的结果是： %d \n", counter)
}

// 原子操作
func Atomtest() {

	var counter int64
	var sw sync.WaitGroup

	for i := 0; i < 10; i++ {
		sw.Add(1)

		go func(sw *sync.WaitGroup) {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
			sw.Done()
		}(&sw)
	}

	sw.Wait()
	fmt.Printf("合计总数为： %d \n", counter)
}
