package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go_cloud_disk/core/models"
	"log"

	"go_cloud_disk/core/internal/svc"
	"go_cloud_disk/core/internal/types"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
)

type CoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoreLogic {
	return &CoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoreLogic) Core(req *types.Request) (resp *types.Response, err error) {
	// request user
	data := make([]*models.UserBasic, 0)
	if err := models.Engine.Find(&data); err != nil {
		log.Println("Get User Failed, ", err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("JSON Parse Failed", err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", " ")
	if err != nil {
		log.Println("JSON Indent Failed, ", err)
	}
	fmt.Println(dst.String())
	// response
	resp = new(types.Response)
	resp.Message = dst.String()
	return
}
