package main

import "fmt"

func main() {
	var a = [...]int{1, 3, 5, 7, 8}
	var sum int = 0
	for _, v := range a {
		sum += v
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
	fmt.Println(sum)

	for i := 0; i < len(a); i++ {
		tmp := 8 - a[i]
		for j := i + 1; j < len(a); j++ {
			if tmp == a[j] {
				fmt.Println(i, j)
				break
			}
		}
	}

	var b = make([]string, 5, 10)
	fmt.Println(b)
	for i := 0; i < 10; i++ {
		b = append(b, fmt.Sprintf("%v", i))
	}
	fmt.Println(b)

	//切片注意的点，append如果底层数组容量还够，修改底层数组，不够开辟新的空间
	var c []int
	c = a[:3]
	fmt.Printf("c=%v, cptr=%p\n", c, c)
	//c = append(c, 9)
	c = append(c, []int{9, 10, 11}...)
	fmt.Printf("a=%v, aptr=%p\n", a, &a)
	fmt.Printf("c=%v, cptr=%p\n", c, c)

	d1 := []int{1, 2, 3}
	d2 := make([]int, 3, 3)
	copy(d2, d1)
	f1(&d2)
	fmt.Println(d1, d2)
}

func f1(s *[]int) {
	*s = append(*s, 4)
}
