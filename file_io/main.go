package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// os_read_single, 单次读取指定长度的文件内容；
// 参数：
// 		os_path: 需要读取内容文件的路径；
// 返回值：
// 		n: 成功读取内容的长度；
//		content: 读取文件的内容；
func os_read_single(os_path string) (n int, content []byte) {

	// 打开文件
	f, err := os.Open(os_path)

	// 打开文件失败
	if err != nil {
		fmt.Println("open file failed!, err: ", err)
		return 0, nil
	}

	// 注册关闭文件的事件
	defer f.Close()

	// 创建一个指定长度的切片
	content = make([]byte, 128)

	// 读取文件
	n, err = f.Read(content)

	// 文件读取完毕
	if err == io.EOF {
		fmt.Println("read over!")
		return
	}

	// 读取文件错误
	if err != nil {
		fmt.Println("read file failed!, err: ", err)
		return 0, nil
	}

	// 返回读取结果
	return
}

// os_read_range, 循环读取文件内容；
// 参数：
// 		os_path: 需要读取内容文件的路径；
// 返回值：
// 		n_count: 成功读取内容的长度；
//		content_all: 读取文件的内容；
func os_read_range(os_path string) (n_count int, content_all []byte) {

	// 打开文件
	f, err := os.Open(os_path)

	// 打开文件失败
	if err != nil {
		fmt.Println("open file failed!, err: ", err)
		return 0, nil
	}

	// 关闭文件
	defer f.Close()

	// 统计读取字符数量变量初始化
	n_count = 0

	// 创建一个用于存储读取内容的切片
	content_all = make([]byte, 0)

	// 创建一个每次读取指定长度的切片
	tmp := make([]byte, 128)

	// 每次读取128个字节的内容，循环读取直到全部读取完毕
	for {
		
		// 读取文件128个字节的内容
		n, err := f.Read(tmp)

		// 判断是否读完了, 如果读完了则跳出循环
		if err == io.EOF {
			fmt.Println("read over!")
			break
		}

		// 读取文件失败
		if err != nil {
			fmt.Println("read file failed!, err: ", err)
			return 0, nil
		}

		// 统计读取的文件内容数量
		n_count = n_count + n

		// 存储每次读取的文件内容
		content_all = append(content_all, tmp[:n]...)

	}

	// 返回读取结果
	return
}

// bufio_read, 使用bufio包按行读取文件内容；
// 参数：
// 		os_path: 需要读取内容文件的路径；
// 返回值：
// 		n_count: 成功读取内容的长度；
//		content_all: 读取文件的内容；
func bufio_read(os_path string) *[]string {

	var content_all []string

	// 打开文件
	f, err := os.Open(os_path)

	// 判断文件是否打开成功
	if err != nil {
		fmt.Println("file open failed! err:", err)
		return nil
	}
	// 关闭打开的文件
	defer f.Close()
	//读取文件信息
	reader := bufio.NewReader(f)
	for {
		// 按行读取数据
		line, err := reader.ReadString('\n')

		// 文件读完了
		if err == io.EOF {
			fmt.Println("read over!")
			break
		}

		// 读取文件失败
		if err != nil {
			fmt.Println("read file failed! err: ", err)
			return nil
		}

		// 打印读取的内容
		content_all = append(content_all, line)
	}
	return &content_all
}

func ioutil_read(os_path string) *string {
	content, err := ioutil.ReadFile(os_path)
	if err != nil {
		fmt.Println("file open failed! err:", err)
		return nil
	}

	contents := string(content)
	return &contents
}

func os_write(os_path string) (is_write bool) {

	// 需要写入的内容
	string_w := "这是一段测试文件写入的内容！"

	// 打开写入内容的目标文件
	file, err := os.OpenFile(os_path, os.O_CREATE|os.O_WRONLY, 0666)

	// 打开文件失败
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return false
	}

	// 注册关闭文件事件
	defer file.Close()

	// 按字节写入文件
	file.Write([]byte(string_w))

	// 按照字符串写入
	file.WriteString(string_w)

	return true
}

func main() {

	path_r := "src/study_go.org/file_io/main.go"
	path_w := "src/study_go.org/file_io/write_test.txt"

	n1, content_os_read_single := os_read_single(path_r)
	n2, content_os_read_range := os_read_range(path_r)
	fmt.Println(n1, n2)
	fmt.Printf("%T %T\n", content_os_read_single, content_os_read_range)

	content_bufio_read := bufio_read(path_r)
	fmt.Printf("%T\n", *content_bufio_read)

	content_ioutil_read := ioutil_read(path_r)
	fmt.Printf("%T\n", *content_ioutil_read)

	is_write := os_write(path_w)

	if is_write {
		fmt.Println("文件写入成功")
	} else {
		fmt.Println("文件写入失败")
	}
}
