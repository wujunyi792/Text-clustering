package Text_clustering

import (
	"github.com/Lost-little-dinosaur/Text-clustering/config"
	"github.com/Lost-little-dinosaur/Text-clustering/myDiff"
	"regexp"
	"strconv"
	"strings"
)

//	testStrArr := []string{"2022年长三角数学建模竞赛", "浙江省第十九届大学生程序设计竞赛", "大创省级立项", "浙江省计算机设计大赛省一", "2022年第二届长三角数学建模竞赛一等奖", "第二届长三角高校数学建模竞赛 SPSSPRO创新应用奖", "浙江省spsspro创新应用奖", "长三角数学建模竞赛二等奖", "杭电第23届数学建模竞赛", "大学生创业竞赛 省级立项", "浙江省大学生物理（理论）创新竞赛", "第二届长三角高校数学建模竞赛", "第十三届中国大学生服务外包创新创业企业命题类（A类）", "认证杯数学建模竞赛", "高教社杯数学建模竞赛", "2022长三角数学建模竞赛三等奖", "2022年第二届长三角高校数学建模竞赛", "互联网➕", "长三角数学建模", "长三角数学建模竞赛", "2022蓝桥杯浙江省一等奖", "浙江省计算机设计大赛一等奖", "浙江省计算机设计大赛省一", "第十二届浙江省大学生物理实验与科技创新竞赛", "服务外包大赛区域赛三等奖", "长三角数模竞赛", "长三角数学建模竞赛", "全国大学生数学竞赛", "浙江省计算机设计大赛", "2022年浙江省第十九届大学生程序设计竞赛", "第二届长三角高校数学建模竞赛\t"}

func TextArrClustering(testStrArr []string) {

	clusteringConfig := config.ClusteringConfig{}
	clusteringConfig.GetConf()
	transFromedArr := [][]string{}
	for _, str := range testStrArr {
		transFromedArr = append(transFromedArr, []string{str})
	}

	//fmt.Println(transFromedArr)
	quitFlag := true
	for quitFlag {
		quitFlag = false
		i := 0
		for i < len(transFromedArr) {
			j := 0
			for j < len(transFromedArr) {
				if i != j {
					if len(transFromedArr[j]) == 1 && myCompare(transFromedArr[i], transFromedArr[j][0], clusteringConfig) > clusteringConfig.LegitimateMatchingRate {
						transFromedArr[i] = append(transFromedArr[i], transFromedArr[j][0])
						//删除transFromedArr[j]
						transFromedArr = append(transFromedArr[:j], transFromedArr[j+1:]...)
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

//比较两个字符串数组
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
