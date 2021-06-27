// @Title  qrcode.go
// @Description  zap日志创建，tools.CreateQrcode() 创建二维码
// @Author  amberhu  20210625
// @Update
package qrcode

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"os"
)

func CreateQrcode(width int, height int, info string, qrname string) string {

	var filecreateinfo = "./statics/images/qrcode/" + qrname + ".png"
	// Create the barcode
	qrCode, _ := qr.Encode(info, qr.M, qr.Auto)
	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, width, height)
	// create the output file
	file, _ := os.Create(filecreateinfo)
	defer file.Close()
	// encode the barcode as png
	png.Encode(file, qrCode)

	return filecreateinfo
}

//func SendQrcode(c *gin.Context){
//	var url = CreateQrcode(200,200,"hihihi","2201")
//
//	c.JSON(http.StatusOK, gin.H{
//		"message": "Hello zaplog!"+url,
//	})
//}
