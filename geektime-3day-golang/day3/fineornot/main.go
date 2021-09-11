package main

import "fmt"

func main() {
	for {
		var weight, height float64
		var age int
		var sexFlag int
		var isMale bool

		fmt.Println("欢迎使用体脂计算器")
		fmt.Println("请输入体重（kg）")
		fmt.Scanln(&weight)
		fmt.Println("身高（m）")
		fmt.Scanln(&height)
		fmt.Println("年龄")
		fmt.Scanln(&age)
		fmt.Println("性别（男 1，女 0）")
		fmt.Scanln(&sexFlag)

		if isMale {
			sexFlag = 1
		} else {
			sexFlag = 0
		}
		//bm=体重(公斤)÷(身高身高)(米)
		var bmi float64 = weight / (height * height)
		//计算体脂
		//体脂率:1.2×BMI+0.23×年龄-5.4-10.8×性别(男为1,女为0)
		var fatRate float64 = 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexFlag)
		// 判别肥胖
		if isMale {
			if age >= 60 {
				dispFatRate(fatRate, judgeFateRate(0, 13, 19, 24, 29, fatRate))
			} else if age >= 49 {
				dispFatRate(fatRate, judgeFateRate(0, 11, 17, 22, 27, fatRate))
			} else if age >= 18 {
				dispFatRate(fatRate, judgeFateRate(0, 10, 16, 21, 26, fatRate))
			} else {
				fmt.Println("小于18岁，不在统计模型中")
			}
		} else {
			if age >= 60 {
				dispFatRate(fatRate, judgeFateRate(0, 22, 29, 36, 41, fatRate))
			} else if age >= 49 {
				dispFatRate(fatRate, judgeFateRate(0, 21, 28, 35, 40, fatRate))
			} else if age >= 18 {
				dispFatRate(fatRate, judgeFateRate(0, 20, 27, 34, 39, fatRate))
			} else {
				fmt.Println("小于18岁，不在统计模型中")
			}
		}
		whetherContinue := "n"
		fmt.Println("是否重新评估？y：继续，n：不继续")
		fmt.Scanln(&whetherContinue)
		if whetherContinue == "n" {
			break
		}
	}
}

/**
根据评估模型
偏瘦
标准
偏重
肥胖
严重肥胖
的体脂率下界，评估
ref:https://www.sohu.com/a/82600953_114873
*/
func judgeFateRate(UnderFatLowerBound,
	StandardMinusLowerBound,
	StandardPlusLowerBound,
	OverfatLowerBound,
	ObesityLowerBound, rate float64) string {
	if rate >= ObesityLowerBound {
		return "严重肥胖"
	} else if rate >= OverfatLowerBound {
		return "肥胖"
	} else if rate >= StandardPlusLowerBound {
		return "偏重"
	} else if rate >= StandardMinusLowerBound {
		return "标准"
	} else if rate >= UnderFatLowerBound {
		return "偏瘦"
	} else {
		return "数据错误，体脂率不能小于0"
	}
}
func dispFatRate(rate float64, judgeRes string) {
	fmt.Println("您的体脂率为：", rate, " 当前体重：", judgeRes)
}
