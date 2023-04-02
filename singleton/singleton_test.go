package singleton

import (
	"reflect"
	"sync"
	"testing"
)

const parCount = 100

func TestNewInstance(t *testing.T) {
	s1 := NewInstance()
	s2 := NewInstance()
	if !reflect.DeepEqual(s1, s2) {
		t.Fatal("singleton model error")
	}
	if s1 != s2 {
		t.Fatal("singleton model error")
	}
}

func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	var (
		instances [parCount]Singleton
		wg        sync.WaitGroup
	)
	wg.Add(parCount)
	for i := 0; i < parCount; i++ {
		go func(index int) {
			//阻塞 并发执行
			<-start
			instances[index] = NewInstance()
			wg.Done()
		}(i)
	}
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("parallel get instance error")
		}
	}
}
