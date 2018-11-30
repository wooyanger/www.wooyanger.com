package controllers

import (
	"www.wooyanger.com/pkg/tools"
)

const SuccessMsg  = "file upload successful."

type FileController struct {
	Controllers
}

func (f *FileController) PostUpload() {
	file, fileheader, err := f.Ctx.FormFile("file")
	defer file.Close()
	if err != nil {
		f.Ctx.JSON(map[string]interface{}{"code": 1, "state": "fatal", "msg": err})
	}
	if err = tools.CreateImgFile(file, *fileheader); err != nil {
		f.Ctx.JSON(map[string]interface{}{"code": 2, "state": "fatal", "msg": err})
	}
	f.Ctx.JSON(map[string]interface{}{"code": 0, "state": "success", "msg": SuccessMsg,})
}