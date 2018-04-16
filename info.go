package bmp

func (bmp Bitmap) GetImgSize() int32 {
	return bmp.Header.Info.BiSizeImage
}

func (bmp Bitmap) GetBitCnt() int16 {
	return bmp.Header.Info.BcBitCount
}

func (bmp Bitmap) GetWidth() int32 {
	return bmp.Header.Info.BcWidth
}

func (bmp Bitmap) GetHeight() int32 {
	return bmp.Header.Info.BcHeight
}
