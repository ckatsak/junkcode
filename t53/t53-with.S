	.file	"t53.c"
	.section	.rodata
.LC0:
	.string	"i == x == %d\n"
	.text
	.globl	main
	.type	main, @function
main:
.LFB2:
	.cfi_startproc
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset 6, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register 6
	subq	$16, %rsp
	movl	$0, %edi
	call	time
	movl	%eax, %edi
	call	srand
	call	rand
	cltq
	movq	%rax, -8(%rbp)
	movl	$0, -12(%rbp)
	jmp	.L2
.L5:
	movl	-12(%rbp), %eax
	cltq
	cmpq	-8(%rbp), %rax
	sete	%al
	movzbl	%al, %eax
	testq	%rax, %rax
	je	.L3
	movq	stderr(%rip), %rax
	movl	-12(%rbp), %edx
	movl	$.LC0, %esi
	movq	%rax, %rdi
	movl	$0, %eax
	call	fprintf
	jmp	.L4
.L3:
	addl	$1, -12(%rbp)
.L2:
	cmpl	$2147483647, -12(%rbp)
	jne	.L5
.L4:
	movl	$0, %eax
	leave
	.cfi_def_cfa 7, 8
	ret
	.cfi_endproc
.LFE2:
	.size	main, .-main
	.ident	"GCC: (Ubuntu 5.4.0-6ubuntu1~16.04.4) 5.4.0 20160609"
	.section	.note.GNU-stack,"",@progbits
