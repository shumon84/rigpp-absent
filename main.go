package main

import (
	"errors"
	"fmt"
	"github.com/nlopes/slack"
	"github.com/shumon84/rigpp-absent/env"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// タイムスタンプを日付に変換
func ParseTimestamp(timestamp string) (string, error) {
	timeStrings := strings.Split(timestamp, ".")
	unixTime, err := strconv.Atoi(timeStrings[0])
	if err != nil {
		return "", err
	}
	datetime := time.Unix(int64(unixTime), 0)
	return datetime.Format("2006/01/02 03:04:05"), nil
}

// textから改行を削除する
func DistinctCRLF(text string) string {
	return strings.Replace(text, "\n", " ", -1)
}

// textから学籍番号を抽出する
func TextToStudentID(text string) (string, error) {
	stripText := strings.Replace(text, "-", "", 10)
	reg, err := regexp.Compile("[0-9]{11}")
	if err != nil {
		log.Fatal(err.Error())
	}
	studentID := string(reg.Find([]byte(stripText)))
	if studentID == "" {
		return "", errors.New("StudentID is not found in 「" + DistinctCRLF(text) + "」")
	}
	return studentID, nil
}

func main() {
	// クライアントの生成
	api := slack.New(env.Token)

	// チャンネルの履歴を取得
	hp := slack.NewHistoryParameters()
	hp.Count = 1000
	history, err := api.GetChannelHistory(env.ChannelID, hp)
	if err != nil {
		log.Fatal(err.Error())
	}

	// 欠席連絡された日付と学籍を抽出して出力
	for _, v := range history.Messages {
		date, err := ParseTimestamp(v.Timestamp)
		if err != nil {
			fmt.Errorf(err.Error())
			continue
		}
		studentID, err := TextToStudentID(v.Text)
		if err != nil {
			fmt.Errorf(err.Error())
			continue
		}
		fmt.Println(date, studentID)
	}
}
