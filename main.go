package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)

	reader := bufio.NewReader(os.Stdin)

}
