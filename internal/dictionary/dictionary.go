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
	if lang == 1 {
		if d.mod == 0 {
			v, ok := d.dictionaryRu[msg]
			if ok {
				return v
			}
		} else {
			res := ""
			for _, w := range strings.Split(msg, "") {
				v, ok := d.dictionaryRuMorze[strings.ToUpper(w)]
				if ok {
					res += v
					res += " "
				}
			}
			return res
		}
	} else {
		if d.mod == 0 {
			v, ok := d.dictionaryEn[msg]
			if ok {
				return v
			}
		} else {
			res := ""
			for _, w := range strings.Split(msg, "") {
				v, ok := d.dictionaryEnMorze[strings.ToUpper(w)]
				if ok {
					res += v
					res += " "
				}
			}
			return res
		}
	}
	return "Not valid"
}

func (d *Dictionary) Parse(lang int, msg string) string {
	res := ""
	if d.mod == 0 {
		if lang == 1 {
			for _, word := range parseString(msg) {
				v, ok := d.dictionaryRu[word]
				if ok {
					res += v
				} else {
					res += " (not valid - '" + word + "' ) "
				}
			}

		} else {
			for _, word := range parseString(msg) {
				v, ok := d.dictionaryEn[word]
				if ok {
					res += v + " "
				} else {
					res += " (not valid - '" + word + "' ) "
				}
			}
		}
	} else {
		if lang == 1 {
			for _, word := range strings.Split(msg, "") {
				v, ok := d.dictionaryRuMorze[strings.ToUpper(word)]
				if ok {
					res += v + " "
				} else {
					res += " (not valid - '" + word + "' ) "
				}
			}

		} else {
			for _, word := range strings.Split(msg, "") {
				v, ok := d.dictionaryEnMorze[strings.ToUpper(word)]
				if ok {
					res += v
				} else {
					res += " (not valid - '" + word + "' ) "
				}
			}
		}
	}
	return res
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
