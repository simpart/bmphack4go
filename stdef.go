/**
 * @file stdef.go
 * @brief bitmap struct define
 * @author simpart
 */
package bmp

import (
	"github.com/simpart/ttrgo-binctl"
)

type FileHeader struct {
	BfType     [2]byte // 2byte
	BfSize     int32   // 6byte
	BfReserved int32   // 10byte
	BfOffBits  int32   // 14byte
}

type InfoHeader struct {
	BcSize         int32 // 18byte
	BcWidth        int32 // 22byte
	BcHeight       int32 // 26byte
	BcPlanes       int16 // 28byte
	BcBitCount     int16 // 30byte
	BiCompression  int32 // 34byte
	BiSizeImage    int32 // 38byte
	BiXPixPerMeter int32 // 42byte
	BiYPixPerMeter int32 // 46byte
	BiClrUsed      int32 // 50byte
	BiCirImportant int32 // 54byte
}

type BmpHeader struct {
	File FileHeader
	Info InfoHeader
}

type Bitmap struct {
	Header BmpHeader
	Binctl *binctl.Target
}
