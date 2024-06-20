package framework

import (
    "github.com/tyler66737/framework/math"
)

func Fib(n int) int {
    if n<=0{
        return 0
    }
    if n==1{
        return 1
    }
    return math.Add(Fib(n-1),Fib(n-2))
}
