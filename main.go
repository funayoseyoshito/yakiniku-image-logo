package main

import (
	"fmt"

	"github.com/funayoseyoshito/yakiniku-image-logo/db"
	"github.com/funayoseyoshito/yakiniku-image-logo/lib"
)

func main() {
	var originImage db.Images
	defer db.GetConnection().Close()

	for i := 0; ; {
		rows := db.SelectProcessingRows(i)

		if !rows.Next() {
			break
		}

		for {
			db.GetConnection().ScanRows(rows, &originImage)

			originImg := originImage.GetOriginImage()
			logoImg := lib.GetLogoImageByKind(originImage.Kind)
			mixRGBA := lib.GetMixImageRGBA(originImg, logoImg)

			fmt.Println(originImage.ID)
			originImage.UpdateImage(mixRGBA)
			//fmt.Println(db.InsertDb(mixRGBA))

			if !rows.Next() {
				break
			}
		}
		i += db.SelectLimit
	}
}
