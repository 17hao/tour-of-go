package main

import (
	"sync"

	"github.com/sirupsen/logrus"
)

// func f2() {
// 	str := "eyJjb2RlIjowLCJtc2ciOiJzdWNjZXNzIiwiZGF0YSI6eyJyZXZpc2lvbiI6NjcxODgsInNwcmVhZHNoZWV0VG9rZW4iOiJzaHRjbnFtbWliZEVXeWJ0VEswWFg2TzZGUGQiLCJ2YWx1ZVJhbmdlIjp7Im1ham9yRGltZW5zaW9uIjoiUk9XUyIsInJhbmdlIjoiY2UwN2RkIUExOkJaMSIsInJldmlzaW9uIjo2NzE4OCwidmFsdWVzIjpbWyLnrb7nuqbkuLvkvZMiLCLnlLPor7fml6XmnJ8iLCLnlLPor7fkuroiLCLkuovkuJrpg6giLCLpg6jpl6giLCLlkIjlkIzmgKfotKgiLCLlkIjlkIzlvaLlvI8iLCLlkIjkvZzmlrnnsbvlnosiLCLpobnnm67lj7ciLCLlr7nmlrnlhazlj7jlkI3np7AiLCLlkIjlkIzph5Hpop0iLCLlvIDlp4vml6XmnJ8iLCLnu5PmnZ/ml6XmnJ8iLCLpo57kuabljZXmja7nvJblj7ciLCJDT1VOVElGKE46TixOMSkiLCLnlKjnq6DnsbvlnosiLCLmmK/lkKbnm5bnq6AiLCLlkIjlkIzlvZLmoaPnvJblj7ciLCLlvZLmoaPml6XmnJ8iLCLljbDoirHnqI7miYDlsZ7mnJ8iLCLlpIfms6giLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsXV19fX0="
// 	b, _ := base64.StdEncoding.DecodeString(str)
// 	fmt.Printf("%s", string(b))
// }

func f1(minID, maxID int64) {
	batchSize := 10
	x := 32
	gap := (maxID - minID + 1) / int64(x)

	consts := []int64{6, 4, 4, 4, 3, 3, 3, 2, 2, 1}
	batch := make([]int64, 10)
	batch[0] = minID
	for i := 1; i < batchSize; i++ {
		batch[i] = batch[i-1] + gap*consts[i-1]
	}

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < batchSize; i++ {
		num := i
		min := batch[num]
		max := maxID
		if i < batchSize-1 {
			max = batch[num+1] - 1
		}
		go func() {
			defer wg.Done()
			if min > max {
				return
			}
			logrus.Infof("num=%d min=%d, max=%d batchSize=%d", num+1, min, max, max-min)
		}()
	}
	wg.Wait()
}

func main() {
	f1(7281956966799327252, 7281956966799327290)
}
