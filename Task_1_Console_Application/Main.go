package main

/*
	Есть последовательность символов (задаётся в консоли).
	Нужно посчитать количество одинаковых символов в последовательности и вывести на экран символы в порядке частности.
	Необходимо использовать:
		- условия
		- циклы
		- slices
		- map
 */

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Pair struct {
	Key string
	Value int
}

type PairList []Pair	// slices done

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

func rankByWordCount(wordFrequencies map[string]int) PairList{
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func createStringFromPairList(list PairList) string {
	var str string
	for _, pair := range list {
		for i := 0; i < pair.Value; i++ {
			str += pair.Key
		}
	}
	return str
}

func main()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	charAmount := make(map[string]int)	// map done
	for _, symbol := range text {		// циклы done
		if symbol != '\n' {				// условия done
			_, exist := charAmount[string(symbol)]
			if exist {
				charAmount[string(symbol)]++
			} else {
				charAmount[string(symbol)] = 1
			}
		}
	}

	fmt.Println(createStringFromPairList(rankByWordCount(charAmount)))
}

