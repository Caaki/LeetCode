package main

import "fmt"

func main() {

	fmt.Print(getWordsInLongestSubsequence([]string{"bc", "cbb", "cc", "bbc", "bdc", "ac", "dc", "add", "db", "cb", "cd", "cdb", "bb"}, []int{13, 8, 1, 8, 5, 6, 3, 10, 5, 11, 7, 9, 12}))

}

func getWordsInLongestSubsequence(words []string, groups []int) []string {

	dp := make([][]string, len(words))

	n := len(groups)

	for i, v := range words {
		dp[i] = []string{v}
	}

	mx := 1
	result := dp[0]

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {

			if groups[i] != groups[j] && len(words[i]) == len(words[j]) && humming(words[i], words[j]) &&
				len(dp[i]) < len(dp[j])+1 {
				temp := make([]string, len(dp[j])+1)
				copy(temp, dp[j])
				temp[len(dp[j])] = words[i]
				dp[i] = temp
				if len(dp[i]) > mx {
					mx = len(dp[i])
					result = dp[i]
				}
			}
		}
	}

	return result
}

func humming(a, b string) bool {

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
			if count > 1 {
				return false
			}
		}
	}

	return count == 1
}
