package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{} // goroutineの完了を待つ機構
	wg.Add(2)

	fmt.Println("Starting Go Routines!!")
	go func() { // goを付けると並行処理になる
		defer wg.Done() // deferはこのScopeを抜ける際に必ず呼び出されるもの

		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}()

	go func() {
		defer wg.Done()

		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()

	fmt.Println("Waiting to Finish")
	wg.Wait() // sync.WaitGroup は Wait() を呼び出すと Add() を呼び出した回数から Done() を呼び出した回数を引いて 0 になるまで待機

	fmt.Println("\nTerminating Program")
}
