package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

type MetaCommandResult int

const (
	MetaCommandSuccess MetaCommandResult = iota
	MetaCommandUnrecognizedCommand
)

type PrepareResult int

const (
	PrepareSuccess PrepareResult = iota
	PrepareUnrecognizedStatement
)

type StatementType int

const (
	StatementInsert StatementType = iota
	StatementSelect
)

type Statement struct {
	Type StatementType
}

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

func DoMetaCommand(buffer string) (MetaCommandResult, error) {

	if buffer == ".exit" {
		os.Exit(0)
	} else {
		return MetaCommandUnrecognizedCommand, nil
	}

	return 1, errors.New("No fucking way")

}

func PrepareStatement(buffer string, statement Statement) PrepareResult {

	if strings.HasPrefix(buffer, "insert") {
		statement.Type = StatementInsert
		return PrepareSuccess
	}

	if buffer == "select" {
		statement.Type = StatementSelect
		return PrepareSuccess
	}

	return PrepareUnrecognizedStatement
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

		// FIXME: Empty buffer error
		if string((*input)[0]) == "." {
			CloseInputBuffer(inputBuffer)
			close(inputChan)
			os.Exit(0)
		} else {
			fmt.Printf("Unrecognized command '%s'.\ndb > ", *input)
		}
	}
}
