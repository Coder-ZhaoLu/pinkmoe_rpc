// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/category"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/comment"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/service"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/sitemeta"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/version"
)

const errInvalidPage = "INVALID_PAGE"

const (
	listField     = "list"
	pageNumField  = "pageNum"
	pageSizeField = "pageSize"
)

type PageDetails struct {
	Page  uint64 `json:"page"`
	Size  uint64 `json:"size"`
	Total uint64 `json:"total"`
}

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

const errInvalidPagination = "INVALID_PAGINATION"

type CategoryPager struct {
	Order  OrderFunc
	Filter func(*CategoryQuery) (*CategoryQuery, error)
}

// CategoryPaginateOption enables pagination customization.
type CategoryPaginateOption func(*CategoryPager)

// DefaultCategoryOrder is the default ordering of Category.
var DefaultCategoryOrder = Desc(category.FieldID)

func newCategoryPager(opts []CategoryPaginateOption) (*CategoryPager, error) {
	pager := &CategoryPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultCategoryOrder
	}
	return pager, nil
}

func (p *CategoryPager) ApplyFilter(query *CategoryQuery) (*CategoryQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// CategoryPageList is Category PageList result.
type CategoryPageList struct {
	List        []*Category  `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (c *CategoryQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...CategoryPaginateOption,
) (*CategoryPageList, error) {

	pager, err := newCategoryPager(opts)
	if err != nil {
		return nil, err
	}

	if c, err = pager.ApplyFilter(c); err != nil {
		return nil, err
	}

	ret := &CategoryPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := c.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		c = c.Order(pager.Order)
	} else {
		c = c.Order(DefaultCategoryOrder)
	}

	c = c.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := c.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type CommentPager struct {
	Order  OrderFunc
	Filter func(*CommentQuery) (*CommentQuery, error)
}

// CommentPaginateOption enables pagination customization.
type CommentPaginateOption func(*CommentPager)

// DefaultCommentOrder is the default ordering of Comment.
var DefaultCommentOrder = Desc(comment.FieldID)

func newCommentPager(opts []CommentPaginateOption) (*CommentPager, error) {
	pager := &CommentPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultCommentOrder
	}
	return pager, nil
}

func (p *CommentPager) ApplyFilter(query *CommentQuery) (*CommentQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// CommentPageList is Comment PageList result.
type CommentPageList struct {
	List        []*Comment   `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (c *CommentQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...CommentPaginateOption,
) (*CommentPageList, error) {

	pager, err := newCommentPager(opts)
	if err != nil {
		return nil, err
	}

	if c, err = pager.ApplyFilter(c); err != nil {
		return nil, err
	}

	ret := &CommentPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := c.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		c = c.Order(pager.Order)
	} else {
		c = c.Order(DefaultCommentOrder)
	}

	c = c.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := c.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type ServicePager struct {
	Order  OrderFunc
	Filter func(*ServiceQuery) (*ServiceQuery, error)
}

// ServicePaginateOption enables pagination customization.
type ServicePaginateOption func(*ServicePager)

// DefaultServiceOrder is the default ordering of Service.
var DefaultServiceOrder = Desc(service.FieldID)

func newServicePager(opts []ServicePaginateOption) (*ServicePager, error) {
	pager := &ServicePager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultServiceOrder
	}
	return pager, nil
}

func (p *ServicePager) ApplyFilter(query *ServiceQuery) (*ServiceQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// ServicePageList is Service PageList result.
type ServicePageList struct {
	List        []*Service   `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (s *ServiceQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...ServicePaginateOption,
) (*ServicePageList, error) {

	pager, err := newServicePager(opts)
	if err != nil {
		return nil, err
	}

	if s, err = pager.ApplyFilter(s); err != nil {
		return nil, err
	}

	ret := &ServicePageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := s.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		s = s.Order(pager.Order)
	} else {
		s = s.Order(DefaultServiceOrder)
	}

	s = s.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := s.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type SitemetaPager struct {
	Order  OrderFunc
	Filter func(*SitemetaQuery) (*SitemetaQuery, error)
}

// SitemetaPaginateOption enables pagination customization.
type SitemetaPaginateOption func(*SitemetaPager)

// DefaultSitemetaOrder is the default ordering of Sitemeta.
var DefaultSitemetaOrder = Desc(sitemeta.FieldID)

func newSitemetaPager(opts []SitemetaPaginateOption) (*SitemetaPager, error) {
	pager := &SitemetaPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultSitemetaOrder
	}
	return pager, nil
}

func (p *SitemetaPager) ApplyFilter(query *SitemetaQuery) (*SitemetaQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// SitemetaPageList is Sitemeta PageList result.
type SitemetaPageList struct {
	List        []*Sitemeta  `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (s *SitemetaQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...SitemetaPaginateOption,
) (*SitemetaPageList, error) {

	pager, err := newSitemetaPager(opts)
	if err != nil {
		return nil, err
	}

	if s, err = pager.ApplyFilter(s); err != nil {
		return nil, err
	}

	ret := &SitemetaPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := s.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		s = s.Order(pager.Order)
	} else {
		s = s.Order(DefaultSitemetaOrder)
	}

	s = s.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := s.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type VersionPager struct {
	Order  OrderFunc
	Filter func(*VersionQuery) (*VersionQuery, error)
}

// VersionPaginateOption enables pagination customization.
type VersionPaginateOption func(*VersionPager)

// DefaultVersionOrder is the default ordering of Version.
var DefaultVersionOrder = Desc(version.FieldID)

func newVersionPager(opts []VersionPaginateOption) (*VersionPager, error) {
	pager := &VersionPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultVersionOrder
	}
	return pager, nil
}

func (p *VersionPager) ApplyFilter(query *VersionQuery) (*VersionQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// VersionPageList is Version PageList result.
type VersionPageList struct {
	List        []*Version   `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (v *VersionQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...VersionPaginateOption,
) (*VersionPageList, error) {

	pager, err := newVersionPager(opts)
	if err != nil {
		return nil, err
	}

	if v, err = pager.ApplyFilter(v); err != nil {
		return nil, err
	}

	ret := &VersionPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := v.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		v = v.Order(pager.Order)
	} else {
		v = v.Order(DefaultVersionOrder)
	}

	v = v.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := v.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}
