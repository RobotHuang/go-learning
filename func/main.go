package main

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	tempCoins := coins
	for _, s := range users {
		for _, x := range s {
			if x == 'e' || x == 'E' {
				if tempCoins >= 1 {
					distribution[s]++
					tempCoins--
				} else {
					continue
				}
			} else if x == 'i' || x == 'I' {
				if tempCoins >= 2 {
					distribution[s] += 2
					tempCoins -= 2
				} else {
					continue
				}
			} else if x == 'o' || x == 'O' {
				if tempCoins >= 3 {
					distribution[s] += 3
					tempCoins -= 3
				} else {
					continue
				}
			} else if x == 'u' || x == 'U' {
				if tempCoins >= 4 {
					distribution[s] += 4
					tempCoins -= 4
				} else {
					continue
				}
			}
		}
	}
	return tempCoins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	//fmt.Println(distribution)
	fmt.Printf("%#v", distribution)
}
