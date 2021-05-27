package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	RandomKindNum        = iota // 纯数字
	RandomKindLower             // 纯小写字母
	RandomKindUpper             // 纯大写字母
	RandomKindLowerNum          // 数字+小字字母
	RandomKindUpperNum          // 数字+大写字母
	RandomKindUpperLower        // 大小写字母
	RandomKindAll               // 数字+大小写字母
)

func GenerateUUID() string {
	return strings.Replace(uuid.Must(uuid.NewUUID()).String(), "-", "", -1)
}

func GenerateRandomID(size int, kind int) string {
	index := kind
	if index > RandomKindAll || index < RandomKindNum {
		index = rand.Intn(3)
	}

	// '0'=48,'a'=97,'A'=65
	kinds := [][]int{{10, 48}, {26, 97}, {26, 65}}
	rand.Seed(time.Now().UnixNano())
	randomID := make([]byte, size)
	for i := 0; i < size; i++ {
		switch kind {
		case RandomKindLowerNum:
			index = rand.Intn(2)
		case RandomKindUpperNum:
			indexList := [2]int{0, 2}
			index = indexList[rand.Intn(2)]
		case RandomKindUpperLower:
			indexList := [2]int{1, 2}
			index = indexList[rand.Intn(2)]
		case RandomKindLower:
			index = 1
		case RandomKindUpper:
			index = 2
		case RandomKindAll:
			index = rand.Intn(3)
		}
		scope, base := kinds[index][0], kinds[index][1]
		randomID[i] = uint8(base + rand.Intn(scope))
	}
	return string(randomID)
}
