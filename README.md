# gFromC
 a simple tool convert C code to Golang

将C语言翻译为Go语言的简单工具

## How to use convert C struct 

Features:

- [x] Support basic operations
- [ ] Support multiple arg name ,eg: int a,b,c,d
- [ ] Support long long

 


```go

go run convertStruct.go < input.txt > ans.txt


paste C struct code in ./input.txt 

converted result is in ./ans.txt

```

eg:

paste this in terminal 

```go
struct redisCommand {
    
    //annotation
    char *name;
    
};
```