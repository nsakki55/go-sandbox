package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "筑肥線九大学研都市駅"
	
	// 正規表現パターンで路線名を抽出
	re := regexp.MustCompile(`(.+)線`)
	matches := re.FindStringSubmatch(text)
	
	if len(matches) >= 2 {
		lineName := matches[1]
		fmt.Printf("路線名: %s\n", lineName)
	} else {
		fmt.Println("路線名が見つかりませんでした。")
	}
}
