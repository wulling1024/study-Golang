package main

// 这一个利用接口来实现的排序算法
// 排序中主要就是三个函数：Len()、Less()、Sort()

// 接口类型
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Sort(i, j int)
}

// 对 int 数组进行排序
type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}

func (p IntArray) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p IntArray) Sort(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// 冒泡排序
func sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i, pass) {
				data.Sort(i, pass)
			}
		}
	}
}

//func main() {
//	arr := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
//	param := IntArray(arr)
//
//	sort(param)
//
//	fmt.Printf("The sorted array is %v\n", param)
//}
