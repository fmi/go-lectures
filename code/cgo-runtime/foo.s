
--- prog list "Add" ---
0000 (foo.go:3) TEXT    Add+0(SB),$0-24
0001 (foo.go:3) LOCALS  ,$0
0002 (foo.go:3) TYPE    a+0(FP){int},$8
0003 (foo.go:3) TYPE    b+8(FP){int},$8
0004 (foo.go:3) TYPE    ~anon2+16(FP){int},$8
0005 (foo.go:4) MOVQ    a+0(FP),BX
0006 (foo.go:4) MOVQ    b+8(FP),BP
0007 (foo.go:4) ADDQ    BP,BX
0008 (foo.go:4) MOVQ    BX,~anon2+16(FP)
0009 (foo.go:4) RET     ,

--- prog list "init" ---
0010 (foo.go:5) TEXT    init+0(SB),$0-0
0011 (foo.go:5) MOVBQZX initdone·+0(SB),AX
0012 (foo.go:5) LOCALS  ,$0
0013 (foo.go:5) CMPB    AX,$0
0014 (foo.go:5) JEQ     ,20
0015 (foo.go:5) CMPB    AX,$2
0016 (foo.go:5) JNE     ,18
0017 (foo.go:5) RET     ,
0018 (foo.go:5) CALL    ,runtime.throwinit+0(SB)
0019 (foo.go:5) UNDEF   ,
0020 (foo.go:5) MOVB    $2,initdone·+0(SB)
0021 (foo.go:5) RET     ,
