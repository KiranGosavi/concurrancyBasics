package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//A program to fetch the book data from cache or from database if cache miss

var cache =map[int]Book{}
var rnd =rand.New(rand.NewSource(time.Now().UnixNano()))

func main(){
	wg :=&sync.WaitGroup{}
	ch :=make(chan Book)
	dbCh :=make(chan Book)
	m :=&sync.RWMutex{}


	for i:=0;i<7;i++{
		id :=rnd.Intn(6)+1
		wg.Add(2)

		//Goroutine to fetch book from cache and send it on channel
		go func(wg *sync.WaitGroup,ch chan<- Book,id int,m *sync.RWMutex){
			b,ok:=queryCache(id,m)
			if ok{
				ch<- b
				//fmt.Printf("\nBook found in cache %v",b)
			}
			wg.Done()
		}(wg,ch,id,m)

		//Goroutine to fetch book from database and send it on channel
		go func(wg *sync.WaitGroup,dbCh chan<- Book,id int,m *sync.RWMutex){
			b,ok:=queryDatabase(id)
			if ok{
				//To avoid concurrent map writes
				m.Lock()
				cache[id]=b
				m.Unlock()
				dbCh <- b
				//fmt.Printf("\nBook found in database %v",b)
			}
			wg.Done()
		}(wg,dbCh,id,m)

		//Goroutine to receive book information from db channel or cache channel
		go func(ch, dbCh <-chan Book){
			select{
					case b := <-ch :
						fmt.Printf("\nBook found in cache %v",b)
					case b:= <-dbCh:
						fmt.Printf("\nBook found in database %v",b)
					default:
						fmt.Printf("\nNon-blocking select")
			}
		}(ch,dbCh)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("Here")
	wg.Wait()
	fmt.Println("after Here")
}

//queryCache function to fetch data from cache
func queryCache(id int,m *sync.RWMutex)(Book,bool){
	//To avoid concurrent map read
	m.RLock()
	defer m.RUnlock()
	b,ok :=cache[id]
	if ok{
		return b,true
	}
	return Book{},false
}

//queryDatabase function to fetch data from database
func queryDatabase(id int)(Book,bool){
	for _,b :=range books{
		if b.ID == id {
			return b,true
		}
	}
	return Book{},false
}