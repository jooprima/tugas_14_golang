package main

import "fmt"
import "os/exec"

func main() {
	var hasil, _ = exec.Command("help").Output()
	fmt.Println(string(hasil))
}
