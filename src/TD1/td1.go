package td1

import (
	"math/rand"
)

func Fill(sl []int) {
	for i := 0; i < len(sl); i++ {
		sl[i] = rand.Int()
	}
}

func Moyenne(sl []int) float64 {
	var sum, i int

	for i = 0; i < len(sl); i++ {
		sum += sl[i]
	}
	return float64(sum) / float64(len(sl))
}

func ValeursCentrales(sl []int) []int {
	//sort sl in ascending order
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl); j++ {
			if sl[i] < sl[j] {
				sl[i], sl[j] = sl[j], sl[i]
			}
		}
	}
	//return the middle value(s)
	if len(sl)%2 == 0 {
		return []int{sl[len(sl)/2], sl[len(sl)/2-1]}
	} else {
		return []int{sl[len(sl)/2]}
	}
}

func Plus1(sl []int) {
	for i := 0; i < len(sl); i++ {
		sl[i]++
	}
}

func Compte(tab []int) {
	//print the whole array
	for i := 0; i < len(tab); i++ {
		print(tab[i], " ")
	}
}
