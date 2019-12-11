package main

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
	解法3
	动态规划中，我们常常可以对空间复杂度进行进一步的优化。
	例如这道题中，可以看到，max_left [ i ] 和 max_right [ i ] 数组中的元素我们其实只用一次，然后就再也不会用到了。所以我们可以不用数组，只用一个元素就行了。我们先改造下 max_left。
	时间复杂度：O(n)
	空间复杂度：O(n)用来保存每一列左边最高的墙和右边最高的墙。
**/
func trap3(height []int) int {
	sum := 0
	max_left := 0
	nlen := len(height)
	max_right := make(map[int]int)
	for i := nlen - 2; i > 0; i-- {
		max_right[i] = max(max_right[i+1], height[i+1])
	}
	for i := 1; i < nlen-1; i++ {
		max_left = max(max_left, height[i-1])
		mi := min(max_left, max_right[i])
		if mi > height[i] {
			sum = sum + (mi - height[i])
		}
	}
	return sum
}

/*
	我们成功将 max_left 数组去掉了。但是会发现我们不能同时把 max_right 的数组去掉，因为最后的 for 循环是从左到右遍历的，而 max_right 的更新是从右向左的。

	所以这里要用到两个指针，left 和 right，从两个方向去遍历。

	那么什么时候从左到右，什么时候从右到左呢？根据下边的代码的更新规则，我们可以知道

	max_left = max(max_left, height[i - 1]);
	height [ left - 1] 是可能成为 max_left 的变量， 同理，height [ right + 1 ] 是可能成为 right_max 的变量。

	只要保证 height [ left - 1 ] < height [ right + 1 ] ，那么 max_left 就一定小于 max_right。

	因为 max_left 是由 height [ left - 1] 更新过来的，而 height [ left - 1 ] 是小于 height [ right + 1] 的，而 height [ right + 1 ] 会更新 max_right，
	所以间接的得出 max_left 一定小于 max_right。

	反之，我们就从右到左更。

	时间复杂度： O(n)
	空间复杂度： O(1)
*/

func trap(height []int) int {
	sum := 0
	max_left := 0
	nlen := len(height)
	max_right := 0
	left := 1
	right := nlen - 2 //  加右指针进去

	for i := 1; i < nlen-1; i++ {
		//从左到右更
		if height[left-1] < height[right+1] {
			max_left = max(max_left, height[left-1])
			mi := max_left
			if mi > height[left] {
				sum = sum + (mi - height[left])
			}
			left++
		} else {
			//从右到左更
			max_right = max(max_right, height[right+1])
			mi := max_right
			if mi > height[right] {
				sum = sum + (mi - height[right])
			}
			right--
		}
	}
	return sum
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
	解法2 动态规划
	我们注意到，解法1中。对于每一列，我们求它左边最高的墙和右边最高的墙，都是重新遍历一遍所有高度，这里我们可以优化一下。

	首先用两个数组，max_left [i] 代表第 i 列左边最高的墙的高度，max_right[i] 代表第 i 列右边最高的墙的高度。（一定要注意下，第 i 列左（右）边最高的墙，是不包括自身的，和 leetcode 上边的讲的有些不同）

	对于 max_left我们其实可以这样求。

	max_left [i] = Max(max_left [i-1],height[i-1])。它前边的墙的左边的最高高度和它前边的墙的高度选一个较大的，就是当前列左边最高的墙了。

	对于 max_right我们可以这样求。

	max_right[i] = Max(max_right[i+1],height[i+1]) 。它后边的墙的右边的最高高度和它后边的墙的高度选一个较大的，就是当前列右边最高的墙了。

	这样，我们再利用解法二的算法，就不用在 for 循环里每次重新遍历一次求 max_left 和 max_right 了。

	时间复杂度：O(n)。
	空间复杂度：O(n)，用来保存每一列左边最高的墙和右边最高的墙。

**/
func trap2(height []int) int {
	sum := 0
	nlen := len(height)
	max_left := make(map[int]int)
	max_right := make(map[int]int)
	// 最两端的列不用考虑，因为一定不会有水。所以下标从1到nlen-2
	for i := 1; i < nlen-1; i++ {
		max_left[i] = max(max_left[i-1], height[i-1])
	}
	for i := nlen - 2; i > 0; i-- {
		max_right[i] = max(max_right[i+1], height[i+1])
	}
	for i := 1; i < nlen-1; i++ {
		mi := min(max_left[i], max_right[i])
		if mi > height[i] {
			sum = sum + (mi - height[i])
		}
	}
	return sum
}

/**
	方法1

	求每一列的水，我们只需要关注当前列，以及左边最高的墙，右边最高的墙就够了。
	装水的多少，当然根据木桶效应，我们只需要看左边最高的墙和右边最高的墙中较矮的一个就够了。
	所以，根据较矮的那个墙和当前列的墙的高度可以分为三种情况。
	较矮的墙的高度大于当前列的墙的高度

		1.较矮的墙的高度大于当前列的墙的高度
			5|
			 |
			4|
			 |											 			_____
			3|										   			|   |
			 |					 _____  		 			|   |
			2|           |   |            |   |
			 | 	 _____   |   |___    	____|   |
			1|   |   |   |   |   |	  |   |   |
			 |___|___|___|___|___|____|___|___|___|___|___|___|_____________
			0    1   2   4 | 5 | 6   7   8  | 9   10  11  12
										 |	 |						|
										 |	 正在求的列    |
								左边最高的墙				右边最高的墙
			 把正在求的列左边最高的墙和右边最高的墙确定后,然后为了方便理解我们把无关的墙都关了

			5|
			 |
			4|
			 |											 			_____
			3|										   			|   |
			 |					 _____  		 			|   |
			2|           |   |            |   |
			 | 	 			   |   |___    			|   |
			1|           |   |   |	      |   |
			 |___|___|___|___|___|________|___|___|___|___|___|_____________
			0    1   2   4 | 5 | 6   7   8  | 9   10  11  12
										 |	 |						|
										 |	 正在求的列    |
								左边最高的墙				右边最高的墙
			这样就很清楚了，现在想想一下，往两边最高的墙之间注水，正在求的列会有多少水？
			很明显，较矮的一边，也就是左边的墙的高度，减去当前列的高度就可以了，也就是 2 - 1 = 1，可以存一个单位的水。


		2.较矮的墙的高度小于于当前列的墙的高度
			5|
			 |
			4|
			 |											 			_____
			3|										   			|   |
			 |					 _____  		 			|   |
			2|           |   |            |   |
			 | 	 _____   |   |___    	____|   |
			1|   |   |   |   |   |	  |   |   |
			 |___|___|___|___|___|____|___|___|___|___|___|___|_____________
			0    1 |  2    3    4   5  6   7   8  | 9   10  11  12
						 |			 |						    |
						 |			正在求的列         |
					 左边最高的墙				   右边最高的墙

					 同样的，我们把其他无关的列去掉。
			5|
			 |
			4|
			 |											 			_____
			3|										   			|   |
			 |					 _____  		 			|   |
			2|           |   |            |   |
			 | 	 _____   |   |      	    |   |
			1|   |   |   |   |    	      |   |
			 |___|___|___|___|________|___|___|___|___|___|___|_____________
			0    1 |  2    3    4   5  6   7   8  | 9   10  11  12
						 |			 |						    |
						 |			正在求的列         |
					 左边最高的墙				   右边最高的墙
			想象下，往两边最高的墙之间注水。正在求的列会有多少水？
			正在求的列不会有水，因为它大于了两边较矮的墙。

		3.较矮的墙的高度等于当前列的墙的高度。和上一种情况是一样的，不会有水。
			5|
			 |
			4|
			 |											 			_____
			3|										   			|   |
			 |	 _____	 _____  		 			|   |
			2|   |   |   |   |            |   |
			 | 	 |   |   |   |            |   |
			1|   |   |   |   |    	      |   |
			 |___|___|___|___|________|___|___|___|___|___|___|_____________
			0    1 |  2    3    4   5  6   7   8  | 9   10  11  12
						 |			 |						    |
						 |			正在求的列         |
					 左边最高的墙				   右边最高的墙
		明白了这三种情况，程序就很好写了，遍历每一列，然后分别求出这一列两边最高的墙。找出较矮的一端，和当前列的高度比较，结果就是上边的三种情况。

		时间复杂度：O(n²）O(n²），遍历每一列需要 nn，找出左边最高和右边最高的墙加起来刚好又是一个 n，所以是 n²
		空间复杂度：O(1）O(1）。

**/
func trap1(height []int) int {
	sum := 0
	nlen := len(height)
	// 最两端的列不用考虑，因为一定不会有水。所以下标从1到nlen-2
	for i := 1; i < nlen-1; i++ {
		max_left := 0
		for j := i - 1; j >= 0; j-- {
			if height[j] > max_left {
				max_left = height[j]
			}
		}
		max_right := 0

		for j := i + 1; j < nlen; j++ {
			if height[j] > max_right {
				max_right = height[j]
			}
		}
		mi := min(max_left, max_right)
		if mi > height[i] {
			sum += (mi - height[i])
		}
	}
	return sum
}

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
			5|
			 |
			4|
			 |											 			_____
			3|										   			|   |
			 |					 _____  		 			|   |___    _____
			2|           |   |            |   |	  |   |		|
			 | 	 _____   |   |___    	____|   |   |___|	  |____
			1|   |   |   |   |   |	  |   |   |   |   |   |   |
			 |___|___|___|___|___|____|___|___|___|___|___|___|_____________
			0    1   2   4 | 5   6   7   8   9   10  11  12
										 |
上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

*/
