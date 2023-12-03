# GO 语言结构与语法基础
1. 包声明
2. 引入包
3. 函数
4. init函数
5. 变量
6. 标识符
7. 行分隔符
8. 语句&表达式
9. 注释
10. 公有成员与私有成员
11. 关键字、保留字和预定义标志

# 引用类型
1. 切片
2. map
3. channel
4. interface
5. func

# 关键词
1. break 跳转语句，跳出循环或跳出switch语句，可跑转到指定的 标签位置
2. default 默认选项，switch 或 select 默认操作项
3. func 定义一个函数
4. interface 声明一个接口
5. select golang语言层面 I/O多路复用机制，用于检测管道是否就绪，与case和default一起使用
6. case 用于switch 或 select 语句块，case用于指定一个或多个值（常量或者表达式）表示满足条件方可执行其他中的语句块
7. defer 方法延迟调用关键字
8. go 启动协程的关键字
9. map 集合类型
10. struct 结构体类型
11. chan 通道类型
12. else 条件表达式 否则
13. goto 跳转到某一个 标签位置
14. package 声明包名
15. switch 流程控制语句，根据不同的条件执行不同的语句块，与 case 和 default一起使用
16. const 声明常量
17. fallthrough 表示通过当前语句块，switch语句中表示可以执行下一个一刻钟块
18. if 条件表达式 如果
19. range 用于 for 循环，遍历数组、切片、集合、管道等类型，获取其中元素、索引、键值
20. type 类型定义的关键词，声明一个类型
21. continue 与 for 循环使用，跳出当次循环
22. for 循环语句
23. import 包导入语句
24. return 方法返回语句
25. var 变量声明的关键词

# 预定义标识
1. append附加，向切片附加元素
2. cap 获取容量、数组、切片、通道的容量
3. close 关闭通道
4. copy 用于切片的拷贝
5. imag 返回复数的虚部
6. real 返回复数的实部
7. panic 抛出异常消息
8. recover 恢复因异常中断的协程，并返回异常消息
9. iota 常量计数器，可与常量配合使用，实现枚举的能力
10. len 获取数组、切片、集合、通道等类型的长度
11. make 用于初始化 切片、集合、通道并返回其他对象
12. new 创建一个类型的变量并为其分配内存空间，并返回类型的指针，常用于结构体变量的创建
13. nil 引用类型的零值
14. false,true 布类型信息
15. print 打印消息，不换行
16. println 打印并换行
17. bool,byte,
18. complex,complex64,complex128
19. float32,float64
20. int,uint,int8,uint8,int16,uint16,int32,uint32,int64,uint64
21. string
22.uintptr

