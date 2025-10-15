package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cake struct {
	BakedBy  int
	BakeTime time.Duration
	PackedBy int
	PackTime time.Duration
}

func main() {
	K := 20 
	N := 3  
	M := 2  

	bakeChan := make(chan Cake)
	packChan := make(chan Cake)

	var wgBake sync.WaitGroup
	var wgPack sync.WaitGroup

	for i := 1; i <= N; i++ {
		wgBake.Add(1)
		go func(id int) {
			defer wgBake.Done()
			for c := 0; c < K/N; c++ {
				T1 := time.Duration(rand.Intn(50)) * time.Millisecond
				time.Sleep(T1)
				bakeChan <- Cake{BakedBy: id, BakeTime: T1}
			}
		}(i)
	}

	go func() {
		wgBake.Wait()
		close(bakeChan)
	}()

	for j := 1; j <= M; j++ {
		wgPack.Add(1)
		go func(id int) {
			defer wgPack.Done()
			for cake := range bakeChan {
				T2 := time.Duration(rand.Intn(100)) * time.Millisecond
				time.Sleep(T2)
				cake.PackedBy = id
				cake.PackTime = T2
				packChan <- cake
			}
		}(j)
	}

	go func() {
		wgPack.Wait()
		close(packChan)
	}()

	count := 0
	for cake := range packChan {
		fmt.Println(cake)
		count++
	}

	fmt.Println("Total cakes:", count)
}
