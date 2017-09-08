package main

/*
获取并解析命令行输入；
从对应文件中读取输入数据；
调用对应的排序函数；
将排序的结果输出到对应的文件中；
打印排序所花费时间的信息。
 */
import (
	"flag" //Go语言标准库提供了用于快迅解析命令行参数的flag包
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
	"time"
	"sorter/algorithms/qsort"
	"sorter/algorithms/bubblesort"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")
//读取文件
func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile) //打开文件失败
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int,0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected")
		}
		str := string(line) // 转换字符数组为字符串
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

//写文件
func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println()
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}
func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		fmt.Println("Read values:", values)
		t1 := time.Now()//当前时间
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")//运行时间
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}
