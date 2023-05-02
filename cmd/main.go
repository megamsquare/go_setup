package main

// "log"
// "net/http"

func main() {
	// http.HandleFunc("/", handler)
	// fmt.Println("Starting Server at port :3000")
	// log.Fatal(http.ListenAndServe(":3000", nil))
	// nums1 := []int{4, 9, 5}
	// nums2 := []int{4, 1, 2, 1, 2}
	// fmt.Println(intersect(nums1, nums2))

	// ch := make(chan int)
	// go send(ch)
	// go receive(ch)
	// time.Sleep(time.Second)

}

// func send(ch chan<- int) {
//     ch <- 1
// }

// func receive(ch <-chan int) {
//     val := <-ch
//     fmt.Println(val)
// }

// func intersect(nums1 []int, nums2 []int) []int {
// 	result := []int{}

// 	freq := make(map[int]int)
// 	for _, num := range nums1 {
// 		freq[num]++
// 	}
// 	for _, num := range nums2 {
// 		if freq[num] > 0 {
// 			result = append(result, num)
// 			freq[num]--
// 		}
// 	}

// 	return result

// }

// func singleNumber(nums []int) int {
// 	checker := make(map[int]int)
// 	result := 0

// 	for _, num := range nums {
// 		_, ok := checker[num]
// 		if ok {
// 			checker[num]++
// 		} else {
// 			checker[num]++
// 		}
// 		if checker[num] > 1 && checker[num] == result {
// 			// fmt.Println(checker)
// 			result = 0
// 		} else {
// 			result = num
// 		}
// 	}

// 	fmt.Println(result)
// 	return result
// }

// func rotate(nums []int, k int) {
// 	if len(nums) == 0 {
// 		return
// 	}
// 	k = k % len(nums)
// 	if k == 0 {
// 		return
// 	}
// 	reverse(nums, 0, len(nums)-1)
// 	reverse(nums, 0, k-1)
// 	reverse(nums, k, len(nums)-1)

// 	fmt.Println(nums)
// }

// func reverse(nums []int, start int, end int) {
// 	for start < end {
// 		nums[start], nums[end] = nums[end], nums[start]
// 		start++
// 		end--
// 	}
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Starting Server")
// }
