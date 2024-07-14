package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader* bufio.Reader=bufio.NewReader(os.Stdin)
var writer* bufio.Writer=bufio.NewWriter(os.Stdout)

var dp[40001] int
var primes[] int
func is_prime(n int) bool{
	if n%2==0{
		return false
	}
	if n%3==0{
		return false
	}
	for i:=3;i*i<=n;i++{
		if(n%i==0){
			return false
		}
	}
	return true
}
func main(){
	defer writer.Flush()
	var N int
	fmt.Fscanln(reader,&N)
	dp[0]=1
	primes=append(primes, 2)
	primes=append(primes, 3)
	for i:=5;i<=N;i++{
		if is_prime(i){
			primes=append(primes, i)
		}
	}
	for i:=0;i<len(primes);i++{
		c:=primes[i]//금액동전
		for j:=c;j<=N;j++{
			dp[j]+=dp[j-c]%123456789
		}
	}
	fmt.Fprintln(writer,dp[N]%123456789)
}