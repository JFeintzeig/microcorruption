3140 0044      mov	#0x4400, sp
1542 5c01      mov	&0x015c, r5
75f3           and.b	#-0x1, r5
35d0 085a      bis	#0x5a08, r5
3f40 0011      mov	#0x1100, r15
0f93           tst	r15
0724           jz	$+0x10
8245 5c01      mov	r5, &0x015c
2f83           decd	r15
0343           clr	4
1e4f 3446      mov	0x4634(r15), r14
8f4e 0024      mov	r14, 0x2400(r15)
ef23           jnz	$-0x20
0f43           clr	r15
0f93           tst	r15
0e24           jz	$+0x1e
8245 5c01      mov	r5, &0x015c
1f83           dec	r15
cf43 0035      mov.b	#0x0, 0x3500(r15)
f923           jnz	$-0xc
3e40 0012      mov	#0x1200, r14
3f40 0024      mov	#0x2400, r15
bf4f feef      mov	@r15+, -0x1002(r15)
3e53           add	#-0x1, r14
fa23           jnz	$-0xa
3b40 0c16      mov	#0x160c, r11
0212           push	sr
3040 be44      br	#0x44be	
