package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var choice string

	fmt.Println("What program do you want to run?\n1. for strings.go\n2. for slice.go\n3. for search.go\n4. for maps\n5. for words")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		stringProgram()
	case "2":
		sliceProgram()
	case "3":
		searchProgram()
	case "4":
		mapProgram()
	case "5":
		wordProgram()
	default:
		fmt.Println("Invalid choice")
	}
}


func stringProgram(){
	s := "èlliot"

	fmt.Printf("%8T %[1]v \n", s)
	fmt.Printf("%8T %[1]v \n", []rune(s))
	b := []byte(s)
	fmt.Printf("%8T %[1]v %d\n",b, len(b))
}

func sliceProgram(){
	t := []byte("string")

	fmt.Println(len(t), t)

	fmt.Println(t[2])
	fmt.Println(t[:2])
	fmt.Println(t[2:])
	fmt.Println(t[3:5], len(t[3:5]))

	h:= append(t, 'a')

	fmt.Println(h)
}

func searchProgram(){
	if len(os.Args)<3{
		fmt.Println(os.Stderr, "not enough args")
		os.Exit(-1)
	}

	old, new:= os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan(){
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s,new)

		fmt.Println(t)
	}
}


func mapProgram(){
	m:= map[string]int{
		"Alice":20,
		"Bob":19,
	}

	//Updating a map
	m["charlie"]=34
	m["adnan"]=21

	age:=m["adnan"]
	
	fmt.Println(age)

	for key, value := range m{
		fmt.Println(key,value)	
	}
}

func wordProgram() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)

	words := make(map[string]int)

	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "unique words")

	type kv struct{
		key string
		val int
	}

	var ss []kv
	
	for k, v := range words{
		ss= append(ss,kv{k,v})
	}

	sort.Slice(ss,func(i,j int) bool {
		return ss[i].val >ss[j].val
	})

	for _, s:=range ss[:4] {
		fmt.Println(s.key, "appears", s.val, "times")
	}
}
