package token

import (
	"context"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.UUIDReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.MmsRpc.BlockUserAllToken(l.ctx, &mms.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: result.Msg}, nil
}
