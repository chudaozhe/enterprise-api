package schemas

import "mime/multipart"

type UploadIn struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}
