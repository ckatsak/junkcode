#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <sys/ptrace.h>
#include <sys/user.h>
#include <sys/reg.h>
#include <signal.h>

int main(void)
{
	int status = 0, pid, r;
	struct user_regs_struct uregs;

	if ((pid = fork()) == 0) {
		fprintf(stderr, "[%d] <--- My pid\n", (int)getpid());
		ptrace(PTRACE_TRACEME, 0, 0, 0);
		kill(getpid(), SIGINT);
		r = getpid();
		fprintf(stderr, "%d", r);
	}
	else {
		wait(&status);
		ptrace(PTRACE_SYSCALL, pid, 0, 0);
		wait(&status);
		ptrace(PTRACE_GETREGS, pid, 0, &uregs);

		/* this prints the syscall number of getpid */
//		fprintf(stderr, "syscall nr: %d\n", uregs.orig_eax);
		fprintf(stderr, "syscall nr: %d\n", uregs.orig_eax);
		/* 64 is syscall number of getppid */
//		uregs.orig_eax = 64;
		uregs.orig_eax = 64;
		ptrace(PTRACE_SETREGS, pid, 0, &uregs);
		ptrace(PTRACE_CONT, pid, 0, 0);
		wait(&status);
		if (WIFEXITED(status))
			fprintf(stderr, "we're done\n");
	}

	return 0;
}
