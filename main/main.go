package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"

	//"go_code/project02/model"
	"strconv"
	"unsafe"
)

var (
	Fun1 = func(n1, n2 int) int {
		return n1 * n2
	}
)

func test(char byte) byte {
	return char + 1
}

func main() {
	var i int
	i = 10
	fmt.Println("i=", i)

	a := 11.1
	fmt.Println("a=", a)

	x, y, z := 100, "ysr", 13.5
	fmt.Println("x =", x, "y =", y, "z =", z)
	var n2 int64 = 10
	fmt.Printf("n2的数据类型是%T sssd%d", n2, unsafe.Sizeof(n2))
	//推荐使用float64用于开发
	var n3 int = '北'
	fmt.Printf("n3=%c", n3)
	//使用反引号 ``输出大段字符
	str3 := `var i int
  i = 10
  fmt.Println("i=", i)

  a := 11.1
  fmt.Println("a=", a)

  x, y, z := 100, "ysr", 13.5
  fmt.Println("x =", x, "y =", y, "z =", z)
  var n2 int64 = 10
  fmt.Printf("n2的数据类型是%T sssd%d", n2, unsafe.Sizeof(n2))
  //推荐使用float64用于开发
  var n3 int = '北'
  fmt.Printf("n3=%c", n3) `
	fmt.Println(str3)

	var num3 int = 100
	var num4 float64 = 32.132
	var str string

	//其他类型转为str类型
	str = fmt.Sprintf("%d", num3)
	fmt.Printf("str%Ttype,str=%q\n", str, str)

	str = fmt.Sprintf("%f", num4)
	fmt.Printf("str%Ttype,str=%q\n", str, str)

	str = strconv.FormatInt(int64(num3), 2) //2代表将num3变成2进制
	fmt.Printf("str%Ttype,str=%q\n", str, str)
	str = strconv.FormatFloat(num4, 'f', 10, 64) //f代表格式，10小数点后10位，64代表为float64
	fmt.Printf("str%Ttype,str=%q\n", str, str)

	// str转为其他类型
	var str1 string = "true"
	var b bool
	//strconv会返回两个值 bool ，error ，只取bool，error用_忽略
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("%T,%v\n", b, b) //%v相应值的默认格式。

	var str2 string = "12345"
	var num5 int64
	num5, _ = strconv.ParseInt(str2, 10, 64) //10进制，64位
	fmt.Printf("%T,%v\n", num5, num5)

	var str4 string = "123,456"
	var num6 float64
	num6, _ = strconv.ParseFloat(str4, 64)
	fmt.Printf("%T,%v\n", num6, num6)
	//变量首字母大写为公有public，小写为私有private

	//fmt.Println(model.Num1)
	var nums1 float64 = 10 / 4
	fmt.Printf("%v\n", nums1)
	var nums2 float64 = 10 / 4.0
	fmt.Printf("%v", nums2)
	//++ --只能独立使用
	var a1 int = 4
	var b1 int
	a1++
	b1 = a1
	fmt.Printf("%v", b1)
	//只能i++，不能++i

	var name string
	var age byte
	var sal float64
	//第一种输入方式
	fmt.Scanln(&name)
	fmt.Scanln(&age)
	fmt.Scanln(&sal)
	//第二种输入方式
	fmt.Scanf("%s %d %f", &name, &age, &sal)

	if age := 18; age > 11 {
		fmt.Printf("%v", age)
	} else {

	}

	var key byte
	fmt.Scanf("%c", &key)

	switch test(key) {
	case 'a':

	case 'b':

	case 'c':

	default:

	}

	var str11 string = "123,123123"
	for index, val := range str11 {
		fmt.Printf("%d,%c", index, val)
	}
	var str12 string = "sads1312"
	str13 := []rune(str12)
	for i := 0; i < len(str13); i++ {
		fmt.Printf("%c", str13[i])
	}

	r := 1
	for {
		if r > 10 {

			break
		}
		r++
	}

	//这里演示break的指定标签的用法
label2:
	for i := 0; i < 4; i++ {
		//label1: //设置一个标签
		for j := 0; j < 3; j++ {
			if j == 2 {
				//break label1
				break label2
			}
			fmt.Printf("%v", j)
		}
	}

	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(200))

	//goto的使用
	fmt.Println("1")
	goto label1
	fmt.Println("2")
label1:
	fmt.Println("3")

	//return 表示退出函数

	//函数可以有多个返回值，一个文件包含全局变量，init函数以及main函数执行顺序按照变量-init-main

	//匿名函数
	res := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(res)

	Res := func(n1, n2 int) int {
		return n1 + n2
	}

	res2 := Res(20, 30)
	fmt.Println(res2)

	//全局匿名函数
	res3 := Fun1(4, 9)
	fmt.Println(res3)

	//字符串常用函数 中文字符占3个字节
	//一般使用rune【】处理含有中文的字符串遍历
	//strconv.Atoi()字符串转整数
	//strconv.Itoa()整数转字符串

	//fmt.Printf("%d-%d-%d %d:%d:%d")格式化时间

	time.Sleep(time.Duration(time.Hour.Milliseconds() * 100))

	// unix时间戳返回从1970.1.1到现在的时间秒数

	//通过这个来算出运行时间
	start := time.Now().Unix()
	end := time.Now().Unix()
	fmt.Println(start - end)

	num2 := new(int) //num2的类型是*int
	fmt.Println(num2)

	//错误处理 defer panic recover go先抛出panic错误再通过defer和recover来处理错误

	//使用defer加recover处理错误
	test3()
	for {
		fmt.Println("")
		time.Sleep(time.Second)
	}

	// 数组

	var arr3 = []int{1: 100, 0: 200, 3: 400}
	fmt.Println(arr3)

	var intarr3 [3]int
	//通过改变seed的值来让rand产生的数变化
	rand.Seed(time.Now().Unix())
	for i := 0; i < 3; i++ {
		intarr3[i] = rand.Intn(100)
	}

	var intarr4 [4]int = [4]int{1, 2, 3, 4}
	slice := intarr4[1:3] //从1到3不包含3
	fmt.Println(slice)
	fmt.Println(cap(slice)) //通过cap获取动态变化slice的容量
	// 修改slice会改变原数组 slice是引用  slice从底层来说是一个结构体

	//通过make创建slice
	var slice2 = make([]int, 4, 10)
	slice2[1] = 100

	//直接定义slice
	var strslice []string = []string{"tom", "ss"}
	fmt.Println(strslice)

	var arr5 [5]int = [5]int{10, 20, 30, 40, 50}
	slice3 := arr5[0:len(arr5)]
	fmt.Println(slice3)

	slice3 = append(slice3, slice...)
	fmt.Println(slice3)

	//string 是不可变的不能通过str【0】=‘z'来改变值
	//要修改string字符串中的字符要通过将其转为byte【】来处理数字和字符
	//而其中的中文字符通过rune【】来处理

	//for-range来遍历二维数组
	var arrint = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i, v := range arrint {
		for j, v2 := range v {
			fmt.Println(i, j, v2)
		}
	}

	// map属于key-value类型,map的key-value无序
	var a11 map[string]string
	a11 = make(map[string]string, 10)
	a11["ss"] = "ss"
	fmt.Println(a11)

	city := make(map[string]string)
	city["1"] = "ws"
	city["2"] = "sw"
	city["3"] = "ss"
	fmt.Println(city)

	heros := map[string]string{
		"h1": "sss",
		"h2": "www",
	}
	fmt.Println(heros)

	//map的查找
	val, ok := city["1"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("no find")
	}

	//动态增加map信息
	var monster []map[string]string
	monster = make([]map[string]string, 2)
	newMonster := map[string]string{
		"name": "ss",
		"age":  "123",
	}
	monster = append(monster, newMonster)
	//排序map
	keys := make(map[int]int, 10)
	keys[1] = 200
	keys[3] = 400
	keys[5] = 300

	var key1 []int
	for k, _ := range keys {
		key1 = append(key1, k)
	}
	sort.Ints(key1)

	for _, k := range key1 {
		fmt.Println(k, keys[k])
	}
	// map是引用类型
	//使用slice和map一定要先make
	var p1 person
	if p1.ptr == nil {
	}
	p1.slice = make([]int, 2)
	p1.map1 = make(map[string]string)

	var p3 *person = new(person)
	(*p3).name = "sssss"

	var person1 *person = &person{}
	(*person1).name = "12312"
	person1.name = "1312312 "

	//结构体之间可以相互转化，但是要求结构体字段完全相同

	monsters := monster1{"s12", 12}
	jsonstr, err := json.Marshal(monsters)
	if err != nil {
	}
	fmt.Println(jsonstr)

	//调用方法 test方法只能person来调用
	var p11 person
	p11.test()

	stu := Student{
		Name: "ts",
		Age:  11,
	}
	fmt.Println(&stu)

	pupil := &pupil{}
	pupil.Student.Name = "sss"
	pupil.Student.Age = 12
	pupil.Student.show()
}

//修改结构体变量的值最好通过结构体指针来改变

func (stu *Student) String() string {
	str := fmt.Sprintf(stu.Name, stu.Age)
	return str
}
func (stu *Student) show() {
	fmt.Println(stu.Name, stu.Age)
}

type Student struct {
	Name string
	Age  byte
}
type pupil struct {
	Student //嵌入结构体，通过改结构体来使用student中的方法以及变量无论大小写
}

func (p person) test() {
	fmt.Println(p.age)
}

type person struct {
	ptr   *int
	slice []int
	map1  map[string]string
	name  string
	age   byte
}

type monster1 struct {
	Name string `json:"name"` //如果想输出小写name
	Age  byte   `json:"age"`
}

func test3() {
	defer func() {
		if err := recover(); err != nil { //也可以 err := recover() if err != nil
			fmt.Println(err)
			//这里就可以加错误信息发送给管理员
			fmt.Println("错误信息")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

// 函数去读取以配置文件init。conf的信息
// 如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取
		return nil

	} else {
		return errors.New("读取文件错误")
	}
}

func test02() {
	err := readConf("config.ini")
	if err != nil {
		//读取文件发送错误，输出这个错误并终止程序
		panic(err)
	}
	fmt.Println("test02继续执行")
}
