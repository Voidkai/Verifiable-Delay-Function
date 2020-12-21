package main

import (
	"fmt"
	"math/big"
	"time"
)

func main(){
	var security uint64 = 300
	n := 24
	T := uint64(1 << 25)
	N := new(big.Int).Set(Setup(security))

	fmt.Println("security:",security,"T: ",T,"= 2^",n, "N:",N)

	instance := Generate(N, T, security)

	Time_Start := time.Now()
	instance.NaiveSolve()
	fmt.Println("time:", time.Since(Time_Start))
	fmt.Println(instance.challenge , instance.y)

}
