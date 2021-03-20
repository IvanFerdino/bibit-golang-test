package main

import (
	"log"
	"sort"
)

func main() {
	seedData:=make([]string,0)
	seedData=append(seedData,"kita","atik","tika","aku","kia","makan","kua")

	myMap:=make(map[string][]string)

	for _,d:=range seedData{
		runes:=[]rune(d)
		sort.Sort(sortRunes(runes)) //sort ascending
		myMap[string(runes)] = append(myMap[string(runes)],d)
	}

	log.Println("Anagram grouping result: ")
	for key,val:=range myMap{
		log.Printf("%+v: %+v\n",key,val)
	}
}


type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}
