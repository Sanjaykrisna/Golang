package golang

import (
	"fmt"
)

func fact(a int) int {
	if a <= 0 {
		return 1
	}
	return a * fact(a-1)
}

func fib(b int) int {
	if b < 2 {
		return 0
	}
	return fib(b-1) + fib(b+1)
}

func Power(x int, count int) int {
	if x <= 0 {
		return 0
	} else if count == 0 {
		return 1
	}

	res := 1

	for count > 0 {
		res = res * x
		count--
	}
	return res
}

func armstrongNum(c int) bool {
	if c > 0 && c <= 9 {
		return true
	}

	var t, digitcount, result int
	t = c
	for t > 0 {
		t = t / 10
		digitcount++
	}

	for c > 0 {
		h := c % 10
		result = result + Power(h, digitcount)
		c = c / 10
	}

	if c != result {
		return false
	}

	return true
}

func reverseInt(d int) int {
	if d >= 0 && d < 10 {
		return d
	}
	var result int
	for d > 0 {
		t := d % 10
		if t != 0 {
			result = result*10 + t
		}
		d = d / 10
	}
	return result
}

func amicableNum(e, f int) bool {
	var eRes, fRes int
	for i := 1; i < e; i++ {
		if e%i == 0 {
			eRes += i //sum of divisors
		}
	}

	for j := 1; j < f; j++ {
		if f%j == 0 {
			fRes += j
		}
	}

	if eRes == f && fRes == e { //1st sum of divisors should be equal to the another original number. Also for 2nd ..
		return true
	}

	return false
}

func disarumNum(g int) bool {
	var t, result, digitcount int
	t = g
	for t > 0 {
		t = t / 10
		digitcount++
	}

	t = g
	for t > 0 {
		h := t % 10
		result = result + Power(h, digitcount)
		digitcount--
		t = t / 10
	}

	if result == g {
		fmt.Println("Match Found.")
		return true
	} else {
		fmt.Println("Match Not Found.")
		return false
	}
}

func happyOrUnhappyNum(k int) bool {
	var t, digitcount int
	t = k
	for t > 0 {
		t = t / 10
		digitcount++
	}

	find := func(num, ct int) int {
		if ct == 0 {
			return 1
		}
		var res int
		for num > 0 {
			h := num % 10
			res = res + Power(h, ct)
			num = num / 10
		}
		return res
	}

	t = k
	m := make(map[int]bool)
	for t != 1 {
		// Have I seen this before?
		if _, b := m[t]; b {
			fmt.Println("Unhappy number.")
			return false
		}
		// Remember it
		m[t] = true
		t = find(t, digitcount)
	}

	if t == 1 {
		fmt.Println("Happy number.")
		return true
	}

	return false

}

func gcd(x, y int) int {
	var res int
	i := 1
	for i <= x && i <= y {
		if x%i == 0 && y%i == 0 {
			res = i
		}
		i++
	}
	return res
}

func primeNum(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			fmt.Println("It's not a prime number.")
			return false
		}
	}

	fmt.Println("It's a prime number.")
	return true
}

func removeDuplicates(sl []int) map[int]int {
	d := make(map[int]int)

	for _, num := range sl {
		if _, i := d[num]; i {
			continue
		}
		d[num] = num
	}
	return d
}
