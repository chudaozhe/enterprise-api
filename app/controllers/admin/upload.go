package admin

import (
	"enterprise-api/app/config"
	"enterprise-api/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 保存编辑器图片
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "msg": "请选择图片"})
		return
	}
	path, err := models.SaveEditorImage(c, file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "", "url": "/" + path, "file_path": config.GetConfig().AppHost + "/" + path})
	//return ['success'=>true, 'msg'=>'', 'file_path'=>$this->config('host').'/data/upload/'.$path];
	//	//	return ['success'=>false, 'msg'=>'上传失败'];
}
