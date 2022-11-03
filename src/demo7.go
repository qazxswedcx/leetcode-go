package main

//var wg sync.WaitGroup
//var i int = 100
//var lock sync.Mutex

/*
func add() {

		defer wg.Done()
		lock.Lock()
		i += 1
		fmt.Println("i++:", i)
		lock.Unlock()
	}

	func sub() {
		defer wg.Done()
		lock.Lock()
		i -= 1
		fmt.Println("i--:", i)
		lock.Unlock()
	}
*/
//func main() {
//	/*	for i := 0; i < 100; i++ {
//			wg.Add(1)
//			go add()
//			wg.Add(1)
//			go sub()
//		}
//		wg.Wait()
//		time.Sleep(time.Second)*/
//
//	ticker := time.NewTicker(time.Second) //周期性实现
//	counter := 1
//	for _ = range ticker.C {
//		fmt.Println("ticker")
//		counter++
//		if counter >= 5 {
//			ticker.Stop()
//			break
//		}
//	}
//}
