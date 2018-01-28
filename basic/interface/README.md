# 接口类型

接口类型指定一个称为 接口 的 方法集。 接口类型变量可存储任何带方法集类型的值，该方法集为此接口的超集。 这种类型表示 实现此接口。未初始化的接口类型变量的值为 nil。

对于所有的方法集，在一个接口类型中，每个方法必须有唯一的名字。

    // 一个简单的File接口
    interface {
    	Read(b Buffer) bool
    	Write(b Buffer) bool
    	Close()
    }
不止一个类型可实现同一接口。例如，若两个类型 S1 和 S2 拥有方法集

    func (p T) Read(b Buffer) bool { return … }
    func (p T) Write(b Buffer) bool { return … }
    func (p T) Close() { … }
（其中 T 代表 S1 或 S2） 那么 File 接口都将被 S1 和 S2所实现， 不论如何，方法 S1 和 S2 都会拥有或共享它。

类型可实现任何接口，包括任何其方法的子集，因此可能实现几个不同的接口。 例如，所有类型都实现了 空接口：

# 嵌入接口

一个接口可通过包含一个名为 T 的接口类型来代替一个方法的实现。 这称之为嵌入接口，其效果等价于在接口中显式枚举出 T 中的方法。

    interface{}

    type ReadWrite interface {
    	Read(b Buffer) bool
    	Write(b Buffer) bool
    }

    type File interface {
    	ReadWrite  // 等价于枚举ReadWrite中的方法
    	Lock       // 等价于枚举Lock中的方法
    	Close()
    }

