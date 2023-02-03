/**
 * formatter print
 */
package format

import "fmt"

/**
 * print green string
 * @params str
 */
func PrintGreen(str string) {
	printColor(str, 32)
}

/**
 * print red string
 * @params str
 */
func PrintRed(str string) {
	printColor(str, 31)
}

/**
 * print color string
 * @params str
 * @params color
 */
func printColor(str string, color int32) {
	fmt.Printf("%c[0;0;%vm%s%c[0m\n", 0x1B, color, str, 0x1B)
}
