# 导入声明

假定我们拥有包含包子句 package math 的已编译包，它导出函数 Sin， 且该包已被安装在标识为"lib/math"的文件中。此表单阐明了在各种类型的导入声明之后， Sin 在导入该包的文件中如何访问。
    导入声明                    Sin的本地名

    import   "lib/math"         math.Sin
    import m "lib/math"         m.Sin
    import . "lib/math"         Sin
导入声明用来声明导入包与被导入包之间的从属关系。包导入其自身或导入一个不涉及任何可导出标识符的包是非法的。 要为包的副作用（初始化）而单独导入它，需使用空白标识符作为明确的包名：

    import _ "lib/math"