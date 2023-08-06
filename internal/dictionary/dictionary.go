package dictionary

import (
	"strings"
)

type Dictionary struct {
	dictionaryRu      map[string]string
	dictionaryRuMorze map[string]string
	dictionaryEn      map[string]string
	dictionaryEnMorze map[string]string
	mod               int
}

func NewDictionary(mod int) *Dictionary {
	return &Dictionary{
		dictionaryRu:      DictionaryMapRu,
		dictionaryEn:      DictionaryMapEn,
		dictionaryEnMorze: DictionaryMapEnMorze,
		dictionaryRuMorze: DictionaryMapRuMorze,
		mod:               mod,
	}
}

func (d *Dictionary) ParseByWorld(lang int, msg string) string {
	dict := d.selectMap(lang)
	if d.mod == 0 {
		v, ok := dict[msg]
		if ok {
			return v
		}
	} else {
		res := ""
		for _, w := range strings.Split(msg, "") {
			v, ok := dict[strings.ToUpper(w)]
			if ok {
				res += v
				res += " "
			}
		}
		return res
	}
	return "Not valid"
}

func (d *Dictionary) Parse(lang int, msg string) string {
	res := ""
	dict := d.selectMap(lang)
	if d.mod == 0 {
		for _, word := range parseString(msg) {
			v, ok := dict[word]
			if ok {
				res += v + " "
			} else {
				res += " (not valid - '" + word + "' ) "
			}
		}
	} else {
		for _, word := range strings.Split(msg, "") {
			v, ok := dict[strings.ToUpper(word)]
			if ok {
				res += v + " "
			} else {
				res += " (not valid - '" + word + "' ) "
			}
		}
	}
	return res
}

func (d *Dictionary) selectMap(lang int) map[string]string {
	if d.mod == 0 {
		if lang == 1 {
			return d.dictionaryRu
		} else {
			return d.dictionaryEn
		}
	} else {
		if lang == 1 {
			return d.dictionaryRuMorze
		} else {
			return d.dictionaryEnMorze
		}
	}
}

func parseString(msg string) []string {
	res := make([]string, 0, len(msg))
	interimMsg := ""
	for _, v := range strings.Split(msg, "") {
		if matchInt(v) {
			interimMsg += v
		} else {
			res = append(res, interimMsg)
			interimMsg = ""
		}
	}
	res = append(res, interimMsg)
	return res
}

func matchInt(msg string) bool {
	if msg == "0" || msg == "1" {
		return true
	}
	return false
}
