# What is Lockey

It is a simple key-based mutex locking package.

## Install

```bash
go get github.com/sertangulveren/lockey
```

## Usage

```go
var lk = lockey.New()

//...

lk.Lock("greatKey")
defer lk.Unlock("greatKey")

// Do a concurrent job
myGreatFunc()
//...
```

## Basic Example

```go
var lk = lockey.New()

func main() {
    //...
    for i := 0; i < bigNumber; i++ {
        go increaseMyMoney(i)
    }   
    //...
}

func increaseMyMoney(i int) {
    lk.Lock("makeMeRich")
    defer lk.Unlock("makeMeRich")
    
    money+=i
    addTransaction(i)
    //...
}
```