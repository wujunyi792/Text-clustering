package clustering

import (
	"clustering/config"
	"clustering/myDiff"
	"regexp"
	"strconv"
	"strings"
)

func TextArrClustering(testStrArr []string) {

	clusteringConfig := config.ClusteringConfig{}
	clusteringConfig.GetConf()
	transformedArr := [][]string{}
	for _, str := range testStrArr {
		transformedArr = append(transformedArr, []string{str})
	}

	quitFlag := true
	for quitFlag {
		quitFlag = false
		i := 0
		for i < len(transformedArr) {
			j := 0
			for j < len(transformedArr) {
				if i != j {
					if len(transformedArr[j]) == 1 && myCompare(transformedArr[i], transformedArr[j][0], clusteringConfig) > clusteringConfig.LegitimateMatchingRate {
						transformedArr[i] = append(transformedArr[i], transformedArr[j][0])
						//transformedArr[j]
						transformedArr = append(transformedArr[:j], transformedArr[j+1:]...)
						quitFlag = true
					}
				}
				j++
			}
			i++

		}
	}
	//for _, strArr := range transFromedArr {
	//	fmt.Println(strArr)
	//}
	//fmt.Println("从", len(testStrArr), "个缩减到了", len(transFromedArr), "个")
}

// 比较两个字符串数组
func myCompare(strArr1 []string, str2 string, clusteringConfig config.ClusteringConfig) float64 {
	var tempStrArr1 []string
	tempStr2 := str2
	//先将strArr中的badWords去掉
	for _, str := range strArr1 {
		for _, badWord := range clusteringConfig.BadWords {
			str = strings.Replace(str, badWord, "", -1)
		}
		tempStrArr1 = append(tempStrArr1, str)
	}
	for _, badWord := range clusteringConfig.BadWords {
		tempStr2 = strings.Replace(tempStr2, badWord, "", -1)
	}
	//再将strArr中的reBadWords通过正则匹配的方式去掉
	tempReStr := ""
	for _, reBadWord := range clusteringConfig.ReBadWords {
		//匹配Start和end中间最多有reBadWord.MaxLength个字符的字符串
		tempReStr += reBadWord.Start + ".{1," + strconv.Itoa(reBadWord.MaxLength) + "}" + reBadWord.End + "|"
	}
	re := regexp.MustCompile(tempReStr)

	for _, str := range tempStrArr1 {
		tempStrArr1 = append(tempStrArr1, re.ReplaceAllString(str, ""))
	}
	tempStr2 = re.ReplaceAllString(tempStr2, "")
	//计算出平均匹配度
	var sum float64
	for _, str := range tempStrArr1 {
		matcher := myDiff.SequenceMatcher{A: strings.Split(str, ""), B: strings.Split(tempStr2, "")}
		sum += matcher.Ratio()
	}
	return sum / float64(len(tempStrArr1))
}
