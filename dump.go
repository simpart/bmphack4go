package bmp

import (
	"fmt"
	"github.com/simpart/ttrgo-binctl"
	"runtime"
	"unsafe"
)

func DumpFile(pth string) error {
	bmp, err := getObj(pth)
	if err != nil {
		fmt.Println(runtime.Caller(1))
		return err
	}
	DumpObj(&(bmp.Header))
	return nil
}

func DumpObj(hdr *BmpHeader) {
	fmt.Println("Dump Bitmap File")
	fmt.Println("==========================")
	fmt.Println("# File Header")
	fmt.Print(
		" bfType     : " + string(hdr.File.BfType[0]) + string(hdr.File.BfType[1]) + "\n",
	)
	fmt.Printf(" bfSize     : %d byte\n", hdr.File.BfSize)
	fmt.Printf(" bfReserved : %d\n", hdr.File.BfReserved)
	fmt.Printf(" bfOffBits  : %d byte\n", hdr.File.BfOffBits)
	fmt.Println("")
	fmt.Println("# Info Header")
	fmt.Printf(" bcSize         : %d byte\n", hdr.Info.BcSize)
	fmt.Printf(" bcWidth        : %d px\n", hdr.Info.BcWidth)
	fmt.Printf(" bcHeight       : %d px \n", hdr.Info.BcHeight)
	fmt.Printf(" bcPlanes       : %d\n", hdr.Info.BcPlanes)
	fmt.Printf(" bcBitCount     : %d bit\n", hdr.Info.BcBitCount)
	fmt.Printf(" biCompression  : %d\n", hdr.Info.BiCompression)
	fmt.Printf(" biSizeImage    : %d byte\n", hdr.Info.BiSizeImage)
	fmt.Printf(" biXPixPerMeter : %d\n", hdr.Info.BiXPixPerMeter)
	fmt.Printf(" biYPixPerMeter : %d\n", hdr.Info.BiYPixPerMeter)
	fmt.Printf(" biClrUsed      : %d\n", hdr.Info.BiClrUsed)
	fmt.Printf(" biCirImportant : %d\n", hdr.Info.BiCirImportant)
}

func getObj(pth string) (*Bitmap, error) {
	var err error
	ret_obj := new(Bitmap)

	// load bitmap file
	ret_obj.Binctl, err = binctl.NewReader(pth)
	if err != nil {
		fmt.Println(runtime.Caller(1))
		return nil, err
	}

	// read header
	err = ret_obj.Binctl.Read(
		&ret_obj.Header,
		int(unsafe.Sizeof(BmpHeader{})),
	)
	if err != nil {
		fmt.Println(runtime.Caller(1))
		return nil, err
	}
	return ret_obj, nil
}
