package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func convertStruct(input []string) {
	suffixAnnotation := ""

	for index, line := range input {

		// ""
		if len(line) == 0 {
			fmt.Println(line)
			continue
		}

		//----------len line > 0--------------
		//handle first line
		if index == 0 {
			lineSlice := strings.Fields(line)
			fmt.Println("type " + lineSlice[1] + " struct {")
			continue
		}
		//handle finish line
		if index == len(input)-1 {
			fmt.Println("}")
			continue
		}
		//skip "//....."
		if len(line) >= 2 && line[0:2] == "//" {
			fmt.Println(line)
			continue
		}

		//skip end 不是";"结尾的
		if line[len(line)-1:] != ";" {
			lineSlice := strings.Split(line, ";")
			if len(lineSlice) >= 2 {
				rawLine := line
				line = lineSlice[0] + ";"
				suffixAnnotation = rawLine[len(line):]
			} else {
				fmt.Println(line)
				continue
			}
		}

		//do convert logic
		//suffix have ';'
		line = doConvertStruct(line, suffixAnnotation)

		//print
		fmt.Println(line)
	}

}

func doConvertStruct(codeLine, suffixAnnotation string) string {
	convertedCodeLine := ""
	if len(codeLine) == 0 {
		return codeLine
	}
	//delete ";"  suffix
	codeLine = codeLine[0 : len(codeLine)-1]
	//split [type,type,...,name]
	lineSlice := strings.Fields(codeLine) // [int a]
	if lineSlice[0] == "struct" {
		lineSlice = lineSlice[1:]
	}
	if len(lineSlice) < 2 {
		return codeLine + suffixAnnotation
	} else if len(lineSlice) == 2{
		typ := lineSlice[0]
		name := lineSlice[1]
		//get "**...." from name
		stars,name := getPrefixStars(name)
		//add prefix stars in typ
		typ = stars + typ
		//package ans
		convertedCodeLine = name + "	" + typ + suffixAnnotation

	//len > 2
	}else {
		//[long long *a, *b, *c]
		//l := len(lineSlice)
		if lineSlice[0]==lineSlice[1] && lineSlice[0]=="long"{
			lineSlice[0] = FIELDS_MAP["long long"]
		}
		//[*a, *b, *c]
		typ := lineSlice[0]
		names := remove(lineSlice,1)[1:]

		stars := ""
		for index,name := range names{
			tmpStars,newName := getPrefixStars(name)
			names[index] = newName
			stars = tmpStars
		}
		typ = stars + typ
		ans := append(names,typ)
		convertedCodeLine = strings.Join(ans,"	")

	}
	//int *a,*b,*c,*d; ---> a,b,c,d *int
	//obj *name; ---> *name obj   ---> name *obj
	//struct structName *name; ----> name *structName

	return convertedCodeLine
}

func getPrefixStars(name string)(string,string){
	//get "**...." from name
	flag := 0
	for _, c := range name {
		if c != '*' {
			break
		} else {
			flag++
		}
	}
	//delete "**.."from name
	stars := name[0:flag]
	name = name[flag:]
	return stars,name
}
func handleInputStruct() []string {
	retStrs := make([]string, 0)
	input := bufio.NewReader(os.Stdin)

	for {
		curLine, err := input.ReadString('\n')
		curLine = strings.TrimSpace(curLine)
		panicErr(err)
		//check input illegal
		if len(retStrs) == 0 && len(curLine) <= 6 {
			continue
		}
		if len(retStrs) == 0 && len(curLine) > 6 && curLine[0:6] != "struct" {
			fmt.Println("please input legal struct C code")
			os.Exit(0)
		}
		retStrs = append(retStrs, curLine)
		//check finish flag
		if len(curLine) > 0 && curLine[0] == '}' {
			break
		}
	}

	return retStrs
}

func main() {

	//--------------input struct C code---------
	input := handleInputStruct()

	//-------------handle input-------------
	convertStruct(input)

}

func panicErr(err error) {
	if err != nil {
		panic(fmt.Errorf("---------%v", err))
	}
}

/**
struct redisCommand {

// 命令名字
char *name;

// 实现函数
redisCommandProc *proc;

// 参数个数
int arity;

// 字符串表示的 FLAG
char *sflags;


// 实际 FLAG
int flags;

// 从命令中判断命令的键参数。在 Redis 集群转向时使用。
redisGetKeysProc *getkeys_proc;


// 指定哪些参数是 key
int firstkey;
int lastkey;
int keystep;

// 统计信息
// microseconds 记录了命令执行耗费的总毫微秒数
// calls 是命令被执行的总次数
};
 **/
