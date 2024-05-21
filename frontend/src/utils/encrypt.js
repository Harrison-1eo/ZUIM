const crypto = require('crypto');

// StreamCipher 类用于持有密钥和位置
class StreamCipher {
    constructor(hexKey) {
        if (hexKey.length !== 64) {
            throw new Error('密钥必须是64个字符的十六进制字符串');
        }

        this.key = [];
        for (let i = 0; i < 64; i += 2) {
            this.key.push(parseInt(hexKey.slice(i, i + 2), 16));
        }
        this.position = 0;
    }

    // encrypt 方法用于加密明文，输出Base64编码的密文
    encrypt(plainText) {
        const plainTextBytes = Buffer.from(plainText);
        const cipherTextBytes = Buffer.alloc(plainTextBytes.length);

        for (let i = 0; i < plainTextBytes.length; i++) {
            cipherTextBytes[i] = plainTextBytes[i] ^ this.key[this.position];
            this.position = (this.position + 1) % this.key.length;
        }

        return cipherTextBytes.toString('base64');
    }

    // decrypt 方法用于解密，输入为Base64编码的密文，输出为明文字符串
    decrypt(cipherText) {
        const cipherTextBytes = Buffer.from(cipherText, 'base64');
        const plainTextBytes = Buffer.alloc(cipherTextBytes.length);

        for (let i = 0; i < cipherTextBytes.length; i++) {
            const keyUnit = this.key[this.position];
            const cipherTextUnit = cipherTextBytes[i];
            const plainTextUnit = cipherTextUnit ^ keyUnit;
            plainTextBytes[i] = plainTextUnit;
            this.position = (this.position + 1) % this.key.length;
        }

        return plainTextBytes.toString();
    }
}

async function main() {
    const hexKey = '0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF';
    const cipher = new StreamCipher(hexKey);
    const plainText = 'Hello, world!';
    const encryptedText = cipher.encrypt(plainText);

    const cipher2 = new StreamCipher(hexKey);
    const decryptedText = cipher2.decrypt(encryptedText);

    console.log('Encrypted Text:', encryptedText);
    console.log('Decrypted Text:', decryptedText);
}

main();
