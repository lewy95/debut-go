package grammar

import "fmt"

func TestSwitch() {

	grade := "A"
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	default:
		fmt.Println("一般")
	}

	// 支持多条件判断
	day := 3
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	}

	// 支持表达式判断
	score := 90
	switch {
	case score >= 90:
		fmt.Println("享受假期")
	case score < 90 && score >= 80:
		fmt.Println("好好学习吧！")
	default:
		fmt.Println("玩命学习！")
	}
}

func TestSwitchFallthrough() {
	value := 100
	switch value {
	case 100:
		fmt.Println("100")
		// break 倘若加上break，则 fallthrough 不会执行，结果为 100
		fallthrough // 使用fallthrough可以执行满足条件的下一个case，结果为 100 200
	case 200:
		fmt.Println("200")
	case 300:
		fmt.Println("300")
	default:
		fmt.Println("other")
	}
}
