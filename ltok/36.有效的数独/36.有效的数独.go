package main

import (
	"fmt"
	"strconv"
)

/*
我们也可以对空间进行些优化，只使用一个 HashSet 来记录已经存在过的状态，
将每个状态编码成为一个字符串，能将如此大量信息的状态编码成一个单一的字符串还是需要有些技巧的。
对于每个1到9内的数字来说，其在每行每列和每个小区间内都是唯一的，将数字放在一个括号中，
每行上的数字就将行号放在括号左边，每列上的数字就将列数放在括号右边，
每个小区间内的数字就将在小区间内的行列数分别放在括号的左右两边，
这样每个数字的状态都是独一无二的存在，就可以在 HashSet 中愉快地查找是否有重复存在啦
*/

func isValidSudoku(board [][]byte) bool {
	leni := len(board)
	lenj := 0
	if leni > 0 {
		lenj = len(board[0])
	}
	strMap := make(map[string]bool)
	for i := 0; i < leni; i++ {
		for j := 0; j < lenj; j++ {
			if board[i][j] != '.' {
				t := "(" + string(board[i][j]) + ")"
				row := strconv.Itoa(i) + t
				col := t + strconv.Itoa(j)
				cell := strconv.Itoa(i/3) + t + strconv.Itoa(j/3)
				_, ok := strMap[row]
				_, ok1 := strMap[col]
				_, ok2 := strMap[cell]
				if ok || ok1 || ok2 {
					return false
				}
				strMap[cell] = true
				strMap[col] = true
				strMap[row] = true
			}
		}
	}
	return true
}

func main() {
	arrs := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(len(arrs))
}

/*
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。


上图是一个部分填充的有效的数独。

数独部分空格内已填入了数字，空白格用 '.' 表示。

示例 1:

输入:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: true
示例 2:

输入:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: false
解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
     但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。
说明:

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
给定数独序列只包含数字 1-9 和字符 '.' 。
给定数独永远是 9x9 形式的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-sudoku
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
