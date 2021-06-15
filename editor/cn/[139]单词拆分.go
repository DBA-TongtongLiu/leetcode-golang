//给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
//
// 说明：
//
//
// 拆分时可以重复使用字典中的单词。
// 你可以假设字典中没有重复的单词。
//
//
// 示例 1：
//
// 输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
//
//
// 示例 2：
//
// 输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//     注意你可以重复使用字典中的单词。
//
//
// 示例 3：
//
// 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false
//
// Related Topics 动态规划
// 👍 1013 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	status := make([]bool, len(s)+1)
	status[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if status[j] && wordMap[s[j:i]] {
				status[i] = true
				break
			}
		}
	}
	return status[len(s)]
}

//leetcode submit region end(Prohibit modification and deletion)
