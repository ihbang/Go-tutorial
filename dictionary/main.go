package main

import (
	"fmt"

	"github.com/ihbang/go-tutorial/dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	newWord := "hello"
	newDef := "Greeting"

	err1 := dictionary.Add(newWord, newDef)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println()

	def, err2 := dictionary.Search(newWord)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Print("Defenition of the word \"", newWord, "\" is \"", def, "\"\n")
	}
	fmt.Println()

	err3 := dictionary.Add(newWord, newDef)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println()

	err4 := dictionary.Update(newWord, "world")
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Print("Defenition of the word \"", newWord, "\" is updated to \"", "world", "\"\n")
	}
	fmt.Println()

	def, err5 := dictionary.Search(newWord)
	if err5 != nil {
		fmt.Println(err5)
	} else {
		fmt.Print("Defenition of the word \"", newWord, "\" is \"", def, "\"\n")
	}
	fmt.Println()

	err6 := dictionary.Delete(newWord)
	if err6 != nil {
		fmt.Println(err6)
	} else {
		fmt.Print("The word \"", newWord, "\" is deleted\n")
	}
	fmt.Println()

	def, err7 := dictionary.Search(newWord)
	if err7 != nil {
		fmt.Println(err7)
	} else {
		fmt.Print("Defenition of the word \"", newWord, "\" is \"", def, "\"\n")
	}
}
