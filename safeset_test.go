package safeset

import (

	"sync/atomic"
	"testing"
)


func TestSize(t *testing.T){
	cmp := func (a,b interface{}) int {
		a1 := a.(int)
		b1 := b.(int)
		if a1 == b1 {
			return 0
		}
		if(a1>b1){
			return 1
		}
		return -1
	}
	set := NewSafeSet(cmp)
	done := make(chan bool)

	for i:=0;i<10;i++{

		go func (ind int){		
			 set.SafeInsert(ind)
			 done <- true 
		}(i)
	}
	
	for i:=0;i<10;i++{
		<-done
	}
	
	if set.SafeSize() != 10{
		t.Fatalf("Expected Size %d but Found %d!",10,set.SafeSize())
	}

	for i:=0;i<10;i++{
		go func (ind int){
			 set.SafeRemove(ind)
			 done<-true
		}(i)
	}
	for i:=0;i<10;i++{
		<-done
	}
	if set.SafeSize() != 0{
		t.Fatalf("Expected Size %d but Found %d!",0,set.SafeSize())
	}
}

func TestSafeGet(t *testing.T){
	cmp := func (a,b interface{}) int {
		a1 := a.(int)
		b1 := b.(int)
		if a1 == b1 {
			return 0
		}
		if(a1>b1){
			return 1
		}
		return -1
	}
	set := NewSafeSet(cmp)
	done := make(chan bool)

	for i:=0;i<10;i++{

		go func (ind int){		
			 set.SafeInsert(ind)
			 done <- true 
		}(i)
	}
	
	for i:=0;i<10;i++{
		<-done
	}
	res := make([]int32,10)

	for i:=0;i<10;i++{
		go func (ind int){		
			atomic.AddInt32(&res[(set.SafeGet(ind)).(int)],1)
			done <- true 
	   }(i)
	}

	for i:=0;i<10;i++{
		<-done
	}

	for i:=0;i<10;i++{
		if res[i]!=1{
			t.Fatalf("Error in  Get Function,Returned Element Mistakenly")
		}
	}
}
