// StreamCipher 类用于持有密钥和位置
export default class StreamCipher {
    constructor(hexKey) {
        this.hexKey = hexKey;
        this.key = new Uint8Array(32); // 密钥数组长度为32，因为每个十六进制字符代表4位，64个字符代表256位
        this.position = 0;
        this.init(hexKey);
    }

    init(hexKey) {
        if (hexKey.length !== 64) {
            throw new Error('密钥必须是64个字符的十六进制字符串');
        }

        for (let i = 0; i < 64; i += 2) {
            this.key[i / 2] = parseInt(hexKey.slice(i, i + 2), 16);
        }
        this.position = 0;
    }

    // encrypt 方法用于加密明文，输出Base64编码的密文
    encrypt(plainText) {
        if (plainText === '') {
            return {
                cipherText: '',
                position: this.position
            };
        }

        const originPosition = this.position;

        const plainTextBytes = new TextEncoder().encode(plainText);
        const cipherTextBytes = new Uint8Array(plainTextBytes.length);

        for (let i = 0; i < plainTextBytes.length; i++) {
            cipherTextBytes[i] = plainTextBytes[i] ^ this.key[this.position];
            this.position = (this.position + 1) % this.key.length;
        }


        return {
            cipherText: btoa(String.fromCharCode.apply(null, cipherTextBytes)),
            position: originPosition
        };
    }

    // decrypt 方法用于解密，输入为Base64编码的密文，输出为明文字符串
    decrypt(cipherText, position) {
        if (position !== this.position) {
            this.position = position;
        }

        const cipherTextBytes = Uint8Array.from(atob(cipherText), c => c.charCodeAt(0));
        const plainTextBytes = new Uint8Array(cipherTextBytes.length);

        for (let i = 0; i < cipherTextBytes.length; i++) {
            const keyUnit = this.key[this.position];
            const cipherTextUnit = cipherTextBytes[i];
            const plainTextUnit = cipherTextUnit ^ keyUnit;
            plainTextBytes[i] = plainTextUnit;
            this.position = (this.position + 1) % this.key.length;
        }

        return new TextDecoder().decode(plainTextBytes);
    }
}

const defaultKey = '0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF';

// 如果在localStorage中backendPassword存在，则使用backendPassword作为密钥，否则使用默认密钥
const backendPassword = localStorage.getItem('backendPassword');
let userCipherBackend;
if (backendPassword) {
    userCipherBackend = new StreamCipher(backendPassword);
} else {
    userCipherBackend = new StreamCipher(defaultKey);
}

const frontendPassword = localStorage.getItem('frontendPassword');
let userCipherFrontend;
if (frontendPassword) {
    userCipherFrontend = new StreamCipher(frontendPassword);
} else {
    userCipherFrontend = new StreamCipher(defaultKey);
}

const websocketBackendPassword = localStorage.getItem('websocketBackendPassword');
let userCipherWebsocketBackend;
if (websocketBackendPassword) {
    userCipherWebsocketBackend = new StreamCipher(websocketBackendPassword);
} else {
    userCipherWebsocketBackend = new StreamCipher(defaultKey);
}

const websocketFrontendPassword = localStorage.getItem('websocketFrontendPassword');
let userCipherWebsocketFrontend;
if (websocketFrontendPassword) {
    userCipherWebsocketFrontend = new StreamCipher(websocketFrontendPassword);
} else {
    userCipherWebsocketFrontend = new StreamCipher(defaultKey);
}


export { userCipherFrontend, userCipherBackend, userCipherWebsocketBackend, userCipherWebsocketFrontend };