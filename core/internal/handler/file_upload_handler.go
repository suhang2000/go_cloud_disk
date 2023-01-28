package handler

import (
	"crypto/md5"
	"fmt"
	"go_cloud_disk/core/helper"
	"go_cloud_disk/core/models"
	"net/http"
	"path"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_cloud_disk/core/internal/logic"
	"go_cloud_disk/core/internal/svc"
	"go_cloud_disk/core/internal/types"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 1. parse file
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 2. check file if exists
		// 2.1 compute file hash
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		// 2.2 query from MySQL
		rp := new(models.RepositoryPool)
		has, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		if has {
			httpx.OkJsonCtx(r.Context(), w, &types.FileUploadResponse{Identity: rp.Identity})
			return
		}
		// 3. store file into MinIO
		uploadFilePath, err := helper.UploadFile(svcCtx.Minio, file, fileHeader)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 4. pass request info to logic function
		req.Name = fileHeader.Filename
		req.Hash = hash
		req.Path = uploadFilePath
		req.Size = fileHeader.Size
		req.Ext = path.Ext(fileHeader.Filename)

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
