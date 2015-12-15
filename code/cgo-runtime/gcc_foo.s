	.file	"foo.go"
	.section	.go_export,"",@progbits
	.ascii	"v1;\n"
	.ascii	"package "
	.ascii	"main"
	.ascii	";\n"
	.ascii	"pkgpath "
	.ascii	"main"
	.ascii	";\n"
	.ascii	"priority 1;\n"
	.ascii	"func "
	.ascii	"Add"
	.ascii	" ("
	.ascii	"a"
	.ascii	" "
	.ascii	"<type -11>"
	.ascii	", "
	.ascii	"b"
	.ascii	" "
	.ascii	"<type -11>"
	.ascii	")"
	.ascii	" "
	.ascii	"<type -11>"
	.ascii	";\n"
	.ascii	"checksum FD80E13A21B63639D0FE0BD8A2F5B84400D3F2A9;\n"
	.globl	main.Add$descriptor
	.section	.rodata.main.Add$descriptor,"a",@progbits
	.align 8
	.type	main.Add$descriptor, @object
	.size	main.Add$descriptor, 8
main.Add$descriptor:
	.quad	main.Add
	.text
	.globl	main.Add
	.type	main.Add, @function
main.Add:
.LFB0:
	.cfi_startproc
	cmpq	%fs:112, %rsp
	jae	.L3
	movl	$8, %r10d
	movl	$0, %r11d
	call	__morestack
	ret
.L3:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset 6, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register 6
	movq	%rdi, -24(%rbp)
	movq	%rsi, -32(%rbp)
	movq	$0, -8(%rbp)
	movq	-32(%rbp), %rax
	movq	-24(%rbp), %rdx
	addq	%rdx, %rax
	movq	%rax, -8(%rbp)
	movq	-8(%rbp), %rax
	popq	%rbp
	.cfi_def_cfa 7, 8
	ret
	.cfi_endproc
.LFE0:
	.size	main.Add, .-main.Add
	.globl	__go_init_main
	.type	__go_init_main, @function
__go_init_main:
.LFB1:
	.cfi_startproc
	cmpq	%fs:112, %rsp
	jae	.L5
	movl	$8, %r10d
	movl	$0, %r11d
	call	__morestack
	ret
.L5:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset 6, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register 6
	popq	%rbp
	.cfi_def_cfa 7, 8
	ret
	.cfi_endproc
.LFE1:
	.size	__go_init_main, .-__go_init_main
	.section	.note.GNU-split-stack,"",@progbits
	.ident	"GCC: (GNU) 4.8.2"
	.section	.note.GNU-stack,"",@progbits
