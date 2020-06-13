package main

func snakeToCamel(camelStr string)string{
	var snakeStr string
	if len(camelStr) ==0 {
		return ""
	}
	camelSlice := []rune(camelStr)
	for i:=1;i<len(camelSlice)-1;i++{
		lowerRune := camelSlice[i+1]
		if camelSlice[i]== '_' && lowerRune >= 'a' && lowerRune <= 'z'{
			camelSlice[i+1] =  lowerRune - 32
			camelSlice = removeRune(camelSlice,i)
		}
	}

	snakeStr = string(camelSlice)
	return snakeStr
}

