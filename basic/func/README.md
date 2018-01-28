# 函数类型

函数类型表示所有带相同形参和返回类型的集。未初始化的函数类型变量的的值为 nil。

    func()
    func(x int) int
    func(a, _ int, z float32) bool
    func(a, b int, z float32) (bool)
    func(prefix string, values ...int)
    func(a, b int, z float64, opt ...interface{}) (success bool)
    func(int, int, float64) (float64, *[]int)
    func(n int) func(p *T)