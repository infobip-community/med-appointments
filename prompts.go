package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func promptOptionRange(message string, nOptions int) int {
	choice := 0
	for choice < 1 || choice > nOptions {
		choice, _ = strconv.Atoi(promptText(fmt.Sprintf("%s\n(1-%d): ", message, nOptions)))
	}

	return choice
}

func promptText(message string) string {
	fmt.Println()
	if len(message) > 0 {
		fmt.Print(message)
	}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.Trim(text, " \n\t")

	return text
}

func promptYesNo(message string) bool {
	fmt.Println()
	text := ""
	for text != "y" && text != "n" {
		text = strings.ToLower(promptText(fmt.Sprintf("%s\n(Y/N/y/n): ", message)))
	}
	if strings.ToLower(text) == "y" {
		return true
	}

	return false
}

func promptChannel(message string) int {
	fmt.Println()

	return promptOptionRange(message, 2)
}

func promptDate(message string, dates []time.Time, amount int) time.Time {
	message = fmt.Sprintf("%s\n", message)
	for i := 0; i < amount; i++ {
		message = fmt.Sprintf("%s %d) %s\n", message, i+1, dates[i])
	}

	choice := promptOptionRange(message, amount)
	return dates[choice-1]
}
