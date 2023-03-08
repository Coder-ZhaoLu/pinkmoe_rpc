package comment

import (
	"context"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"

	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateComment management
func (l *CreateCommentLogic) CreateComment(in *pinkmoe.CommentInfo) (*pinkmoe.BaseUUIDResp, error) {
	result, err := l.svcCtx.DB.Comment.Create().
		SetContent(in.Content).
		SetServiceUUID(in.ServiceUuid).
		SetStatus(uint8(in.Status)).
		SetUserUUID(in.UserUuid).
		SetType(in.Type).
		Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(err, in)
	}

	return &pinkmoe.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
