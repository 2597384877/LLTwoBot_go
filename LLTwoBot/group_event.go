package main

import (
	"LLTwoBot/mod/MatchingJson"
)

// Group 群消息事件
func Group(data MatchingJson.Group) {
	println(data.Message)

}
