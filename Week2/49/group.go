package _9

import "sort"

func groupAnagrams(strs []string) [][]string {
	resultMap := make(map[string][]string)
	for _, str := range strs {
		tmp := []byte(str)
		sort.Slice(tmp, func(i, j int) bool {
			if tmp[i] < tmp[j] {
				return true
			}
			return false
		})
		if _, ok := resultMap[string(tmp)]; ok {
			resultMap[string(tmp)] = append(resultMap[string(tmp)], str)
		} else {
			resultMap[string(tmp)] = []string{str}
		}
	}
	resultArr := make([][]string, len(resultMap))
	i := 0
	for _, strings := range resultMap {
		resultArr[i] = strings
		i++
	}
	return resultArr
}
