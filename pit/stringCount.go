package pit

/*
在 golang 中不能直接用 len 函数来统计字符串长度，查看了下源码发现字符串是以 UTF-8 为格式存储的，说明 len 函数是取得包含 byte 的个数

*/

//func main() {
//
//	s := "Hello, 世界"
//	fmt.Println(len(s))    // 13
//	fmt.Println([]byte(s)) // [72 101 108 108 111 44 32 228 184 150 231 149 140]
//
//	//fmt.Print(f1(s))
//}

//func f1(s string) int {
//	return bytes.Count([]byte(s), nil) - 1
//}
//
//func f2(s string) int {
//	return strings.Count(s, "") - 1
//}
//
//func f3(s string) int {
//	return len([]rune(s))
//}
//
//func f4(s string) int {
//	return utf8.RuneCountInString(s)
//}
//
//var s = "Hello, 世界"
//
//func Benchmark1(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		f1(s)
//	}
//}
//
//func Benchmark2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		f2(s)
//	}
//}
//
//func Benchmark3(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		f3(s)
//	}
//}
//
//func Benchmark4(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		f4(s)
//	}
//}
