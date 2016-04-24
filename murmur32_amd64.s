// +build amd64

#define c1_32 0xcc9e2d51
#define c2_32 0x1b873593

// AX = h1
// SI = &data[0]

#define sum32loop \
	MOVQ  CX, DX;                   \
	SHRQ  $2, DX;                   \
	SHLQ  $2, DX;                   \ // DX = end
	MOVQ  $0, BP;                   \ // BP = word pointer
	loop:                           \
	CMPQ  BP, DX;                   \
	JE    end                       \
	                                \
	MOVL  (SI)(BP*1), DI;           \ // DI = k1
	ADDQ  $4, BP;                   \
	IMULL $c1_32, DI;               \
	ROLL  $15, DI;                  \
	IMULL $c2_32, DI;               \
	XORL  AX, DI;                   \
	ROLL  $13, DI;                  \
	LEAL  0xe6546b64(DI)(DI*4), AX; \
	JMP   loop;                     \
	      \
	end:  \

// DX/DI = scratch
// BX = original clen
#define sum32tail \
	MOVL    $0, DX;    \ // DX is k1, DI is intermediate calc
	CMPQ    CX, $3;    \
	JNE     chk2;      \
	                   \
	on3:               \
	MOVBLZX 2(SI), DI; \
	SHLL    $16, DI;   \
	XORL    DI, DX;    \
	JMP     on2;       \
	                   \
	chk2:              \
	CMPQ    CX, $2;    \
	JNE     chk1;      \
	                   \
	on2:               \
	MOVBLZX 1(SI), DI; \
	SHLL    $8, DI;    \
	XORL    DI, DX;    \
	JMP     on1;       \
	                         \
	chk1:                    \
	CMPQ    CX, $1;          \
	JNE     chk0;            \
	                         \
	on1:                     \
	MOVBLZX (SI), DI;        \
	XORL    DI, DX;          \
	IMULL   $c1_32, DX;      \
	ROLL    $15, DX;         \
	IMULL   $c2_32, DX;      \
	XORL    DX, AX;          \
	                         \
	chk0:                    \
	XORL    BX, AX;          \
	MOVL    AX, DI;          \
	SHRL    $16, DI;         \
	XORL    DI, AX;          \
	IMULL   $0x85ebca6b, AX; \
	MOVL    AX, DI;          \
	SHRL    $13, DI;         \
	XORL    DI, AX;          \
	IMULL   $0xc2b2ae35, AX; \
	MOVL    AX, DI;          \
	SHRL    $16, DI;         \
	XORL    DI, AX;          \

#define sum32 \
	sum32loop    \
	MOVQ CX, BX; \ // save clen for sum32tail
	ADDQ DX, SI; \ // adjust SI to partial block
	SUBQ DX, CX; \ // adjust CX to partial block len
	sum32tail    \

TEXT ·Sum32(SB), $0-28
	MOVL $0, AX
	MOVQ data_base+0(FP), SI
	MOVQ data_len+8(FP), CX
	sum32
	MOVL AX, h1+24(FP)
	RET

TEXT ·SeedSum32(SB), $0-36
	MOVL seed+0(FP), AX
	MOVQ data_base+8(FP), SI
	MOVQ data_len+16(FP), CX
	sum32
	MOVL AX, h1+32(FP)
	RET
