package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s:%d", p.Name, p.Age)
}

//
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	people := []Person{
		{Name: "Bob", Age: 31},
		{Name: "John", Age: 42},
		{Name: "Michael", Age: 17},
		{Name: "Jenny", Age: 26},
	}
	fmt.Println(people)

	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})

	fmt.Println(people)
}
