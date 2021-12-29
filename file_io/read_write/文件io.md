# Go语言的文件读取与写入
* 本文主要介绍了使用Go的 `os`/`bufio`/`ioutil`三个go内置包对文件进行读取操作；

## 文件打开和关闭
* 使用`os.Open()`方法可以打开一个文件,使用该方法打开一个文件后会返回一个`*File` 和 一个`error`;
* 使用`f.Close()`方法可以关闭一个文件，通常为了防止忘记关闭文件，我们都使用`defer`来注册文件关闭语句；

代码示例

```go
func main(){
	
	// 需要打开文件的地址
	path := "./main.go"

	// 使用只读的方式打开文件
	f, err := os.Open(path)

	// 打开文件失败
	if err != nil {
		fmt.Println("open file failed!, err: ", err)
		return
	}

	// 关闭文件
	defer f.Close()
}
```


## 文件读取
### 使用 f.Read()读取文件内容

> Read方法定义
```go
func (f, *File) Read(b []byte) (n int, err error) 
```
它接收一个字节切片，返回读取的字节数和可能出现的错误，当读到文件的末尾时会返回`0`和`io.EOF`。

#### 单次读取文件内容

代码示例：
```go
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
```

#### 循环读取文件内容
代码示例：
```go
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
```

### 使用bufio读取文件
```go

```
