package lockey

import (
	"sync"
	"testing"
)

var abacusRed = 0
var abacusYellow = 0
var abacusBlue = 0

var lockey = New()

const total = 300000
const per = 100000

func Test_Lock(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(total)

	//inc red
	go func() {
		for i := 0; i < per; i++ {
			go func() {
				IncRed()
				wg.Done()
			}()
		}
	}()

	//inc yellow
	go func() {
		for i := 0; i < per; i++ {
			go func() {
				IncYellow()
				wg.Done()
			}()
		}
	}()

	//inc blue
	go func() {
		for i := 0; i < per; i++ {
			go func() {
				IncBlue()
				wg.Done()
			}()
		}
	}()
	wg.Wait()

	if abacusRed != per {
		t.Errorf("got %d, want %d for red", abacusRed, per)
	}

	if abacusYellow != per {
		t.Errorf("got %d, want %d for yellow", abacusYellow, per)
	}

	if abacusBlue != per {
		t.Errorf("got %d, want %d for blue", abacusBlue, per)
	}

}

func IncRed() {
	lockey.Lock("red")
	defer lockey.Unlock("red")
	abacusRed++
}

func IncYellow() {
	lockey.Lock("yellow")
	defer lockey.Unlock("yellow")
	abacusYellow++
}

func IncBlue() {
	lockey.Lock("blue")
	defer lockey.Unlock("blue")
	abacusBlue++
}
