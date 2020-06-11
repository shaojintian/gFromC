# gFromC
 a simple tool convert C code to Golang


## How to use convert C struct 

Features:

-[x] Support basic operations
-[ ] Support multiple arg name ,eg: int a,b,c,d
-[ ] Support long long

 


```go

go run convertStruct.go > ans.txt

paste C struct code in terminal AND press RETURN key in keyboard


```

eg:

paste this in terminal 

```go
struct redisCommand {
    
    //annotation
    char *name;
    
};
```