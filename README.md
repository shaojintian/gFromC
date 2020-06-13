# gFromC
 a simple tool convert C code to Golang

将C语言翻译为Go语言的简单工具

# Quick Usage V0.1

## How to convert C function 

Features:

- [x] Support simple functions conversion
- [x] Support channel/goroutine from pthread
- [ ] Support Go function type bind

## How to convert C struct 

Features:

- [x] Support basic operations
- [x] Support multiple arg name ,eg: int a,b,c,d
- [x] Support C array -> Go slice

 

```go

go run ./ < input.txt > ans.txt


paste C struct code in ./input.txt 

converted result is in ./ans.txt

```

eg:

paste this in input.txt

```c
struct structName {
    
    //xxxxxxxx
    char *name;
    
};
```

##