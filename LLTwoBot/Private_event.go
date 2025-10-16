package main

import (
	"LLTwoBot/mod/MatchingJson"
)

// Private 私聊消息事件
func Private(data MatchingJson.Private) {
	println(data.Message)

}
