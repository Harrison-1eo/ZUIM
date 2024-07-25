// zuc.h

#ifndef ZUC_H
#define ZUC_H

#include <stdint.h>

typedef unsigned char u8;
typedef unsigned int u32;

typedef struct
{
    u32 LFSR_S0;
    u32 LFSR_S1;
    u32 LFSR_S2;
    u32 LFSR_S3;
    u32 LFSR_S4;
    u32 LFSR_S5;
    u32 LFSR_S6;
    u32 LFSR_S7;
    u32 LFSR_S8;
    u32 LFSR_S9;
    u32 LFSR_S10;
    u32 LFSR_S11;
    u32 LFSR_S12;
    u32 LFSR_S13;
    u32 LFSR_S14;
    u32 LFSR_S15;
    u32 F_R1;
    u32 F_R2;
    u32 BRC_X0;
    u32 BRC_X1;
    u32 BRC_X2;
    u32 BRC_X3;
} ZUC_State;

void ZUC_Init(ZUC_State* state, u8* k, u8* iv);
void ZUC_Next(ZUC_State* state, u32* pKeystream, int KeystreamLen);

#endif // ZUC_H
