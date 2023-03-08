package comment

import (
	"context"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/comment"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/pkg/uuidx"

	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentByIdLogic {
	return &GetCommentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentByIdLogic) GetCommentById(in *pinkmoe.UUIDReq) (*pinkmoe.CommentInfo, error) {
	result, err := l.svcCtx.DB.Comment.Query().Where(comment.IDEQ(uuidx.ParseUUIDString(in.Id))).First(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(err, in)
	}

	return &pinkmoe.CommentInfo{
		Id:          result.ID.String(),
		CreatedAt:   result.CreatedAt.UnixMilli(),
		UpdatedAt:   result.CreatedAt.UnixMilli(),
		Content:     result.Content,
		Status:      uint32(result.Status),
		ServiceUuid: result.ServiceUUID,
		UserUuid:    result.UserUUID,
		Type:        result.Type,
	}, nil
}
