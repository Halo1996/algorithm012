package main
//127. 单词接龙
//给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
//
//每次转换只能改变一个字母。
//转换过程中的中间单词必须是字典中的单词。
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if idxOf(endWord, wordList) == -1 {
		return 0
	}
	step := 0
	startQueue := []string{beginWord}
	endQueue := []string{endWord}
	used := make([]bool, len(wordList))

	for len(startQueue) > 0 {
		step++
		l := len(startQueue)
		for i := 0; i < l; i++ {
			for k := 0; k < len(endQueue); k++ {
				if hasOneDiff(startQueue[i], endQueue[k]) {
					return step + 1
				}
			}
			for j, w := range wordList {
				if !used[j] && hasOneDiff(startQueue[i], w) {
					startQueue = append(startQueue, w)
					used[j] = true
				}
			}
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
		}
	}
	return 0
}
func idxOf(str string, bank []string) int {
	for i, s := range bank {
		if str == s {
			return i
		}
	}
	return -1
}

func hasOneDiff(x, y string) bool {
	count := 0
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	if count == 1 {
		return true
	}
	return false
}
func main() {
	
}
