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

// 与えられた日付から見てもっとも近い次の活動日を返す
func GetNextActivityDate(date time.Time) time.Time {
	NextActivityDateDuration := []int{2, 1, 0, 2, 1, 0, 3}
	return date.AddDate(0, 0, NextActivityDateDuration[date.Weekday()])
}

// タイムスタンプを日付に変換
func TimestampToDateTime(timestamp string) (time.Time, error) {
	timeStrings := strings.Split(timestamp, ".")
	second, err := strconv.Atoi(timeStrings[0])
	if err != nil {
		return time.Time{}, err
	}
	nanoSecond, err := strconv.Atoi(timeStrings[1])
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(int64(second), int64(nanoSecond)), nil
}

// 次の活動日を返す
func TimestampToDateString(timestamp string) (string, error) {
	dateTime, err := TimestampToDateTime(timestamp)
	if err != nil {
		return "", err
	}
	nextActivityDate := GetNextActivityDate(dateTime)
	return nextActivityDate.Format("2006/01/02 Mon"), nil
}

// textから改行を削除する
func RemoveNewLine(text string) string {
	return strings.Replace(text, "\n", " ", -1)
}

// textから学籍番号を抽出する
func ExtractStudentID(text string) (string, error) {
	stripText := strings.Replace(text, "-", "", 10)
	reg, err := regexp.Compile("[0-9]{11}")
	if err != nil {
		log.Fatal(err.Error())
	}
	studentID := string(reg.Find([]byte(stripText)))
	if studentID == "" {
		return "", errors.New("StudentID is not found in 「" + RemoveNewLine(text) + "」")
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

	// 欠席する活動日と学籍番号を抽出して出力
	for _, message := range history.Messages {
		dateString, err := TimestampToDateString(message.Timestamp)
		if err != nil {
			log.Println(err)
			continue
		}
		studentID, err := ExtractStudentID(message.Text)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(dateString, studentID)
	}
}
