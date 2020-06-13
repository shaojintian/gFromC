package main

import (
	"reflect"
)

func removeStr(slice []string, s int) []string{
	return append(slice[:s], slice[s+1:]...)
}

func removeRune(slice []rune, s int) []rune{
	return append(slice[:s], slice[s+1:]...)
}


func remove(slice []interface{}, s int) []interface{}{
	return append(slice[:s], slice[s+1:]...)
}

// cannot force convert []type --> []interface{}
// interface{}  -> []interface{}
func toSlice(slice interface{})[]interface{}{
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic(v)
	}
	newSlice := make([]interface{},0)
	for i:=0;i<v.Len();i++{
		newSlice = append(newSlice,v.Index(i).Interface())
	}
	return newSlice
}

func interfaceToRuneSlice(slice []interface{})[]rune{
	if len(slice) == 0{
		return []rune{}
	}

	runeSlice := make([]rune,len(slice))
	for i:=0;i<len(runeSlice);i++{
		runeSlice[i] = slice[i].(rune)
	}
	return runeSlice
}

