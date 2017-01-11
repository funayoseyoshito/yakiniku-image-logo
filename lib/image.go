package lib

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

var (
	mediumLogo image.Image
	largeLogo  image.Image
	originLogo image.Image
)

func init() {

	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	//ロゴ画像の準備
	mediumLogo = getLogoImageByPath(Config.GetMediumLogoPath())
	largeLogo = getLogoImageByPath(Config.GetLargeLogoPath())
	originLogo = getLogoImageByPath(Config.GetOriginLogoPath())
}

//getLogoImageByPath ファイルパスからロゴimageを生成
func getLogoImageByPath(path string) image.Image {
	logoFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		panic("logoファイルを開くことができませんでした。")
	}
	logoImg, _, err := image.Decode(logoFile)
	if err != nil {
		fmt.Println(err)
		panic("logoファイルのデコードに失敗")
	}

	return logoImg
}

//GetLogoImage ロゴ画像のimage.Imageを返却します。
func GetLogoImageByKind(kind int) image.Image {

	switch kind {
	case Config.Cooking.MediumID, Config.Other.MediumID:
		return mediumLogo
	case Config.Cooking.LargeID, Config.Other.LargeID:
		return largeLogo
	case Config.Cooking.OriginID, Config.Other.OriginID:
		return originLogo
	default:
		fmt.Println(kind)
		panic("logo画像の取得に失敗")
	}
}

//GetMixImageRGBA 合成画像を生成する
func GetMixImageRGBA(originImg image.Image, logoImg image.Image) image.Image {
	startPointLogo := image.Point{
		originImg.Bounds().Dx() - logoImg.Bounds().Dx() - 10, originImg.Bounds().Dy() - logoImg.Bounds().Dy() - 10}

	logoRectangle := image.Rectangle{startPointLogo, startPointLogo.Add(logoImg.Bounds().Size())}
	originRectangle := image.Rectangle{image.Point{0, 0}, originImg.Bounds().Size()}

	rgba := image.NewRGBA(originRectangle)
	draw.Draw(rgba, originRectangle, originImg, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, logoRectangle, logoImg, image.Point{0, 0}, draw.Over)

	return rgba
}
