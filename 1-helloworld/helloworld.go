package __helloworld

const helloPrefix = "Hello, "

func Hello(name string) string {

	// return "Hello, world" // 1. 这样写，测试会不通过，提示 got 'Hello, world' want 'Hello, Kobe'
	// return "Hello, " + name  // 2. 测试通过, 这一步后，就可以提交到git了，但是不能提交到master
	// 3. TDD 的最后一步，是重构， 这里的重构过程中，消除了硬编码，重构之后在此测试，依然是通过的
	return helloPrefix + name
}

func Hello2(name string) string {
	return helloPrefix + name
}

func Hello3(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}
