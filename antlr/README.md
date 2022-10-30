Calc.g4文件解读：
- 可以看到有的规则由大写开头的单词描述，有的由小写开头的单词描述
    - 大写开头的单词描述的是parser rule（语法规则）
    - 小写开头的单词描述的是lexical rule（词法规则）
- 第1行注释是生成目标代码的命令行，-o参数对应的值需要修改成自己需要的值
- 第2行grammar为关键字，表示定义了一种名为Calc的语法
- 第5行定义了该语法的一条规则，它的名称为prog，值由多个stat元素构成，stat元素的描述从第8行开始；# calc表示为名称为prog的规则打上一个名为calc的标签，用于生成visitor代码时访问rule，可以参考下一小节的示例代码
- 第8行到第11行描述名称为stat的规则，它由3种情况构成，'|'操作符表示逻辑运算符或
    - 第8行表示第一种情况，由expr规则和NEWLINE标识符组成
    - 第9行表示第二种情况，由左中右3部分构成，左值为ID标识符，中值为操作符，右值为expr规则和NEWLINE标识符
    - 第10行表示第三种情况，仅为NEWLINE标识符
- 第13行到第18行描述名称为expr的规则，它由5种情况构成
    - 第13行 op=('*'|'/') op的用途是给*、/这2个操作符打上了op的标签，可以在代码中通过GetOp()f方法获取到这2个操作符
- 第20行到第23行定义了4种规则，表达方式和正则表达式相似
    - 第20行定义了ID规则，它由大小写英文字母组成
    - 第21行定义了INT规则，它是阿拉伯数字构成的正整数
    - 第22行定义了NEWLINE规则，它由\r或\n表示
    - 第23行定义了WS规则，它表示遇到空格就跳过
- 第25行到第28行为语法中定义的操作符打上标签，用途是生成visitor代码时可以访问相应的操作符，可以参考下一小节的示例代码

实现代码2022/04/22_2.go解读：
- ANTLR鼓励开发者定义完基础的语法后，将更多的逻辑处理放到开发者自己的应用代码中，基于访问者模式实现业务逻辑。
- ANLTR为我们生成了visitor接口和最基础的实现类，分别位于calc_visitor.go和calc_base_visitor.go中。
- 第11行到第14行定义了开发者自定义的visitor，在其中实现visitor中定义的接口
- 第23行Visit方法负责解析由parser生成的语法树，也就是开始执行计算工作的入口
- 第24行tree.Accept(visitor)使用访问者模式，将遍历语法树过程中所执行的逻辑与语法树本身解耦
- 第32行开始VisitAssign、VisitPrintExpr等方法表示处理Calc.g4文件中定义的各类规则，方法名为visit+g4文件中#符号后的标签名称
  - 同时也可以看到有CalcContext、AssignContext等以XXXContext命名的结构体，他们是语法树。在处理树形的数据结构时，函数的入参往往是一个节点，在这里同理，XXXContext可以视为TreeNode
  - VisitXXX方法要做的就是遍历这颗语法树，处理每一个树节点
- 第93行到96行对应lexer、parser中的概念
  - 由lexer对输入做词法分析，生成token流
  - parser对lexer生成的token流进行语法分析，生成语法树
  - 这部分代码为获取语法树的模板代码
- 当有了语法树以后，就可以遍历语法树进行计算
- 执行go run calc.go input.file

ANTLR在生成代码时可以通过-listener、-visitor选项生成listener和visitor。
在上面的demo中，我们在visitor中实现了遍历语法树时需要做的操作。在listener也可以实现对语法树的操作，举一个简单的例子看一下listener的用法2022/04/22_1.go：
- 在执行ANTLR命令时加上-listener选项生成listener接口
- 这段代码实现的功能是：打印语法树每个节点，从而展示整个语法树
- 注意第18行，listener需要通过walker来调用，walker是遍历语法树的代码
  listener和visitor之间的区别是：
  listener需要通过walker来调用执行，而visitor则通过主动调用自身的Visit方法来遍历语法树。当walker执行时，它会完整地遍历语法树；如果不显式地调用visitor的Visit方法，则当前节点下的子树不会被遍历。可以看出，visitor的功能更加灵活。