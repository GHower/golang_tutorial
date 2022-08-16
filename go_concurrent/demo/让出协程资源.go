package test

//func main() {
//	for i := 0; i < 5; i++ {
//		go func(no int) {
//			if no == 3 {
//				// 让出协程资源，降低执行的优先级
//				//runtime.Gosched()
//				// 让3号协程退出
//				//runtime.Goexit()
//				// 设置CPU逻辑核心数为4，返回原配置
//				v := runtime.GOMAXPROCS(4)
//				fmt.Println("原配置数:", v)
//			}
//
//			for j := 0; j < 10; j++ {
//				fmt.Printf("协程%d:%d\n", no, j)
//			}
//		}(i)
//	}
//	time.Sleep(3 * time.Second)
//}
