package main

import (
	"bufio"
	"gomorz/internal/dictionary"
	"os"
	"strings"
)

func main() {
	var dict *dictionary.Dictionary
	reader := bufio.NewReader(os.Stdin)

	mod := 0
	lang := 0
	translate := 0
	start := true

	for {
		if start {
			println("Select mode: ( 0 - Morze-Lang, 1 - Lang-Morze )")
			rawText, _ := reader.ReadString('\n')
			text := strings.Split(rawText, "\n")[0]
			if text != "0" && text != "1" {
				println("Not valid value")
				continue
			} else {
				if text == "1" {
					translate = 1
				}
			}
			dict = dictionary.NewDictionary(translate)

			println("Select a language: ( 0 - English, 1 - Русский )")
			rawText, _ = reader.ReadString('\n')
			text = strings.Split(rawText, "\n")[0]
			if text != "0" && text != "1" {
				println("Not valid value")
				continue
			} else {
				if text == "1" {
					lang = 1
				}
			}

			if lang == 1 {
				println("Выберите в каком режиме будете набирать: ( 0 - по словам, 1 - целиком )")
			} else {
				println("Choose in which mode you will type: ( 0 - by words, 1 - in whole )")
			}

			rawText, _ = reader.ReadString('\n')
			text = strings.Split(rawText, "\n")[0]
			if text != "0" && text != "1" {
				println("Not valid value")
				continue
			} else {
				if text == "1" {
					mod = 1
				}
				start = false
			}
		}

		println("Enter text: ")

		if mod == 0 {
			text, _ := reader.ReadString('\n')
			println("Result: ", dict.ParseByWorld(lang, strings.Split(text, "\n")[0]))
			print("\n")
		} else {
			text, _ := reader.ReadString('\n')
			println("Result: ", dict.Parse(lang, strings.Split(text, "\n")[0]))
			print("\n")
		}
	}
}
