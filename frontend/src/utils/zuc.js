const S0 = [
    0x3e, 0x72, 0x5b, 0x47, 0xca, 0xe0, 0x00, 0x33, 0x04, 0xd1, 0x54, 0x98, 0x09, 0xb9, 0x6d, 0xcb,
    0x7b, 0x1b, 0xf9, 0x32, 0xaf, 0x9d, 0x6a, 0xa5, 0xb8, 0x2d, 0xfc, 0x1d, 0x08, 0x53, 0x03, 0x90,
    0x4d, 0x4e, 0x84, 0x99, 0xe4, 0xce, 0xd9, 0x91, 0xdd, 0xb6, 0x85, 0x48, 0x8b, 0x29, 0x6e, 0xac,
    0xcd, 0xc1, 0xf8, 0x1e, 0x73, 0x43, 0x69, 0xc6, 0xb5, 0xbd, 0xfd, 0x39, 0x63, 0x20, 0xd4, 0x38,
    0x76, 0x7d, 0xb2, 0xa7, 0xcf, 0xed, 0x57, 0xc5, 0xf3, 0x2c, 0xbb, 0x14, 0x21, 0x06, 0x55, 0x9b,
    0xe3, 0xef, 0x5e, 0x31, 0x4f, 0x7f, 0x5a, 0xa4, 0x0d, 0x82, 0x51, 0x49, 0x5f, 0xba, 0x58, 0x1c,
    0x4a, 0x16, 0xd5, 0x17, 0xa8, 0x92, 0x24, 0x1f, 0x8c, 0xff, 0xd8, 0xae, 0x2e, 0x01, 0xd3, 0xad,
    0x3b, 0x4b, 0xda, 0x46, 0xeb, 0xc9, 0xde, 0x9a, 0x8f, 0x87, 0xd7, 0x3a, 0x80, 0x6f, 0x2f, 0xc8,
    0xb1, 0xb4, 0x37, 0xf7, 0x0a, 0x22, 0x1, 0x28, 0x7c, 0xcc, 0x3c, 0x89, 0xc7, 0xc3, 0x96, 0x56,
    0x07, 0xbf, 0x7e, 0xf0, 0x0b, 0x2b, 0x97, 0x52, 0x35, 0x41, 0x79, 0x61, 0xa6, 0x4c, 0x10, 0xfe,
    0xbc, 0x26, 0x95, 0x88, 0x8a, 0xb0, 0xa3, 0xfb, 0xc0, 0x18, 0x94, 0xf2, 0xe1, 0xe5, 0xe9, 0x5d,
    0xd0, 0xdc, 0x11, 0x66, 0x64, 0x5c, 0xec, 0x59, 0x42, 0x75, 0x12, 0xf5, 0x74, 0x9c, 0xaa, 0x23,
    0x0e, 0x86, 0xab, 0xbe, 0x2a, 0x02, 0xe7, 0x67, 0xe6, 0x44, 0xa2, 0x6c, 0xc2, 0x93, 0x9f, 0xf1,
    0xf6, 0xfa, 0x36, 0xd2, 0x50, 0x68, 0x9e, 0x62, 0x71, 0x15, 0x3d, 0xd6, 0x40, 0xc4, 0xe2, 0x0f,
    0x8e, 0x83, 0x77, 0x6b, 0x25, 0x05, 0x3f, 0x0c, 0x30, 0xea, 0x70, 0xb7, 0xa1, 0xe8, 0xa9, 0x65,
    0x8d, 0x27, 0x1a, 0xdb, 0x81, 0xb3, 0xa0, 0xf4, 0x45, 0x7a, 0x19, 0xdf, 0xee, 0x78, 0x34, 0x60
];

const S1 = [
    0x55, 0xc2, 0x63, 0x71, 0x3b, 0xc8, 0x47, 0x86, 0x9f, 0x3c, 0xda, 0x5b, 0x29, 0xaa, 0xfd, 0x77,
    0x8c, 0xc5, 0x94, 0x0c, 0xa6, 0x1a, 0x13, 0x00, 0xe3, 0xa8, 0x16, 0x72, 0x40, 0xf9, 0xf8, 0x42,
    0x44, 0x26, 0x68, 0x96, 0x81, 0xd9, 0x45, 0x3e, 0x10, 0x76, 0xc6, 0xa7, 0x8b, 0x39, 0x43, 0xe1,
    0x3a, 0xb5, 0x56, 0x2a, 0xc0, 0x6d, 0xb3, 0x05, 0x22, 0x66, 0xbf, 0xdc, 0x0b, 0xfa, 0x62, 0x48,
    0xdd, 0x20, 0x11, 0x06, 0x36, 0xc9, 0xc1, 0xcf, 0xf6, 0x27, 0x52, 0xbb, 0x69, 0xf5, 0xd4, 0x87,
    0x7f, 0x84, 0x4c, 0xd2, 0x9c, 0x57, 0xa4, 0xbc, 0x4f, 0x9a, 0xdf, 0xfe, 0xd6, 0x8d, 0x7a, 0xeb,
    0x2b, 0x53, 0xd8, 0x5c, 0xa1, 0x14, 0x17, 0xfb, 0x23, 0xd5, 0x7d, 0x30, 0x67, 0x73, 0x08, 0x09,
    0xee, 0xb7, 0x70, 0x3f, 0x61, 0xb2, 0x19, 0x8e, 0x4e, 0xe5, 0x4b, 0x93, 0x8f, 0x5d, 0xdb, 0xa9,
    0xad, 0xf1, 0xae, 0x2e, 0xcb, 0x0d, 0xfc, 0xf4, 0x2d, 0x46, 0x6e, 0x1d, 0x97, 0xe8, 0xd1, 0xe9,
    0x4d, 0x37, 0xa5, 0x75, 0x5e, 0x83, 0x9e, 0xab, 0x82, 0x9d, 0xb9, 0x1c, 0xe0, 0xcd, 0x49, 0x89,
    0x01, 0xb6, 0xbd, 0x58, 0x24, 0xa2, 0x5f, 0x38, 0x78, 0x99, 0x15, 0x90, 0x50, 0xb8, 0x95, 0xe4,
    0xd0, 0x91, 0xc7, 0xce, 0xed, 0x0f, 0xb4, 0x6f, 0xa0, 0xcc, 0xf0, 0x02, 0x4a, 0x79, 0xc3, 0xde,
    0xa3, 0xef, 0xea, 0x51, 0xe6, 0x6b, 0x18, 0xec, 0x1b, 0x2c, 0x80, 0xf7, 0x74, 0xe7, 0xff, 0x21,
    0x5a, 0x6a, 0x54, 0x1e, 0x41, 0x31, 0x92, 0x35, 0xc4, 0x33, 0x07, 0x0a, 0xba, 0x7e, 0x0e, 0x34,
    0x88, 0xb1, 0x98, 0x7c, 0xf3, 0x3d, 0x60, 0x6c, 0x7b, 0xca, 0xd3, 0x1f, 0x32, 0x65, 0x04, 0x28,
    0x64, 0xbe, 0x85, 0x9b, 0x2f, 0x59, 0x8a, 0xd7, 0xb0, 0x25, 0xac, 0xaf, 0x12, 0x03, 0xe2, 0xf2
];

const d = [
    0x22, 0x2f, 0x24, 0x2a,
    0x6d, 0x40, 0x40, 0x40,
    0x40, 0x40, 0x40, 0x40,
    0x40, 0x52, 0x10, 0x30,
];

function S(a) {
    const x = new Uint8Array(4);
    const y = new Uint8Array(4);
    x[0] = a >> 24;
    x[1] = (a >> 16) & 0xff;
    x[2] = (a >> 8) & 0xff;
    x[3] = a & 0xff;

    for (let i = 0; i < 4; i++) {
        y[i] = (i === 0 || i === 2) ? S0[x[i]] : S1[x[i]];
    }
    return (y[0] << 24) | (y[1] << 16) | (y[2] << 8) | y[3];
}

function rot(x, move) {
    return ((x << move) | (x >>> (32 - move))) >>> 0;
}

function L1(x) {
    return (x ^ rot(x, 2) ^ rot(x, 10) ^ rot(x, 18) ^ rot(x, 24)) >>> 0;
}

function L2(x) {
    return (x ^ rot(x, 8) ^ rot(x, 14) ^ rot(x, 22) ^ rot(x, 30)) >>> 0;
}

function modAdd(a, b) {
    let c = (a + b) >>> 0;
    c = (c & 0x7fffffff) + (c >>> 31);
    return c >>> 0;
}

function mod2ExpMul(x, exp) {
    return ((x << exp) | (x >>> (31 - exp))) & 0x7fffffff >>> 0;
}

class ZUCState {
    constructor() {
        this.key = new Uint8Array(32);
        this.iv = new Uint8Array(25);
        this.LFSR = new Uint32Array(16);
        this.X = new Uint32Array(4);
        this.R1 = 0;
        this.R2 = 0;
        this.W = 0;
    }

    printState() {
        // // Print LFSR
        // let lfsrOutput = "LFSR: " + Array.from(this.LFSR).map(num => num.toString(16).padStart(8, '0')).join(" ");
        // console.log(lfsrOutput);

        // // Print X
        // let xOutput = "X: " + Array.from(this.X).map(num => num.toString(16).padStart(8, '0')).join(" ");
        // console.log(xOutput);

        // // Print R1, R2, W
        // console.log("R1: " + this.R1.toString(16).padStart(8, '0'));
        // console.log("R2: " + this.R2.toString(16).padStart(8, '0'));
        // console.log("W: " + this.W.toString(16).padStart(8, '0'));
        // console.log(""); // For extra newline


        // 打印 LFSR
        let lfsrOutput = "LFSR: ";
        for (let i = 0; i < 16; i++) {
            lfsrOutput += this.LFSR[i].toString(2).padStart(32, '0') + " ";
        }
        console.log(lfsrOutput);

        // 打印 X
        let xOutput = "X: ";
        for (let i = 0; i < 4; i++) {
            xOutput += this.X[i].toString(2).padStart(32, '0') + " ";
        }
        console.log(xOutput);

        // 打印 R1, R2, W
        console.log("R1: " + this.R1.toString(2).padStart(32, '0'));
        console.log("R2: " + this.R2.toString(2).padStart(32, '0'));
        console.log("W: " + this.W.toString(2).padStart(32, '0'));
        console.log(""); // 输出额外的换行
    }

    bitReconstruction() {
        const X = new Uint32Array(4);
        X[0] = ((this.LFSR[15] & 0x7fff8000) << 1) | (this.LFSR[14] & 0xffff);
        X[1] = (this.LFSR[11] << 16) | (this.LFSR[9] >> 15);
        X[2] = (this.LFSR[7] << 16) | (this.LFSR[5] >> 15);
        X[3] = (this.LFSR[2] << 16) | (this.LFSR[0] >> 15);
        return X;
    }

    F() {
        const W1 = (this.R1 + this.X[1]) >>> 0;
        const W2 = (this.R2 ^ this.X[2]) >>> 0;
        this.W = (this.X[0] ^ this.R1) + this.R2 >>> 0;

        this.R1 = S(L1((W1 << 16) | (W2 >>> 16)));
        this.R2 = S(L2((W2 << 16) | (W1 >>> 16)));
    }

    keyIVInsert() {
        this.LFSR[0] = (this.key[0] << 23) | (d[0] << 16) | (this.key[21] << 8) | (this.key[16]);
        this.LFSR[1] = (this.key[1] << 23) | (d[1] << 16) | (this.key[22] << 8) | (this.key[17]);
        this.LFSR[2] = (this.key[2] << 23) | (d[2] << 16) | (this.key[23] << 8) | (this.key[18]);
        this.LFSR[3] = (this.key[3] << 23) | (d[3] << 16) | (this.key[24] << 8) | (this.key[19]);
        this.LFSR[4] = (this.key[4] << 23) | (d[4] << 16) | (this.key[25] << 8) | (this.key[20]);

        this.LFSR[5] = (this.iv[0] << 23) | ((d[5] | this.iv[17]) << 16) | (this.key[5] << 8) | (this.key[26]);
        this.LFSR[6] = (this.iv[1] << 23) | ((d[6] | this.iv[18]) << 16) | (this.key[6] << 8) | (this.key[27]);
        this.LFSR[7] = (this.iv[10] << 23) | ((d[7] | this.iv[19]) << 16) | (this.key[7] << 8) | (this.iv[2]);
        this.LFSR[8] = (this.key[8] << 23) | ((d[8] | this.iv[20]) << 16) | (this.iv[3] << 8) | (this.iv[11]);
        this.LFSR[9] = (this.key[9] << 23) | ((d[9] | this.iv[20]) << 16) | (this.iv[12] << 8) | (this.iv[4]);

        this.LFSR[10] = (this.iv[5] << 23) | ((d[10] | this.iv[22]) << 16) | (this.key[10] << 8) | (this.key[28]);
        this.LFSR[11] = (this.key[11] << 23) | ((d[11] | this.iv[23]) << 16) | (this.iv[6] << 8) | (this.iv[13]);
        this.LFSR[12] = (this.key[12] << 23) | ((d[12] | this.iv[24]) << 16) | (this.iv[7] << 8) | (this.iv[14]);

        this.LFSR[13] = (this.key[13] << 23) | (d[13] << 16) | (this.iv[15] << 8) | (this.iv[8]);
        this.LFSR[14] = (this.key[14] << 23) | ((d[14] | (this.key[31] >> 4)) << 16) | (this.iv[16] << 8) | (this.iv[9]);
        this.LFSR[15] = (this.key[15] << 23) | ((d[15] | (this.key[31] & 0xf)) << 16) | (this.iv[16] << 8) | (this.iv[9]);
    }

    LFSRWithInitMode(u) {
        let v = this.LFSR[0];
        let tmp = mod2ExpMul(this.LFSR[0], 8);
        v = modAdd(v, tmp);
        tmp = mod2ExpMul(this.LFSR[4], 20);
        v = modAdd(v, tmp);
        tmp = mod2ExpMul(this.LFSR[10], 21);
        v = modAdd(v, tmp);
        tmp = mod2ExpMul(this.LFSR[13], 17);
        v = modAdd(v, tmp);
        tmp = mod2ExpMul(this.LFSR[15], 15);
        v = modAdd(v, tmp);

        if (v === 0) {
            v = 0x7fffffff;
        }

        let s16 = modAdd(v, u);
        if (s16 === 0) {
            s16 = 0x7fffffff;
        }

        for (let i = 0; i < 15; i++) {
            this.LFSR[i] = this.LFSR[i + 1];
        }
        this.LFSR[15] = s16;
    }

    LFSRWithWorkMode() {
        let s16 = this.LFSR[0];
        let tmp = mod2ExpMul(this.LFSR[0], 8);
        s16 = modAdd(s16, tmp);
        tmp = mod2ExpMul(this.LFSR[4], 20);
        s16 = modAdd(s16, tmp);
        tmp = mod2ExpMul(this.LFSR[10], 21);
        s16 = modAdd(s16, tmp);
        tmp = mod2ExpMul(this.LFSR[13], 17);
        s16 = modAdd(s16, tmp);
        tmp = mod2ExpMul(this.LFSR[15], 15);
        s16 = modAdd(s16, tmp);

        if (s16 === 0) {
            s16 = 0x7fffffff;
        }

        for (let i = 0; i < 15; i++) {
            this.LFSR[i] = this.LFSR[i + 1];
        }
        this.LFSR[15] = s16;
    }

    init(k, iv) {
        for (let i = 0; i < 32; i++) {
            this.key[i] = k[i];
        }
        for (let i = 0; i < 25; i++) {
            this.iv[i] = iv[i];
        }
        for (let i = 0; i < 16; i++) {
            this.LFSR[i] = 0;
        }
        for (let i = 0; i < 4; i++) {
            this.X[i] = 0;
        }
        this.R1 = 0;
        this.R2 = 0;

        this.keyIVInsert();
        // this.printState();

        let u = 0;
        for (let i = 0; i < 32; i++) {
            // console.log("============ round", i, "============");
            this.X = this.bitReconstruction();
            // this.printState();
            this.F();
            // this.printState();
            u = this.W >>> 1;
            this.LFSRWithInitMode(u);
            // this.printState();
        }
        this.X = this.bitReconstruction();
        this.F();
        this.LFSRWithWorkMode();
    }

    keyStreamGenerator(keylen) {
        const keystream = new Uint32Array(keylen);
        for (let i = 0; i < keylen; i++) {
            this.X = this.bitReconstruction();
            this.F();
            keystream[i] = this.W ^ this.X[3];
            this.LFSRWithWorkMode();
        }
        return keystream;
    }
}


// Initialize the state with key and IV as all zeros
const key = new Uint8Array(32).fill(0);
const iv = new Uint8Array(25).fill(0);

const state = new ZUCState();
state.init(key, iv);

// Generate the first 20 bits of keystream
const keylen = 20;
const keystream = state.keyStreamGenerator(keylen);

// Convert keystream to hexadecimal format
const hexKeystream = Array.from(keystream).map(num => num.toString(16).padStart(8, '0')).join('\n');

console.log("Keystream (hex):", hexKeystream);

