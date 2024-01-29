3540 0026      mov	#0x2600, r5
0643           clr	r6
2455           add	@r5, r4
8410           swpb	r4
36e5           xor	@r5+, r6
06e4           xor	r4, r6
04e6           xor	r6, r4
06e4           xor	r4, r6
8593 0000      tst	0x0(r5)
0742           mov	sr, r7
21f3           and	#0x2, r7
0711           rra	r7
17e3           xor	#0x1, r7
8710           swpb	r7
0711           rra	r7
8711           sxt	r7
8710           swpb	r7
8711           sxt	r7
3840 184b      mov	#0x4b18, r8
08f7           and	r7, r8
37e3           xor	#-0x1, r7
37f0 aa47      and	#0x47aa, r7
0857           add	r7, r8
0743           clr	r7
0c48           mov	r8, r12
2455           add	@r5, r4
8410           swpb	r4
36e5           xor	@r5+, r6
06e4           xor	r4, r6
04e6           xor	r6, r4
06e4           xor	r4, r6
8593 0000      tst	0x0(r5)
0742           mov	sr, r7
27f3           and	#0x2, r7
0711           rra	r7
17e3           xor	#0x1, r7
8710           swpb	r7
0711           rra	r7
8711           sxt	r7
8710           swpb	r7
8711           sxt	r7
3840 184b      mov	#0x4b18, r8
08f7           and	r7, r8
37e3           xor	#-0x1, r7
37f0 aa47      and	#0x47aa, r7
0857           add	r7, r8
0743           clr	r7
0c48           mov	r8, r12
3490 b1fe      cmp	#0xfeb1, r4
0742           mov	sr, r7
0443           clr	r4
3690 9892      cmp	#0x9298, r6
07f2           and	sr, r7
0643           clr	r6
0711           rra	r7
8710           swpb	r7
0711           rra	r7
0711           rra	r7
0711           rra	r7
0711           rra	r7
02d7           bis	r7, sr
