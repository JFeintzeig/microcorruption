0e4a           mov	r10, r14
8e10           swpb	r14
3e80 d204      sub	#0x4d2, r14
0e52           add	sr, r14
0e10           rrc	r14
0eaa           dadd	r10, r14
0e11           rra	r14
c1d0 0e10      bis.b	pc, 0x100e(sp)
2ea0           dadd	@pc, r14
0e52           add	sr, r14
0e10           rrc	r14
2e50           add	@pc, r14
0e10           rrc	r14
0e10           rrc	r14
0fee           xor	r14, r15
3e41           pop	r14
2e53           incd	r14
004e           br	r14
3e41           pop	r14
2e53           incd	r14
004e           br	r14
3041           ret	
