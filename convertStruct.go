package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func convertStruct(input []string){
	//ansStrs := make([]string,0)

	for _,line := range input{

		//skip some lines
		if line == "" || (len(line)>2 && line[0:2] == "//") || (len(line)>0 && line[len(line)-1:] != ";"){
			fmt.Println(line)
			continue
		}

		//do convert logic
		//have ';'
		line = doConvertStruct(line)

		//print
		fmt.Println(line)
	}

}

func doConvertStruct(codeLine string)string{
	convertedCodeLine := ""
	//delete ";"  suffix
	if len(codeLine) == 0{
		return codeLine
	}
	codeLine = codeLine[0:len(codeLine)-1]
	lineSice := strings.Fields(codeLine) // [int a]
	if len(lineSice) < 2{
		return codeLine
	}else{
		if lineSice[0] == "struct"{
			lineSice = lineSice[1:]
		}
		//
		typ := lineSice[0]
		name := lineSice[1]
		//get "**...." from name
		flag := 0
		for _,c := range name{
			if c != '*'{
				break
			}else {
				flag++
			}
		}
		//delete "**.."from name
		stars := name[0:flag]
		name = name[flag:]
		//add prefix stars in typ
		typ =  stars + typ
		//package ans
		convertedCodeLine = name + "	" + typ

	}
	//int *a,*b,*c,*d; ---> a,b,c,d *int
	//obj *name; ---> *name obj   ---> name *obj
	//struct structName *name; ----> name *structName

	return convertedCodeLine
}




func handleInputStruct() []string{
	retStrs := make([]string,0)
	input := bufio.NewReader(os.Stdin)
	fmt.Println("converted result is :\n")
	for{
		curLine,err := input.ReadString('\n')
		curLine = strings.TrimSpace(curLine)
		panicErr(err)
		//check input illegal
		if len(retStrs) == 0 && len(curLine)>6 && curLine[0:6] != "struct"{
			fmt.Println("please input legal struct C code")
			os.Exit(0)
		}
		retStrs = append(retStrs,curLine)
		//check finish flag
		if curLine == "};"{
			break
		}
	}

	return retStrs
}

func main()  {

	//--------------input struct C code---------
	input := handleInputStruct()

	//-------------handle input-------------
	convertStruct(input)

}


func panicErr(err error){
	if err != nil{
		panic(fmt.Errorf("---------%v",err))
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
long long microseconds, calls;
};
 **/

