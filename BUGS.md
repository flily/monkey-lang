Known Bugs in offical implements
=================================

Bugs are verified and reproduced in offical final work in both WAIIG and WACIG.

Versions are presented below, and the links to the code will not be provided. You can easily find
the links in preface of the books, and get the codes. 

+ Interpreter, WAIIG 1.7, in `04` directory.
+ Compiler, WACIG 1.2, in `10` directory.


## 1. Let statement crush in variable self-assignment
Code:
```monkey
let a = 1;
let a = a + 1;
puts(a);
```

In interpreter, it works well.
```monkey
>> let a = 1;
>> let a = a + 1;
>> puts(a)
2
null
>> 
```

But in compiler, it makes VM crush.
```monkey
>> let a = 1;
1
>> let a = a + 1;
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x2 addr=0x20 pc=0x1043784dc]

goroutine 1 [running]:
monkey/vm.(*VM).executeBinaryOperation(0x14000108150, 0x0?)
    /tmp/wacig_code_1.2/code/10/src/monkey/vm/vm.go:304 +0x8c
monkey/vm.(*VM).Run(0x14000108150)
    /tmp/wacig_code_1.2/code/10/src/monkey/vm/vm.go:87 +0x1cc
monkey/repl.Start({0x1043c1228?, 0x1400000e010?}, {0x1043c1248, 0x1400000e018})
    /tmp/wacig_code_1.2/code/10/src/monkey/repl/repl.go:55 +0x418
main.main()
    /tmp/wacig_code_1.2/code/10/src/monkey/main.go:18 +0xc0
```
