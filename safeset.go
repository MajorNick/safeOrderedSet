package safeset

import (
	
	"sync"

	oset "github.com/MajorNick/orderedSet"
)
type SafeSet struct {
	set oset.OrderedSet
	locker  *sync.RWMutex

}

func NewSafeSet(compare func(a,b interface{})int)  SafeSet{
	
	safeSet := SafeSet{}
	safeSet.locker  = &sync.RWMutex{}
	safeSet.set = oset.NewSet(compare)


	return safeSet
}
func (st *SafeSet)SafeSize()int{
	st.locker.RLock()
	defer st.locker.RUnlock()
	return st.set.Size()
}

func (st *SafeSet)SafeGet(index int) interface{}{
	st.locker.RLock()
	defer st.locker.RUnlock()
	return st.set.Get(index)
}

func (st *SafeSet)SafeInsert(elem interface{})bool{
	st.locker.Lock()
	defer st.locker.Unlock()
	return st.set.Insert(elem)
}

func (st *SafeSet)SafeRemove(elem interface{})bool{
	st.locker.Lock()
	defer st.locker.Unlock()
	return st.set.Remove(elem)
}

func (st *SafeSet)SafeBsearch(elem interface{})int{
	st.locker.RLock()
	defer st.locker.RUnlock()
	return st.set.Bsearch(elem)
}

func (st *SafeSet)SafeToString()string{
	st.locker.RLock()
	defer st.locker.RUnlock()
	return st.set.ToString()
}


