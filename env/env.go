package env

import (
	"log"
	"os"
)

// 環境変数のkey
const (
	envChannelID = "CHANNEL_ID"
	envToken     = "TOKEN"
)

// 環境変数のvalue
var (
	ChannelID string
	Token     string
)

// 必須環境変数の読み込み
// 未定義の場合はプログラムを終了させる
func required(key string) string {
	value, isDefined := os.LookupEnv(key)
	if !isDefined {
		log.Fatalln("Undefined environment variable \"" + key + "\"")
	}
	return value
}

// オプショナル環境変数の読み込み
// 未定義の場合はvalueには空文字列が入る
func optional(key string) string {
	value, isDefined := os.LookupEnv(key)
	if !isDefined {
		value = ""
	}
	return value
}

// 環境変数をセット
func init() {
	ChannelID = required(envChannelID)
	Token = required(envToken)
}
