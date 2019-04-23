package __helloworld

import "testing"

// 需求： 编写一个函数，实现向指定人打招呼的功能
// 先写测试
// 写好测试后，运行 go test -run ^TestHello$
// 此时当然会提示 Hello 函数不存在
func TestHello(t *testing.T) {
	got := Hello("Kobe")
	want := "Hello, Kobe"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

// 接下来，需求改了，当传入的参数为空字符串时，返回"Hello, World"，而不是 "Hello, ".
// 注意，这就要求有多个测试用例
func TestHello2(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello2("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello2("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}

// 上面的测试代码可以用自测试的方式重构
// 运行测试用例发现报错，修改代码直到测试通过
func TestHello3(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// t.Helper() 可以告诉测试用例，这里是一个辅助函数，当代码失败时报告的行号是此函数被调用的行号
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello3("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello3("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {

}
