// File: time_complexity_types.go
// Created Time: 2022-12-13
// Author: cathay (cathaycchen@gmail.com)

package chapter_computational_complexity

// constant 常数阶
func constant(n int) int {
	count := 0
	var size = 100000
	for i := 0; i < size; i++ {
		count++
	}
	return count
}

// linear 线性阶
func linear(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count++
	}
	return count
}

// arrayTraversal 线性阶（遍历数组）
func arrayTraversal(nums []int) int {
	count := 0
	// 循环次数与数组长度成正比
	for range nums {
		count++
	}
	return count
}

// quadratic 平方阶
func quadratic(n int) int {
	count := 0
	// 循环次数与数组长度成平方关系
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			count++
		}
	}
	return count
}

// bubbleSort 平方阶（冒泡排序）
func bubbleSort(nums []int) int {
	count := 0  // 计数器
	// 外循环：待排序元素数量为 n-1, n-2, ..., 1
	for i := len(nums) - 1; i > 0; i-- {
		// 内循环：冒泡操作
		for j := 0; j < i; j++ {
			if nums[j] > nums[j + 1] {
				// 交换 nums[j] 与 nums[j + 1]
				tmp := nums[j]
				nums[j] = nums[j + 1]
				nums[j + 1] = tmp
				count += 3 	// 元素交换包含 3 个单元操作
			}
		}
	}
	return count
}

// exponential 指数阶（循环实现）
func exponential(n int) int {
	count := 0
	base := 1
	// cell 每轮一分为二，形成数列 1, 2, 4, 8, ..., 2^(n-1)
	for i := 0; i < n; i++ {
		for j := 0; j < base; j++ {
			count++
		}
		base *= 2
	}
	// count = 1 + 2 + 4 + 8 + .. + 2^(n-1) = 2^n - 1
	return count
}

// expRecur 指数阶（递归实现）
func expRecur(n int) int {
	if n == 1 {
		return 1
	}
	return expRecur(n - 1) + expRecur(n - 1) + 1
}

// logarithmic 对数阶（循环实现）
func logarithmic(n float32) int {
	count := 0
	for n > 1 {
		n = n / 2
		count++
	}
	return count
}

// logRecur 对数阶（递归实现）
func logRecur(n float32) int {
	if n <= 1 {
		return 0
	}
	return logRecur(n / 2) + 1
}

// 线性对数阶
func linearLogRecur(n float32) int {
	if n <= 1 {
		return 1
	}
	count := linearLogRecur(n / 2) + linearLogRecur(n / 2)
	for i := 0; float32(i) < n; i++ {
		count++
	}
	return count
}

// factorialRecur 阶乘阶（递归实现）
func factorialRecur(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	// 从 1 个分裂出 n 个
	for i := 0; i < n; i++ {
		count += factorialRecur(n - 1)
	}
	return count
}


