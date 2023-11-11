package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type InputBuffer struct {
	buffer       *string
	bufferLength int
	inputLength  int
}

func NewInputBuffer() *InputBuffer {
	inputBuffer := InputBuffer{}
	return &inputBuffer
}

func PrintPrompt() {
	fmt.Print("db > ")
}

func CloseInputBuffer(inputBuffer *InputBuffer) {
	inputBuffer.buffer = nil
}

func ReadInput(inputBuffer *InputBuffer, wg *sync.WaitGroup, inputChan chan *string) {
	defer wg.Done()
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error Reading input")
		os.Exit(1)
	}

	inputBuffer.inputLength = len(input) - 1
	trimmedInput := strings.TrimSpace(input)
	inputBuffer.buffer = &trimmedInput
	(*inputBuffer).bufferLength = len(input) - 1
	inputChan <- inputBuffer.buffer

}

func main() {
	inputBuffer := NewInputBuffer()
	var wg sync.WaitGroup
	inputChan := make(chan *string)

	go func() {
		for {
			PrintPrompt()
			wg.Add(1)
			go ReadInput(inputBuffer, &wg, inputChan)
			wg.Wait()
		}
	}()

	for {
		input := <-inputChan

		if *input == ".exit" {
			CloseInputBuffer(inputBuffer)
			close(inputChan)
			os.Exit(0)
		} else {
			fmt.Printf("Unrecognized command '%s'.\ndb > ", *input)
		}
	}
}
