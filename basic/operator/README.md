# 运算符与分隔符

以下字符序列表示运算符，分隔符和其它特殊标记：

    +    &     +=    &=     &&    ==    !=    (    )
    -    |     -=    |=     ||    <     <=    [    ]
    *    ^     *=    ^=     <-    >     >=    {    }
    /    <<    /=    <<=    ++    =     :=    ,    ;
    %    >>    %=    >>=    --    !     ...   .    :
    &^          &^=

# 操作符优先级

一元操作符拥有最高优先级。 ++ 和 -- 操作符是语句，而非表达式，它们不属于运算符一级。 因此，语句 *p++ 等价于 (*p)++。

二元操作符有五种优先级。 乘法操作符结合性最强，其次为加法操作符、比较操作符、&&（逻辑与）， 最后为 ||（逻辑或）：

          优先级        操作符
    	5             *  /  %  <<  >>  &  &^
    	4             +  -  |  ^
    	3             ==  !=  <  <=  >  >=
    	2             &&
    	1             ||
相同优先级的二元操作符从左到右结合。 例如，x / y * z 等价于 (x / y) * z。

    +x
    23 + 3*x[i]
    x <= f()
    ^a >> b
    f() || g()
    x == y+1 && <-chanPtr > 0

# 算数操作符

算数操作符适用于数值，并产生相同类型的结果作为第一个操作数。四个基本算数操作符 （+，-，*，/）适用于整数、浮点数和复数类型； + 也适用于字符串。其它所有算数操作符仅适用于整数。

    +    和              integers, floats, complex values, strings
    -    差              integers, floats, complex values
    *    积              integers, floats, complex values
    /    商              integers, floats, complex values
    %    余              integers

    &    按位与          integers
    |    按位或          integers
    ^    按位异或        integers
    &^   位清除（与非）  integers

    <<   向左移位        integer << unsigned integer
    >>   向右移位        integer >> unsigned integer
字符串可使用 + 操作符连结或 += 赋值操作符：

    s := "hi" + string(c)
    s += " and good bye"
字符串加法通过连结操作数创建一个新的字符串。