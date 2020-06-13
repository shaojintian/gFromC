package main

var FIELDS_MAP = map[string]string{
	//custom config
	"sds":"string",
	"redis": "gkv",
	"list": "list.List",

	"long long": "int64",
	"long": "int32",
	"unsigned long":"uint32",
	"int": "int",
	"char": "rune",
}
