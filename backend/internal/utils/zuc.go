package utils

/*
#cgo CFLAGS: -I./cmodule
#cgo LDFLAGS: -L./cmodule -lzuc
#include "zuc.h"
*/
import "C"
import (
	"encoding/base64"
	"unsafe"
)

type ZUCState struct {
	State    C.ZUC_State
	Position int
}

// ZUCInit 初始化 ZUC 状态
func ZUCInit(k []byte, iv []byte) *ZUCState {
	if len(k) != 32 || len(iv) != 25 {
		panic("key length must be 32 bytes, iv length must be 16 bytes")
	}

	// 创建 ZUC 状态结构体
	var state ZUCState
	state.Position = 0

	// 转换 Go 切片为 C 指针
	kPtr := (*C.uchar)(unsafe.Pointer(&k[0]))
	ivPtr := (*C.uchar)(unsafe.Pointer(&iv[0]))

	// 调用 C 函数初始化
	C.ZUC_Init(&state.State, kPtr, ivPtr)

	return &state
}

// ZUCNext 生成 keystream
func ZUCNext(state *ZUCState, keystreamLen int) []uint32 {
	keystream := make([]uint32, keystreamLen)
	keystreamPtr := (*C.uint32_t)(unsafe.Pointer(&keystream[0]))

	C.ZUC_Gene(&state.State, keystreamPtr, C.int(keystreamLen))
	return keystream
}

// Encrypt 加密
func (state *ZUCState) Encrypt(plainText []byte) (string, int, error) {
	// 计算密文长度
	cipherText := make([]byte, len(plainText))

	// 需要的密钥流长度是明文长度的四分之一，因为 ZUC 生成的是 32 位的密钥流
	// 向上取整
	originPosition := state.Position
	keystreamLen := (len(plainText) + 3) / 4
	keystream := ZUCNext(state, keystreamLen)

	// 加密
	for i := 0; i < len(plainText); i++ {
		cipherText[i] = plainText[i] ^ byte(keystream[i/4]>>uint(8*(i%4)))
	}
	state.Position = state.Position + keystreamLen

	//返回Base64编码的密文
	return base64.StdEncoding.EncodeToString(cipherText), originPosition, nil

}

// Decrypt 解密
func (state *ZUCState) Decrypt(encryptedText string, wantedPosition int) (string, error) {
	// 计算密文长度
	encryptedTextBytes, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}
	plainText := make([]byte, len(encryptedTextBytes))

	// 需要的密钥流长度是密文长度的四分之一，因为 ZUC 生成的是 32 位的密钥流
	// 向上取整
	keystreamLen := (len(encryptedTextBytes) + 3) / 4
	keystream := ZUCNext(state, keystreamLen)
	// 解密
	for i := 0; i < len(encryptedTextBytes); i++ {
		plainText[i] = encryptedTextBytes[i] ^ byte(keystream[i/4]>>uint(8*(i%4)))
	}
	state.Position = state.Position + keystreamLen

	return string(plainText), nil
}
