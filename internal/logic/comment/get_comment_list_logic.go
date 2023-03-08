package comment

import (
	"context"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/comment"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"

	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentListLogic) GetCommentList(in *pinkmoe.CommentListReq) (*pinkmoe.CommentListResp, error) {
	var predicates []predicate.Comment
	if in.Content != "" {
		predicates = append(predicates, comment.ContentContains(in.Content))
	}
	result, err := l.svcCtx.DB.Comment.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, errorhandler.DefaultEntError(err, in)
	}

	resp := &pinkmoe.CommentListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &pinkmoe.CommentInfo{
			Id:          v.ID.String(),
			CreatedAt:   v.CreatedAt.UnixMilli(),
			Status:      uint32(v.Status),
			Content:     v.Content,
			ServiceUuid: v.ServiceUUID,
			UserUuid:    v.UserUUID,
			Type:        v.Type,
		})
	}

	return resp, nil
}
