package main

import ("fmt")

func HelloWorld(str string) string {
	return fmt.Sprintf("hello %s", "world")
}

func main() {
	fmt.print(HelloWorld("WuYunlong"))
}