> 阅读了《程序员的职业素养》之后，对书中提倡的TDD（测试驱动开发）产生好奇，遂进行一些练习，记录于此。  

google一下找到了很好的练习文档：https://github.com/quii/learn-go-with-tests  

开启从零实践TDD之路。
### 理解TDD
虽然Bob 大叔在书中5.2中提到了TDD的三个法则，但我并不打算码在这里，因为并没有真的理解。  
所以，这里记录我从0开始对TDD的认知，看看能不能慢慢体会到大师的精髓。
#### TDD 模式开发流程
1. 理解需求 
2. 根据写测试代码
3. 快速完成业务代码，跑通测试用例
4. 提交代码到git工作分支
5. 重构业务代码，用测试用例保证重构不会引发bug

当面临改需求的时候，循环上面5个步骤。  
最重要的原则：先写测试，不要偷懒。


### golang 中的一些测试工具和方法 

#### 普通测试
```
func TestHello(t *testing.T) {
	got := Hello("Kobe")
	want := "Hello, Kobe"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
```
#### 子测试

```
func TestHello(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, got, want string) {
        // t.Helper() 可以告诉测试用例，这里是一个辅助函数，当代码失败时报告的行号是此函数被调用的行号
        t.Helper()
        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'world'", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

}
```

#### 基准测试
```
func BenchmarkRepeat(b *testing.B) {
      for i := 0; i < b.N; i++ {
          Repeat("a")
      }
  }`
  
  
 go test -bench=.

```

#### 检查测试覆盖率
```
go test -cover
```

#### Table driven tests（表格驱动测试）
```
func TestArea(t *testing.T) {
    // 可以通过如下命令来运行列表中指定的测试用例： go test -run TestArea/Rectangle
    areaTests := []struct {
        name    string
        shape   Shape
        hasArea float64
    }{
        {name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
        {name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
        {name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
    }

    for _, tt := range areaTests {
        // using tt.name from the case to use it as the `t.Run` test name
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.Area()
            if got != tt.hasArea {
                t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
            }
        })

    }

}
```

```


```