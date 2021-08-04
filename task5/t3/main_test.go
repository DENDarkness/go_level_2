package main

import (
	"math/rand"
	"sync"
	"testing"
)

var (
	mutex    sync.Mutex
	mutexRW  sync.RWMutex
	wg       sync.WaitGroup
	massages []int
)

func BenchmarkMutex1090(t *testing.B) {
	rand.Seed(89)
	massages = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	wg.Add(100)

	go func() {
		for i := 0; i < 10; i++ {
			go func() {
				rand.Seed(89)
				c := rand.Intn(20)
				mutex.Lock()
				massages[c] = rand.Intn(100)
				mutex.Unlock()
				wg.Done()
			}()
		}
	}()

	go func() {
		for i := 0; i < 90; i++ {
			go func() {
				mutex.Lock()
				c := rand.Intn(10)
				_ = massages[c]
				mutex.Unlock()
				wg.Done()
			}()
		}
	}()

	wg.Wait()

}

func BenchmarkMutex5050(t *testing.B) {
	rand.Seed(89)
	massages = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	wg.Add(100)

	go func() {
		for i := 0; i < 50; i++ {
			go func() {
				rand.Seed(89)
				c := rand.Intn(20)
				mutex.Lock()
				massages[c] = rand.Intn(100)
				mutex.Unlock()
				wg.Done()
			}()
		}
	}()

	go func() {
		for i := 0; i < 50; i++ {
			go func() {
				mutex.Lock()
				c := rand.Intn(10)
				_ = massages[c]
				mutex.Unlock()
				wg.Done()
			}()
		}
	}()

	wg.Wait()

}

func BenchmarkMutex9010(t *testing.B) {
	rand.Seed(89)
	massages = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	wg.Add(100)

	go func() {
		for i := 0; i < 90; i++ {
			go func() {
				rand.Seed(89)
				c := rand.Intn(20)
				mutex.Lock()
				massages[c] = rand.Intn(100)
				mutex.Unlock()
				wg.Done()
			}()
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			go func() {
				mutex.Lock()
				c := rand.Intn(10)
				_ = massages[c]
				mutex.Unlock()
				wg.Done()
			}()
		}
	}()

	wg.Wait()

}

func BenchmarkMutexRW1090(t *testing.B) {
	rand.Seed(89)
	massages = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	wg.Add(100)

	go func() {
		for i := 0; i < 10; i++ {
			go func() {
				rand.Seed(89)
				c := rand.Intn(20)
				mutex.Lock()
				massages[c] = rand.Intn(100)
				mutex.Unlock()
				wg.Done()
			}()
		}
	}()

	go func() {
		for i := 0; i < 90; i++ {
			go func() {
				mutexRW.RLock()
				c := rand.Intn(10)
				_ = massages[c]
				mutexRW.RUnlock()
				wg.Done()
			}()
		}
	}()

	wg.Wait()

}

func BenchmarkMutexRW5050(t *testing.B) {
	rand.Seed(89)
	massages = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	wg.Add(100)

	go func() {
		for i := 0; i < 50; i++ {
			go func() {
				rand.Seed(89)
				c := rand.Intn(20)
				mutex.Lock()
				massages[c] = rand.Intn(100)
				mutex.Unlock()
				wg.Done()
			}()
		}
	}()

	go func() {
		for i := 0; i < 50; i++ {
			go func() {
				mutexRW.RLock()
				c := rand.Intn(10)
				_ = massages[c]
				mutexRW.RUnlock()
				wg.Done()
			}()
		}
	}()

	wg.Wait()

}

func BenchmarkMutexRW9010(t *testing.B) {
	rand.Seed(89)
	massages = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	wg.Add(100)

	go func() {
		for i := 0; i < 90; i++ {
			go func() {
				rand.Seed(89)
				c := rand.Intn(20)
				mutex.Lock()
				massages[c] = rand.Intn(100)
				mutex.Unlock()
				wg.Done()
			}()
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			go func() {
				mutexRW.RLock()
				c := rand.Intn(10)
				_ = massages[c]
				mutexRW.RUnlock()
				wg.Done()
			}()
		}
	}()

	wg.Wait()

}
