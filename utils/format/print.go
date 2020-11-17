/**
 * 格式化-输出
 */
package format

import "fmt"

/**
 * 输出绿色字符串
 * @params str	待输出的字符串
 */
func PrintGreen(str string) {
	printColor(str, 32)
}

/**
 * 输出红色字符串
 * @params str	待输出的字符串
 */
func PrintRed(str string) {
	printColor(str, 31)
}

/**
 * 输出待颜色字符串
 * @params str	待输出的字符串
 * @params color 待输出颜色
 */
func printColor(str string, color int32) {
	fmt.Printf("%c[0;0;%vm%s%c[0m\n", 0x1B, color, str, 0x1B)
}
