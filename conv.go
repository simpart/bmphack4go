package bmp

import (
	"errors"
	"fmt"
	"github.com/simpart/tetraring4go/binctl"
	"runtime"
	//"unsafe"
)

const white_val = 15
const black_val = 0
const red_val = 9

func ConvBitcnt(src string, dst string, bc int) error {
	// check bit count
	if 4 != bc {
		return errors.New("not supported bit count yet")
	}

	// get bitmap object
	src_bmp, err := getObj(src)
	if err != nil {
		fmt.Println(runtime.Caller(1))
		return err
	}
	// check bitcount
	if 24 != int(src_bmp.GetBitCnt()) {
		return errors.New("source bitmap is not supported bit count")
	}

	// start convert
	read_cnt := int(src_bmp.GetImgSize() / int32(src_bmp.GetBitCnt()))
	dot := make([]byte, 3)
	dst_dot := make([]byte, bc*int(src_bmp.GetWidth())*int(src_bmp.GetHeight())/8)
	var dot_buf byte
	for i := 0; i < read_cnt; i++ {
		src_bmp.Binctl.Read(&dot, 3)
		if (0 != i) && (0 == i%2) {
			dot_buf = convDot(dot)
		} else {
			dst_dot[i/2] = (dot_buf << 4) | convDot(dot)
		}
	}
	// write convert bitmap
	dst_hdr := src_bmp.Header
	dst_hdr.File.BfSize = ((int32(bc) * int32(src_bmp.GetWidth()) * int32(src_bmp.GetHeight())) / 8) + 118
	dst_hdr.File.BfOffBits = 118
	dst_hdr.Info.BcBitCount = 4
	dst_hdr.Info.BiSizeImage = ((int32(bc) * int32(src_bmp.GetWidth()) * int32(src_bmp.GetHeight())) / 8)
	dst_hdr.Info.BiXPixPerMeter = 2835
	dst_hdr.Info.BiYPixPerMeter = 2835
	d_tgt, err := binctl.NewWriter(dst)
	if err != nil {
		fmt.Println(runtime.Caller(1))
		return err
	}
	d_tgt.Write(&dst_hdr)
	// set palette
	dst_plt := make([]uint32, 16)
	dst_plt[0] = 0x00000000  /* palette0 */ /* BLACK */
	dst_plt[1] = 0x00008000  /* palette1 */
	dst_plt[2] = 0x00800000  /* palette2 */
	dst_plt[3] = 0x00808000  /* palette3 */
	dst_plt[4] = 0x80000000  /* palette4 */
	dst_plt[5] = 0x80008000  /* palette5 */
	dst_plt[6] = 0x80800000  /* palette6 */
	dst_plt[7] = 0x80808000  /* palette7 */
	dst_plt[8] = 0xC0C0C000  /* palette8 */
	dst_plt[9] = 0x0000FF00  /* palette9 */ /* RED */
	dst_plt[10] = 0x00FF0000 /* palette10 */
	dst_plt[11] = 0x00FFFF00 /* palette11 */
	dst_plt[12] = 0xFF000000 /* palette12 */
	dst_plt[13] = 0xFF00FF00 /* palette13 */
	dst_plt[14] = 0xFFFF0000 /* palette14 */
	dst_plt[15] = 0xFFFFFF00 /* palette15 */ /* WHITE */
	d_tgt.Write(&dst_plt)
	d_tgt.Write(&dst_dot)

	return nil
}

func convDot(dot []byte) byte {

	blue := dot[0]
	green := dot[1]
	red := dot[2]

	if (200 <= blue) &&
		(200 <= green) &&
		(200 <= red) {
		/* white */
		return white_val
	} else if (100 >= blue) &&
		(100 >= green) &&
		(100 >= red) {
		/* black */
		return black_val
	} else if (50 >= blue) &&
		(50 >= green) &&
		(200 <= red) {
		/* red */
		return red_val
	} else {
		/* other */
		return white_val
	}
}

//imgbmp.ConvBitcnt
