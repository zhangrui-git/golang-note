## Golang中的隐藏特性

### 一、初级篇

#### 1. 开大括号不能放在单独的一行

>go:5:6: missing function body
>go:6:1: syntax error: unexpected semicolon or newline before {

  ```go
  package main
  
  import "fmt"
  
  func main()
  {
  	fmt.Println("abc")
  }
  ```

---

> 以下示例不会报错

  ```go
  package main
  
  import "fmt"
  
  func main()  {fmt.Println("abc")}
  ```

  ```go
  package main
  
  import "fmt"
  
  func main()  {fmt.Println("abc")
  }
  ```

  ```go
  package main
  
  import "fmt"
  
  func main()  {
  	fmt.Println("abc")}
  ```

#### 2. 未使用的变量

> 全局变量不会报错

> go:5:2: v2 declared but not used
>
> go:6:2: v3 declared but not used
>
> go:7:6: v4 declared but not used

```go
package main

var v1 = 1
func main() {
	v2 := 2
	v3 := 3
  v4 := 4
}
```

---

> 以下示例不会报错

```go
package main

var v1 = 1
func main() {
	v2 := 2
	_ = v2
	v3 := 3
	_ = v3
	var v4 = 1
	_ = v4
}
```

#### 3. 未使用的import

> go:4:2: imported and not used: "fmt"

```go
package main

import "fmt"

func main() {
	
}
```

#### 4. 简式的变量声明仅可以在函数内部使用

> go:5:1: syntax error: non-declaration statement outside function body

```go
package main

import "fmt"

v1 := 1
//var v1 = 1
func main() {
	v2 := 2
	fmt.Println(v1)
	fmt.Println(v2)
}
```

#### 5. 使用简式声明重复声明变量

> go:7:5: no new variables on left side of :=

```go
package main

import "fmt"

func main() {
	v1 := 1
	v1 := 2
	fmt.Println(v1)
}
```

---

> 以下示例不会报错

```go
package main

import "fmt"

func main() {
	var v1 = 1
	v1 = 2
	fmt.Println(v1)
}
```

#### 6. 偶然的变量隐藏

```go
package main

import "fmt"

func main() {
	var v1 = 1
	fmt.Println(v1)
	{
		fmt.Println(v1)
		v1 = 2
		fmt.Println(v1)
	}
	fmt.Println(v1)
}
```

> > 1
> > 1
> > 2
> > 2

---

```go
package main

import "fmt"

func main() {
	var v1 = 1
	fmt.Println(v1)
	{
		fmt.Println(v1)
		v1 := 2
		fmt.Println(v1)
	}
	fmt.Println(v1)
}
```

> > 1
> > 1
> > 2
> > 1

---

```go
package main

import "fmt"

func main() {
	v1 := 1
	fmt.Println(v1)
	{
		fmt.Println(v1)
		v1 := 2
		fmt.Println(v1)
	}
	fmt.Println(v1)
}
```

> > 1
> > 1
> > 2
> > 1

---

```go
package main

import "fmt"

func main() {
   v1 := 1
   fmt.Println(v1)
   {
      fmt.Println(v1)
      v1 = 2
      fmt.Println(v1)
   }
   fmt.Println(v1)
}
```

> > 1
> > 1
> > 2
> > 2

#### 7. 不使用显式类型，无法使用“nil”来初始化变量

> go:4:6: use of untyped nil

```go
package main

func main() {
	var i = nil
	_ = i
}
```

---

> go:4:6: cannot use nil as type int in assignment

```go
package main

func main() {
	var i int = nil
	_ = i
}
```

---

```go
package main

func main() {
	var i interface{} = nil
	_ = i
}
```

#### 8. 使用“nil” Slices and Maps

> slice可以不用初始化操作

```go
package main

func main() {
	var s []int
	s = append(i, 1,2)
}
```

---

> panic: assignment to entry in nil map

```go
package main

func main() {
	var m map[string]int
	m["a"] = 1
}
```

---

```go
package main

func main() {
	var i = make(map[string]int)
	i["a"] = 1
}
```

#### 9. Map的容量

> go:8:17: invalid argument m (type map[string]int) for cap

```go
package main

import "fmt"

func main() {
	var m = make(map[string]int, 99)
	m["a"] = 1
	fmt.Println(cap(m))
}
```

---

```go
package main

import "fmt"

func main() {
	var m = make(map[string]int, 99)
	m["a"] = 1
	fmt.Println(len(m))
}
```

> > 1

#### 10. Slice的容量

```go
package main

import "fmt"

func main() {
	s := make([]int, 20)
	fmt.Println(len(s))
}
```

> > 20

---

```go
package main

import "fmt"

func main() {
	var s = make([]int, 20)
	fmt.Println(len(s))
}
```

> > 20

---

```go
package main

import "fmt"

func main() {
   var s []int
   fmt.Println(len(s))
}
```

> > 0

---

```go
package main

import "fmt"

func main() {
   var s = make([]int, 20)
   fmt.Println(cap(s))
}
```

> > 20

---

```go
package main

import "fmt"

func main() {
	s := make([]int, 20)
	fmt.Println(cap(s))
}
```

> > 20

---

> cap() 可以作用在array、slice、channel
>
> len() 可以作用在array、slice、channel

#### 11. 字符串

> go:6:6: cannot use nil as type string in assignment

```go
package main

import "fmt"

func main() {
	var s string = nil
	fmt.Println(s)
}
```

---

```go
package main

import "fmt"

func main() {
	str := "123abc中国"
	for i, i2 := range str {
		fmt.Println(i, i2)
	}
}
```

> 字节索引---utf8编码
>
> > 0 49
> > 1 50
> > 2 51
> > 3 97
> > 4 98
> > 5 99
> > 6 20013
> > 9 22269

---

```go
fmt.Println(string(22269))
```

> 国

#### 12. 数组

> 数组的参数是值传递

```go
package main

import "fmt"

func main() {
	var arr = [3]int{1,2,3}
	fmt.Println(arr)
	func(arr1 [3]int) {
		arr1[0] = 4
		arr1[1] = 5
		arr1[2] = 6
		fmt.Println(arr1)
	}(arr)
	fmt.Println(arr)
}
```

> > [1 2 3]
> > [4 5 6]
> > [1 2 3]

---

```go
package main

import "fmt"

func main() {
	var arr = [3]int{1,2,3}
	fmt.Println(arr)
	func(arr1 *[3]int) {
		arr1[0] = 4
		arr1[1] = 5
		arr1[2] = 6
		fmt.Println(arr1)
	}(&arr)
	fmt.Println(arr)
}
```

> > [1 2 3]
> > &[4 5 6]
> > [4 5 6]

#### 13. 闭包

```go
package main

import "fmt"

func getFun(i int) func() int {
	if (1 & i) > 0 {
		return func() int {
			i++; i++
			return i
		}
	} else {
		return func() int {
			i++
			return i
		}
	}
}
func main() {
	i := 1
	fmt.Println(getFun(i)())
	i2 := 2
	fmt.Println(getFun(i2)())
	i3 := 3
	fmt.Println(getFun(i3)())
	i4 := 4
	fmt.Println(getFun(i4)())
}
```

> > 3
> > 3
> > 5
> > 5

#### 14. 不存在的map key

```go
package main

import "fmt"

func main() {
	m := map[string]string{"one": "1", "two": "", "three": "3"}
	if v := m["two"]; v == "" {
		fmt.Println("empty")
	}
	v1, v2 := m["two"]
	fmt.Println(v1, v2)
	v3, v4 := m["three"]
	fmt.Println(v3, v4)
	if _, ok := m["four"]; !ok {
		fmt.Println("empty")
	}
}
```

> > empty
> >    true
> > 3 true
> > empty

#### 15. string修改

> go:5:9: cannot assign to str[0]

```go
package main

func main() {
	str := "big"
	str[0] = "p"
}
```

---

> 正确做法

```go
package main

import "fmt"

func main() {
	str := "big"
	//str[0] = 'p'
	byt := []byte(str)
	byt[0] = 'p'
	fmt.Println(byt)
	fmt.Println(string(byt))
}
```

> > [112 105 103]
> > pig

#### 16. 字符串不总是UTF8文本

> 为了知道字符串是否是UTF8，你可以使用“unicode/utf8”包中的 `ValidString()`函数

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "z中\xfec"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))

	fmt.Println(utf8.ValidString(str))
}
```

> > 6
> > 4
> > false

#### 17. 字符串的长度

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "z中"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
}
```

> > 4
> > 2

#### 18. 在多行的Slice、Array和Map语句中遗漏逗号

> go:11:4: syntax error: unexpected newline, expecting comma or }

```go
package main

import (
	"fmt"
)

func main() {
	s := []int{
		1,
		2,
		3
	}
	fmt.Println(s)
}
```

正确语法

```go
package main

import (
	"fmt"
)

func main() {
	s := []int{
		1,
		2,
		3,
	}
	fmt.Println(s)
  s1 := []int{1, 2, 3,}
	s2 := []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Println(s2)
}

```

#### 19. log.Fatal和log.Panic不仅仅是Log

> Go中log包在你调用它的 `Fatal*()`和 `Panic*()`函数时，可以做的不仅仅是log。当你的应用调用这些函数时，Go也将会终止应用。

 ```go
package main

import "log"

func main() {
	log.Fatalln("")
	log.Panicln("")
}
 ```

#### 20. switch声明中的行为

> switch语句中的case默认不会break进入下一个next。

```go
package main

import "fmt"

func main() {
	isSpace := func(c byte) bool {
		switch c {
		case '\t':
		case ' ':return true
		}
		return false
	}
	fmt.Println(isSpace('\t'))
	fmt.Println(isSpace(' '))
}
```

> > false
> > true

正确语法

```go
isSpace2 := func(c byte) bool {
  switch c {
    case '\t', ' ':return true
    }
  return false
}
fmt.Println(isSpace2('\t'))
fmt.Println(isSpace2(' '))
```

> > true
> > true

#### 21. 自增自减

> go:8:20: syntax error: unexpected ++, expecting :

```go
package main

import "fmt"

func main() {
	data := [3]int{1, 2, 3}
	i := 0
  //i := 0++
  //i := 0--
	fmt.Println(data[i++])
}
```

正确语法

```go
package main

import "fmt"

func main() {
	data := [3]int{1, 2, 3}
	i := 0
	i++
	fmt.Println(data[i])
}
```

> > 2

#### 22. 未导出的结构体不会被编码

> 以小写字母开头的结构体将不会被（json、xml、gob等）编码，因此当你编码这些未导出的结构体时，你将会得到零值。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type MyData struct {
	One int
	two int
}

type MyData2 struct {
	One int
	Two int
}

func main() {
	md := MyData{1, 2}
	jsonStr, _ := json.Marshal(md)
	fmt.Println(string(jsonStr))

	md2 := MyData2{1, 2}
	jsonStr2, _ := json.Marshal(md2)
	fmt.Println(string(jsonStr2))
}
```

> > {"One":1}
> > {"One":1,"Two":2}

#### 23. 异或

```go
package main

import "fmt"

func main() {
	var a uint = 254
	fmt.Printf("%b\n", a)
	var b uint = 128
	fmt.Printf("%b\n", b)
	fmt.Printf("%b\n", a ^ b)
	fmt.Printf("%b\n", a &^ b)
}
```

> > 11111110
> > 10000000
> > 01111110
> > 01111110

#### 24. goroutines的退出

> 主goroutines不会等待所有子goroutines结束。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	workerCount := 2
	for i := 1; i < workerCount; i++ {
		go func(workerId int) {
			fmt.Printf("[%v] running\n", workerId)
			time.Sleep(3 * time.Second)
			fmt.Printf("[%v] done\n", workerId)
		}(i)
	}
	time.Sleep(3 * time.Second)
	fmt.Printf("all done")
}
```

> > [1] running
> > [1] done
> > all done

正确方法

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	workerCount := 2
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func(workerId int) {
			fmt.Printf("[%v] running\n", workerId)
			time.Sleep(3 * time.Second)
			fmt.Printf("[%v] done\n", workerId)
			defer wg.Done()
		}(i)
	}
	//time.Sleep(3 * time.Second)
	wg.Wait()
	fmt.Printf("all done")
}
```

> > [1] running
> > [0] running
> > [0] done
> > [1] done
> > all done

---

#### 25. WaitGroup产生死锁

> 当`Add()`与`Done()`不匹配时会出现死锁。

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg  = sync.WaitGroup{}

func main() {
	workerCount := 2
	wg.Add(workerCount)
	// for loop 1
	for i := 1; i < workerCount; i++ {
		go func(workerId int, wg sync.WaitGroup) {
			fmt.Printf("[%v] running\n", workerId)
			time.Sleep(3 * time.Second)
			fmt.Printf("[%v] done\n", workerId)
			defer wg.Done()
		}(i, wg)
	}
	//time.Sleep(3 * time.Second)
	wg.Wait()
	fmt.Printf("all done")
}
```

> > [1] running
> > [1] done
> > fatal error: all goroutines are asleep - deadlock!

---

> 当WaitGroup通过参数传递时，其实是值传递。所以会导致死锁。

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	workerCount := 2
	var wg sync.WaitGroup
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go worker(i, wg)
	}
	wg.Wait()
	fmt.Printf("all done")
}

func worker(workerId int, wg sync.WaitGroup)  {
	fmt.Printf("[%v] running\n", workerId)
	time.Sleep(3 * time.Second)
	fmt.Printf("[%v] done\n", workerId)
	defer wg.Done()
}
```

> > [1] running
> > [0] running
> > [0] done
> > [1] done
> > fatal error: all goroutines are asleep - deadlock!

正确方法

> 使用引用传递。

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	workerCount := 2
	var wg sync.WaitGroup
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Printf("all done")
}

func worker(workerId int, wg *sync.WaitGroup)  {
	fmt.Printf("[%v] running\n", workerId)
	time.Sleep(3 * time.Second)
	fmt.Printf("[%v] done\n", workerId)
	defer wg.Done()
}
```

> > [1] running
> > [0] running
> > [0] done
> > [1] done
> > all done

#### 26. Channel通道

> 只要目标接收就会立即返回。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		for m := range ch {
			time.Sleep(1 * time.Second)
			fmt.Println("processed:", m)
		}
	}()
	ch<- "cmd.1"
	ch<- "cmd.2"
	ch<- "cmd.3"
}
```

> > processed: cmd.1
> > processed: cmd.2
> > ~~processed: cmd.3~~

---

> 向已关闭的Channel发送会引起Panic。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		for m := range ch {
			time.Sleep(1 * time.Second)
			fmt.Println("processed:", m)
		}
	}()
	ch<- "cmd.1"
	close(ch)
	ch<- "cmd.2"
	ch<- "cmd.3"
}
```

> > panic: send on closed channel

---

> 不能向nil的Channel发送和接受。

```go
package main

import (
	"fmt"
)

func main() {
	var ch chan string
	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
		}
	}()
	ch<- "cmd.1"
	ch<- "cmd.2"
	ch<- "cmd.3"
}
```

> > fatal error: all goroutines are asleep - deadlock!



### 二、进阶篇

#### 1. 关闭http的响应

> 当使用标准http库发起请求时，会得到一个http响应，不管是否使用都需要关闭。
>
> 对于成功请求关闭是没有问题，但如果http请求失败，变量可能是nil，关闭时会导致一个runtime panic。

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.baidu2.com/")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
```

> > panic: runtime error: invalid memory address or nil pointer dereference
> > [signal SIGSEGV: segmentation violation code=0x1 addr=0x40 pc=0x123dac4]

正确方法

> `defer resp.Body.Close()`放到错误检查之后。

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.baidu2.com/")
	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
```

#### 2. 关闭HTTP的连接

> 一些HTTP服务器保持会保持一段时间的网络连接（根据HTTP 1.1的说明和服务器端的“keep-alive”配置）。默认情况下，标准http库只在目标HTTP服务器要求关闭时才会关闭网络连接。这意味着你的应用在某些条件下消耗完sockets/file的描述符。你可以通过设置请求变量中的 `Close`域的值为 `true`，来让http库在请求完成时关闭连接。另一个选项是添加一个 `Connection`的请求头，并设置为`close`。目标HTTP服务器应该也会响应一个 `Connection: close`的头。当http库看到这个响应头时，它也将会关闭连接。

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, err := http.NewRequest("get", "https://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
  //-------------------------------------
	req.Close = true
	//req.Header.Add("Connection", "close")
  //-------------------------------------
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
```

> 你也可以取消http的全局连接复用。你将需要为此创建一个自定义的http传输配置。

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	tr :=&http.Transport{DisableKeepAlives:true}
	client :=&http.Client{Transport: tr}

	resp, err := client.Get("http://golang.org")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(len(string(body)))
}
```

#### 3. 比较Structs, Arrays, Slices, and Maps

> 以下情况时可比较的。

```go
package main

import"fmt"

type data struct{
	num int
	fp float32
	complex complex64
	str string
	char rune
	yes bool
	events <-chan string
	handler interface{}
	ref *byte
	raw [10]byte}

func main(){
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:",v1 == v2)
}
```

> > v1 == v2: true

---

> 以下情况时不可比较的。

```go
package main

import"fmt"

type data struct{
	num int//ok
	checks [10]func()bool//not comparable
	doit func()bool//not comparable
	m map[string]string//not comparable
	bytes []byte//not comparable
}

func main(){
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:",v1 == v2)
}
```

> > go:16:29: invalid operation: v1 == v2 (struct containing [10]func() bool cannot be compared)
> >
> > go:16:29: invalid operation: v1 == v2 (struct containing func() bool cannot be compared)
> >
> > go:16:29: invalid operation: v1 == v2 (struct containing map[string]string cannot be compared)
> >
> > go:16:29: invalid operation: v1 == v2 (struct containing []byte cannot be compared)

---

> Go提供了一些助手函数，用于比较那些无法使用等号比较的变量。最常用的方法是使用 `reflect`包中的 `DeepEqual()`函数。

```go
package main

import (
	"fmt"
	"reflect"
)

func main(){
	v1 := [3]int{1, 2, 3}
	v2 := [3]int{1, 2, 3}
	v3 := [3]int{2, 3, 4}
	v4 := [2]int{1, 2}
	fmt.Println(reflect.DeepEqual(v1, v2))
	fmt.Println(reflect.DeepEqual(v2, v3))
	fmt.Println(reflect.DeepEqual(v3, v4))
}
```

> > true
> > false
> > false

> 除了很慢（这个可能会也可能不会影响你的应用）， `DeepEqual()`也有其他自身的技巧。
>
> `DeepEqual()`不会认为空的slice与“nil”的slice相等。这个行为与你使用 `bytes.Equal()`函数的行为不同。 `bytes.Equal()`认为“nil”和空的slice是相等的。
>
> `DeepEqual()`在比较slice时并不总是完美的。
>
> 如果你的byte slice（或者字符串）中包含文字数据，而当你要不区分大小写形式的值时（在使用 `==`， `bytes.Equal()`，或者 `bytes.Compare()`），你可能会尝试使用“bytes”和“string”包中的 `ToUpper()`或者 `ToLower()`函数。对于英语文本，这么做是没问题的，但对于许多其他的语言来说就不行了。这时应该使用 `strings.EqualFold()`和 `bytes.EqualFold()`。
>
> 如果你的byte slice中包含需要验证用户数据的隐私信息（比如，加密哈希、tokens等），不要使用 `reflect.DeepEqual()`、 `bytes.Equal()`，或者 `bytes.Compare()`，因为这些函数将会让你的应用易于被定时攻击。为了避免泄露时间信息，使用 `'crypto/subtle'`包中的函数（即， `subtle.ConstantTimeCompare()`）。

#### 4. 从

> `recover()`函数可以用于获取/拦截panic。仅当在一个defer函数中被完成时，调用 `recover()`将会完成这个小技巧。

```go
package main

import "fmt"

func main(){
	recover()
	fmt.Println("recover1")
	panic("not good")
	//recover()
	//fmt.Println("recover2")
}
```

正确做法

```go
package main

import "fmt"

func main(){
	defer func() {
		recover()
		fmt.Println("recover")
	}()
	panic("not good")
}
```

> > recover

#### 5. 在Slice, Array, and Map "range"语句中更新引用元素的值

```go
package main

import "fmt"

func main(){
	s := []int{1, 2, 3}
	for _, i2 := range s {
		i2 *= 10
	}
	fmt.Println(s)
  // 使用索引是有效的更新
	for i, _ := range s {
		s[i] *= 10
	}
	fmt.Println(s)
}
```

> > [1 2 3]
> > [10 20 30]

```go
package main

import "fmt"

func main(){
	s := []*struct{num int}{{1}, {2}, {3}}
	for _, i2 := range s {
		i2.num *= 10
	}
	fmt.Println(s[0], s[1], s[2])
}
```

> > &{10} &{20} &{30}

#### 6. 在Slice中"隐藏"数据

> 容易造成陷阱。slice会拷贝数组的引用。

```go
package main

import "fmt"

func main(){
	data := get()
	fmt.Println(data)
	fmt.Printf("%d\n", &data[0])
	fmt.Printf("%d\n", &data[1])
}

func get() []int {
	s := make([]int, 3)
	s[0] = 10
	s[1] = 20
	s[2] = 30
	fmt.Printf("%d\n", &s[0])
	fmt.Printf("%d\n", &s[1])
	return s[:2]
}
```

> > 824633819360
> > 824633819368
> > [10 20]
> > 824633819360
> > 824633819368

正确做法

```go
package main

import "fmt"

func main(){
	data := get()
	fmt.Println(data)
	fmt.Printf("%d\n", &data[0])
	fmt.Printf("%d\n", &data[1])
}

func get() []int {
	s := make([]int, 3)
	s[0] = 10
	s[1] = 20
	s[2] = 30
	fmt.Printf("%d\n", &s[0])
	fmt.Printf("%d\n", &s[1])
	s1 := make([]int, 2)
	copy(s1, s[:2])
	return s1
}
```

> > 824633819360
> > 824633819368
> > [10 20]
> > 824633811040
> > 824633811048

> 结果并不是想象中预料的。

```go
package main

import "fmt"

func main(){
	s := []int{1, 2, 3}
	s2 := s[:2]

	for i, _ := range s2 {
		s2[i] *= 10
	}
	fmt.Println(s)
	fmt.Println(s2)
}
```

> > [10 20 3]
> > [10 20]

#### 7. 新的类型声明

> `type myUser user` 使用原始类型重新定义一个类型，但你无法使用原始类型方法。

```go
package main

import "fmt"

type user struct {
	name string
}

func (u *user) getName() string {
	return u.name
}

type myUser user

func main(){
	var u myUser
	fmt.Println(u.getName())
}
```

> > go:17:15: u.getName undefined (type myUser has no field or method getName)

正确做法

> `type myUser struct {user}` 重新定义一个struct，匿名嵌入原始类型。

```go
package main

import "fmt"

type user struct {
	name string
}

func (u *user) getName() string {
	return u.name
}

type myUser struct {
	user
}

func main(){
	var u myUser
	fmt.Println(u.getName())
}
```

#### 8. for switch、for select、for for代码块跳出

> for switch 使用label加break跳出。

```go
package main

import "time"

func main(){
	loop:for {
		switch {
		case true:
			println("break")
			break loop
		}
		time.Sleep(3 * time.Second)
	}
	println("end")
}
```

> > break
> > end

---

> for switch 使用label加goto跳出。

```go
package main

import "time"

func main(){
	for {
		switch {
		case true:
			println("break")
			goto endLoop
		}
		time.Sleep(3 * time.Second)
	}
	endLoop:
	println("end")
}
```

---

> 使用bool变量加if break跳出。

```go
package main

import "time"

func main(){
	endFor := false
	for {
		switch {
		case true:
			println("break")
			endFor = true
		}
		if endFor {
			break
		}
		time.Sleep(3 * time.Second)
	}

	println("end")
}
```

#### 9. for中迭代的变量和闭包

> 运行以下代码会发现结果并不是预料的。

```go
package main

import (
	"fmt"
	"time"
)

func main(){
	s := []string{"one", "two", "three"}
	for _, v := range s {
		fmt.Println("for:" + v)
		go func() {
			fmt.Println("goroutines:" + v)
		}()
	}

	time.Sleep(3 * time.Second)
}
```

> > for:one
> > for:two
> > for:three
> > goroutines:three
> > goroutines:three
> > goroutines:three

通过加一行赋值语句

```go
package main

import (
	"fmt"
	"time"
)

func main(){
	s := []string{"one", "two", "three"}
	for _, v := range s {
		fmt.Println("for:" + v)
		v := v
		go func() {
			fmt.Println("goroutines:" + v)
		}()
	}

	time.Sleep(3 * time.Second)
}
```

> > for:one
> > for:two
> > for:three
> > goroutines:three
> > goroutines:one
> > goroutines:two

通过闭包参数传递

```go
package main

import (
	"fmt"
	"time"
)

func main(){
	s := []string{"one", "two", "three"}
	for _, v := range s {
		fmt.Println("for:" + v)
		go func(v string) {
			fmt.Println("goroutines:" + v)
		}(v)
	}

	time.Sleep(3 * time.Second)
}
```

> > for:one
> > for:two
> > for:three
> > goroutines:three
> > goroutines:two
> > goroutines:one

---

下面是个陷阱例子

```go
package main

import (
	"fmt"
	"time"
)

type user struct {
	name string
}

func (u *user)Print() {
	fmt.Println(u.name)
}

func main(){
	s := []user{{"one"}, {"two"}, {"three"}}
	for _, v := range s {
		go v.Print()
	}

	time.Sleep(3 * time.Second)
}
```

> > three
> > three
> > three

通过加一行赋值语句

```go
package main

import (
	"fmt"
	"time"
)

type user struct {
	name string
}

func (u *user)Print() {
	fmt.Println(u.name)
}

func main(){
	s := []user{{"one"}, {"two"}, {"three"}}
	for _, v := range s {
		v := v
		go v.Print()
	}

	time.Sleep(3 * time.Second)
}
```

> > three
> > two
> > one

slice初始化为struct的引用

```go
package main

import (
	"fmt"
	"time"
)

type user struct {
	name string
}

func (u *user)Print() {
	fmt.Println(u.name)
}

func main(){
	s := []*user{{"one"}, {"two"}, {"three"}}
	for _, v := range s {
		go v.Print()
	}

	time.Sleep(3 * time.Second)
}
```

> > one
> > three
> > two

#### 10. 被defer的函数参数

> defer在声明时就已经确定参数值，而不是在实际执行时。

```go
package main

import "fmt"

func main(){
	var i = 10
	defer func(v int) {
		fmt.Println("defer print:", v)
	}(i)
	i++
	fmt.Println(i)
}
```

> > 11
> > defer print: 10

#### 11. 被defer的函数执行

> defer的函数是在被包含的函数末尾执行，而不是在被包含的代码块末尾执行。

```go
package main

import (
	"fmt"
	"time"
)

type box struct {
	name string
}

func (b *box)Open() bool {
	fmt.Println(b.name, " box open")
	return true
}

func (b *box)Close()  {
	fmt.Println(b.name, " box close")
}

func main(){
	b := []*box{{"one"}, {"two"}, {"three"}}
	for _, v := range b {
		if v.Open() {
			defer v.Close()
		}
		time.Sleep(3 * time.Second)
	}
	fmt.Println("end")
}
```

> > one  box open
> > two  box open
> > three  box open
> > end
> > three  box close
> > two  box close
> > one  box close

#### 12. 失败的类型断言

> 在使用类型断言时使用了断言结果覆盖了变量的值。

```go
package main

import "fmt"

func main(){
	var data interface{} = "great"
	if data, ok := data.(int); ok {
		fmt.Println("is int")
		fmt.Println(data)
	} else {
		fmt.Println("not int")
		fmt.Println(data)
	}
}
```

> > not int
> > 0

在实际应用中不因该覆盖变量

```go
	if res, ok := data.(int); ok {
		fmt.Println("is int")
		fmt.Println(res)
	} else if res, ok := data.(string); ok {
		fmt.Println("is string")
		fmt.Println(res)
	}
```

> > is string
> > great

### 三、高级篇

#### 1. 使用指针接收方法的值的实例

> 并不是所有的变量是可取址的。

```go
package main

import "fmt"

type user struct {
	name string
}

func (u *user)print()  {
	fmt.Println(u.name)
}

type printer interface {
	print()
}

func main(){
	u := user{"rui"}
	u.print()

	var p printer = &user{name: "zhang"}
	p.print()

	//m := map[string]user{"u": {name:"zhangrui"}}
	//m["u"].print()
}
```

> > rui
> > zhang

> > go:25:8: cannot call pointer method on m["u"]
> > go:25:8: cannot take the address of m["u"]

正确做法

```go
package main

import "fmt"

type user struct {
	name string
}

func (u *user)print()  {
	fmt.Println(u.name)
}

type printer interface {
	print()
}

func main(){
	u := user{"rui"}
	u.print()

	//var p printer = user{name: "zhang"}
	var p printer = &user{name: "zhang"}
	p.print()

	//m := map[string]user{"u": {name:"zhangrui"}}
	m := map[string]*user{"u": {name:"zhangrui"}}
	m["u"].print()
}
```

> > rui
> > zhang
> > zhangrui

#### 2. 更新Map的值

> struct值的map，你无法更新单个的struct值，map元素是无法取址的。
>
> slice元素是可以取址的。

```go
package main

import "fmt"

type user struct {
	name string
}

func (u *user)print()  {
	fmt.Println(u.name)
}

func main(){
	u := user{"rui"}
	u.print()

	m := map[string]user{"u": {name:"zhangrui"}}
	m["u"].name = "zhang"
}
```

> > go:18:14: cannot assign to struct field m["u"].name in map

```go
package main

import "fmt"

type user struct {
	name string
}

func (u *user)print()  {
	fmt.Println(u.name)
}

func main(){
	u := user{"rui"}
	u.print()

	//m := map[string]user{"u": {name:"zhangrui"}}
	//m["u"].name = "zhang"
	s := []user{{name:"zhangrui"}}
	s[0].name = "zhang"
	fmt.Println(s)
}
```

> > rui
> > [{zhang}]

> 使用struct引用的map是可以的。当访问一个不存在的key时会出现异常。

```go
package main

import "fmt"

type user struct {
	name string
}

func (u *user)print()  {
	fmt.Println(u.name)
}

func main(){
	m := map[string]*user{"u": {"zhangrui"}}
	m["u"].name = "rui"
	m["u"].print()
  //--------------
	m["u2"].print()
}
```

> > rui
> > panic: runtime error: invalid memory address or nil pointer dereference

#### 3. "nil" Interface和"nil" Interface的值

```go
package main

import "fmt"

func main(){
	var data *byte
	var in interface{}
	fmt.Println(data == nil)
	fmt.Println(in == nil)
	in = data
	fmt.Println(in == nil)
}
```

> > true
> > true
> > false

#### 4. 栈和堆变量

> 你并不总是知道变量是分配到栈还是堆上。在C++中，使用 `new`创建的变量总是在堆上。在Go中，即使是使用 `new()`或者 `make()`函数来分配，变量的位置还是由编译器决定。编译器根据变量的大小和“泄露分析”的结果来决定其位置。这也意味着在局部变量上返回引用是没问题的，而这在C或者C++这样的语言中是不行的。
>
> 如果你想知道变量分配的位置，在“go build”或“go run”上传入“-m“ gc标志（即， `go run -gcflags -m app.go`）。

#### 5. GOMAXPROCS并发和并行

> 默认情况下，Go仅使用一个执行上下文/OS线程（在当前的版本）。这个数量可以通过设置 `GOMAXPROCS`来提高。
>
> 一个常见的误解是， `GOMAXPROCS`表示了CPU的数量，Go将使用这个数量来运行goroutine。而 `runtime.GOMAXPROCS()`函数的文档让人更加的迷茫。 `GOMAXPROCS`变量描述（[https://golang.org/pkg/runtime/](https://links.jianshu.com/go?to=https%3A%2F%2Fgolang.org%2Fpkg%2Fruntime%2F)）所讨论OS线程的内容比较好。
>
> 你可以设置 `GOMAXPROCS`的数量大于CPU的数量。 `GOMAXPROCS`的最大值是256。

```go
package main

import(
	"fmt"
	"runtime"
)

func main(){
	fmt.Println(runtime.GOMAXPROCS(-1))
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(20)
	fmt.Println(runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(300)
  // 这个和网络文档上的并不一样
  // 当前机器
  // MacBook Pro (13-inch, 2017, Two Thunderbolt 3 ports)
  // 2.3 GHz 双核Intel Core i5 
  // 8 GB 2133 MHz LPDDR3
	fmt.Println(runtime.GOMAXPROCS(-1))
}
```

> > 4
> > 4
> > 20
> > 300

#### 6. 读写操作的重排顺序

> Go可能会对某些操作进行重新排序，但它能保证在一个goroutine内的所有行为顺序是不变的。然而，它并不保证多goroutine的执行顺序。
>
> 如果你需要在多goroutine内放置读写顺序的变化，你将需要使用channel，或者使用"sync"包构建合适的结构体。

```go
package main

import(
	"runtime"
	"time"
	)
var _ = runtime.GOMAXPROCS(3)
var a, b int

func u1(){
	a =1
	b =2
}

func u2(){
	a =3
	b =4
}

func p(){
	println(a)
	println(b)
}

func main(){
	go u1()
	go u2()
	go p()
	time.Sleep(1 * time.Second)
}
```

> > 1
> > 4

> > 3
> > 4

> > 0
> > 0

#### 7. 优先调度

> 有可能会出现这种情况，一个无耻的goroutine阻止其他goroutine运行。当你有一个不让调度器运行的 `for`循环时，这就会发生。
>
> `for`循环并不需要是空的。只要它包含了不会触发调度执行的代码，就会发生这种问题。
>
> 调度器会在GC、“go”声明、阻塞channel操作、阻塞系统调用和lock操作后运行。它也会在非内联函数调用后执行。
>
> 要想知道你在 `for`循环中调用的函数是否是内联的，你可以在“go build”或“go run”时传入“-m” gc标志（如， `go build -gcflags -m`）。
>
> 另一个选择是显式的唤起调度器。你可以使用“runtime”包中的 `Goshed()`函数。