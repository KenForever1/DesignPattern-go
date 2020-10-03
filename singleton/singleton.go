package singleton

import "fmt"

type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}
func main() {
	fmt.Println("vim-go")
}
