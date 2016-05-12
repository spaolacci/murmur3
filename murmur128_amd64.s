// +build go1.5,amd64

#define c1_128 0x87c37b91114253d5
#define c2_128 0x4cf5ad432745937f

// uses DI
#define fmix64(k) \
	MOVQ  k, DI;   \
	SHRQ  $33, DI; \
	XORQ  DI, k;   \
	IMULQ R9, k;   \
	MOVQ  k, DI;   \
	SHRQ  $33, DI; \
	XORQ  DI, k;   \
	IMULQ R10, k;  \
	MOVQ  k, DI;   \
	SHRQ  $33, DI; \
	XORQ  DI, k;   \

// AX = h1
// BX = h2
// SI = &data[0]

// R9 = big const
#define sum128loop \
	MOVQ  CX, DX;                   \
	SHRQ  $4, DX;                   \
	SHLQ  $4, DX;                   \ // DX = end
	MOVQ  $0, BP;                   \ // BP = quad pointer
	loop:                           \
	CMPQ  BP, DX;                   \
	JE    end;                      \
	                                \
	MOVQ  (SI)(BP*1), R11;          \ // R11 = k1
	MOVQ  8(SI)(BP*1), R12;         \ // R12 = k2
	ADDQ  $16, BP;                  \
	IMULQ R9, R11;                  \
	ROLQ  $31, R11;                 \
	IMULQ R10, R11;                 \
	XORQ  R11, AX;                  \
	                                \
	ROLQ  $27, AX;                  \
	ADDQ  BX, AX;                   \
	LEAQ  0x52dce729(AX)(AX*4), AX; \
	                                \
	IMULQ R10, R12;                 \
	ROLQ  $33, R12;                 \
	IMULQ R9, R12;                  \
	XORQ  R12, BX;                  \
	                                \
	ROLQ  $31, BX;                  \
	ADDQ  AX, BX;                   \
	LEAQ  0x38495ab5(BX)(BX*4), BX; \
	JMP   loop;                     \
	      \
	end:  \

// DX/DI = scratch
// BP = original clen
#define sum128tail \
	MOVQ    $0, DX;     \ // DX is k1 and k2, DI is intermediate calc
	CMPQ    CX, $15;    \
	JNE     chk14;      \
	                    \
	on15:               \
	MOVBQZX 14(SI), DI; \
	SHLQ    $48, DI;    \
	XORQ    DI, DX;     \
	JMP     on14;       \
	                    \
	chk14:              \
	CMPQ    CX, $14;    \
	JNE     chk13;      \
	                    \
	on14:               \
	MOVBQZX 13(SI), DI; \
	SHLQ    $40, DI;    \
	XORQ    DI, DX;     \
	JMP     on13;       \
	                    \
	chk13:              \
	CMPQ    CX, $13;    \
	JNE     chk12;      \
	                    \
	on13:               \
	MOVBQZX 12(SI), DI; \
	SHLQ    $32, DI;    \
	XORQ    DI, DX;     \
	JMP     on12;       \
	                    \
	chk12:              \
	CMPQ    CX, $12;    \
	JNE     chk11;      \
	                    \
	on12:               \
	MOVBQZX 11(SI), DI; \
	SHLQ    $24, DI;    \
	XORQ    DI, DX;     \
	JMP     on11;       \
	                    \
	chk11:              \
	CMPQ    CX, $11;    \
	JNE     chk10;      \
	                    \
	on11:               \
	MOVBQZX 10(SI), DI; \
	SHLQ    $16, DI;    \
	XORQ    DI, DX;     \
	JMP     on10;       \
	                   \
	chk10:             \
	CMPQ    CX, $10;   \
	JNE     chk9;      \
	                   \
	on10:              \
	MOVBQZX 9(SI), DI; \
	SHLQ    $8, DI;    \
	XORQ    DI, DX;    \
	JMP     on9;       \
	                   \
	chk9:              \
	CMPQ    CX, $9;    \
	JNE     chk8;      \
	                   \
	on9:               \
	MOVBQZX 8(SI), DI; \
	XORQ    DI, DX;    \
	IMULQ   R10, DX;   \
	ROLQ    $33, DX;   \
	IMULQ   R9, DX;    \
	XORQ    DX, BX;    \
	MOVQ    $0, DX;    \ // reset for k2
	JMP     on8;       \
	                   \
	chk8:              \
	CMPQ    CX, $8;    \
	JNE     chk7;      \
	                   \
	on8:               \
	MOVBQZX 7(SI), DI; \
	SHLQ    $56, DI;   \
	XORQ    DI, DX;    \
	JMP     on7;       \
	                   \
	chk7:              \
	CMPQ    CX, $7;    \
	JNE     chk6;      \
	                   \
	on7:               \
	MOVBQZX 6(SI), DI; \
	SHLQ    $48, DI;   \
	XORQ    DI, DX;    \
	JMP     on6;       \
	                   \
	chk6:              \
	CMPQ    CX, $6;    \
	JNE     chk5;      \
	                   \
	on6:               \
	MOVBQZX 5(SI), DI; \
	SHLQ    $40, DI;   \
	XORQ    DI, DX;    \
	JMP     on5;       \
	                   \
	chk5:              \
	CMPQ    CX, $5;    \
	JNE     chk4;      \
	                   \
	on5:               \
	MOVBQZX 4(SI), DI; \
	SHLQ    $32, DI;   \
	XORQ    DI, DX;    \
	JMP     on4;       \
	                   \
	chk4:              \
	CMPQ    CX, $4;    \
	JNE     chk3;      \
	                   \
	on4:               \
	MOVBQZX 3(SI), DI; \
	SHLQ    $24, DI;   \
	XORQ    DI, DX;    \
	JMP     on3;       \
	                   \
	chk3:              \
	CMPQ    CX, $3;    \
	JNE     chk2;      \
	                   \
	on3:               \
	MOVBQZX 2(SI), DI; \
	SHLQ    $16, DI;   \
	XORQ    DI, DX;    \
	JMP     on2;       \
	                   \
	chk2:              \
	CMPQ    CX, $2;    \
	JNE     chk1;      \
	                   \
	on2:               \
	MOVBQZX 1(SI), DI; \
	SHLQ    $8, DI;    \
	XORQ    DI, DX;    \
	JMP     on1;       \
	                                  \
	chk1:                             \
	CMPQ    CX, $1;                   \
	JNE     chk0;                     \
	                                  \
	on1:                              \
	MOVBQZX (SI), DI;                 \
	XORQ    DI, DX;                   \
	IMULQ   R9, DX;                   \
	ROLQ    $31, DX;                  \
	IMULQ   R10, DX;                  \
	XORQ    DX, AX;                   \
	                                  \
	chk0:                             \
	XORQ    BP, AX;                   \
	XORQ    BP, BX;                   \
	ADDQ    BX, AX;                   \
	ADDQ    AX, BX;                   \
	MOVQ    $0xff51afd7ed558ccd, R9;  \
	MOVQ    $0xc4ceb9fe1a85ec53, R10; \
	fmix64(AX);                       \
	fmix64(BX);                       \
	ADDQ    BX, AX;                   \
	ADDQ    AX, BX;                   \

#define sum128 \
	MOVQ $c1_128, R9;  \
	MOVQ $c2_128, R10; \
	sum128loop         \
	MOVQ CX, BP;       \ // save clen for sum128tail
	ADDQ DX, SI;       \ // adjust SI to partial block
	SUBQ DX, CX;       \ // adjust CS to partial block len
	sum128tail         \

TEXT ·Sum128(SB), $0-40
	MOVQ $0, AX
	MOVQ $0, BX
	MOVQ data_base+0(FP), SI
	MOVQ data_len+8(FP), CX
	sum128
	MOVQ AX, h1+24(FP)
	MOVQ BX, h2+32(FP)
	RET

TEXT ·SeedSum128(SB), $0-56
	MOVQ seed1+0(FP), AX
	MOVQ seed2+8(FP), BX
	MOVQ data_base+16(FP), SI
	MOVQ data_len+24(FP), CX
	sum128
	MOVQ AX, h1+40(FP)
	MOVQ BX, h2+48(FP)
	RET
