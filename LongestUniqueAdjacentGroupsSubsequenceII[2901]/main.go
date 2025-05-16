package main

func main() {}

func getWordsInLongestSubsequence(words []string, groups []int) []string {

	dp := make([][]string, len(words))

	n := len(groups)

	for i, v := range words {
		dp[i] = []string{v}
	}

	mx := 1
	result := dp[0]

	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {

			if groups[i] != groups[j] && len(words[i]) == len(words[j]) && humming(words[i], words[j]) && len(dp[j]) < len(dp[i])+1 {
				dp[j] = append(dp[i], words[j])
				if len(dp[j]) > mx {
					mx = len(dp[j])
					result = dp[j]
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
