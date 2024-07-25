package utils

import (
	"encoding/base64"
	"encoding/hex"
)

// StreamCipher struct to hold the Key and Position
type StreamCipher struct {
	Key      []uint32
	Position int
}

// NewStreamCipher constructor to initialize the StreamCipher with a hexadecimal Key
func NewStreamCipher(hexKey string) *StreamCipher {
	key := make([]uint32, 0, 32)
	for i := 0; i < 64; i += 2 {
		parsedKey, _ := hex.DecodeString(hexKey[i : i+2])
		key = append(key, uint32(parsedKey[0]))
	}

	return &StreamCipher{Key: key, Position: 0}
}

// Encrypt 将明文加密为密文
// 输入 plainText: 明文字符串
// 输出：被**Base64编码**的密文字符串，加密前的Position，错误信息
func (sc *StreamCipher) Encrypt(plainText []byte) (string, int, error) {
	originPosition := sc.Position
	plainTextBytes := plainText
	cipherTextBytes := make([]byte, len(plainTextBytes))
	unitSize := 1

	for i := 0; i < len(plainTextBytes); i += unitSize {
		keyUnit := sc.Key[sc.Position]
		plainTextUnit := uint32(plainTextBytes[i])

		cipherTextUnit := plainTextUnit ^ keyUnit

		cipherTextBytes[i] = byte(cipherTextUnit)

		sc.Position = (sc.Position + 1) % len(sc.Key)
	}
	//println("Hex format: ", hex.EncodeToString(cipherTextBytes))
	//println("String format: ", string(cipherTextBytes))
	//println("Base64 encode format: ", base64.StdEncoding.EncodeToString(cipherTextBytes))
	return base64.StdEncoding.EncodeToString(cipherTextBytes), originPosition, nil
}

// Decrypt 将密文解密为明文
// 输入 encryptedText: 被**Base64编码**的密文字符串
// 输入 wantedPosition: 加密前的Position
// 输出：明文字符串，错误信息
func (sc *StreamCipher) Decrypt(encryptedText string, wantedPosition int) (string, error) {
	if len(encryptedText) == 0 {
		return "", nil
	}
	//println("Encrypted text: ", encryptedText)
	//println("wantedPosition: ", wantedPosition)
	//println("sc.Position: ", sc.Position)
	if wantedPosition != sc.Position {
		println("Position changed from ", sc.Position, " to ", wantedPosition)
		sc.Position = wantedPosition
	}

	encryptedTextBytes, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	plainTextBytes := make([]byte, len(encryptedTextBytes))
	unitSize := 1

	for i := 0; i < len(encryptedTextBytes); i += unitSize {
		keyUnit := sc.Key[sc.Position]
		encryptedTextUnit := uint32(encryptedTextBytes[i])

		plainTextUnit := encryptedTextUnit ^ keyUnit

		plainTextBytes[i] = byte(plainTextUnit)

		sc.Position = (sc.Position + 1) % len(sc.Key)
	}

	return string(plainTextBytes), nil
}

func ZUCMac32(key string, data string) string {
	// ZUC算法
	return Sha256OfString(key + data)[0:32]
}
