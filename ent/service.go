// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/service"
	uuid "github.com/gofrs/uuid/v5"
)

// Service is the model entity for the Service schema.
type Service struct {
	config `json:"-"`
	// ID of the ent.
	// UUID
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// status 1 normal 2 ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// Title | 标题
	Title string `json:"title,omitempty"`
	// Desc | 描述
	Desc string `json:"desc,omitempty"`
	// Content | 内容
	Content string `json:"content,omitempty"`
	// CategoryId | 分类ID
	CategoryID uint64 `json:"category_id,omitempty"`
	// AuthorUuid | 作者ID
	AuthorUUID string `json:"author_uuid,omitempty"`
	// Cover | 封面
	Cover string `json:"cover,omitempty"`
	// Document | 文档
	Document string `json:"document,omitempty"`
	// Type | 类型 1: 微服务 2: 组件库 3: 文章
	Type uint32 `json:"type,omitempty"`
	// Price | 价格
	Price uint32 `json:"price,omitempty"`
	// View | 查看数
	View uint64 `json:"view,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Service) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case service.FieldStatus, service.FieldCategoryID, service.FieldType, service.FieldPrice, service.FieldView:
			values[i] = new(sql.NullInt64)
		case service.FieldTitle, service.FieldDesc, service.FieldContent, service.FieldAuthorUUID, service.FieldCover, service.FieldDocument:
			values[i] = new(sql.NullString)
		case service.FieldCreatedAt, service.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case service.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Service", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Service fields.
func (s *Service) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case service.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case service.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case service.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case service.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = uint8(value.Int64)
			}
		case service.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				s.Title = value.String
			}
		case service.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				s.Desc = value.String
			}
		case service.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				s.Content = value.String
			}
		case service.FieldCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field category_id", values[i])
			} else if value.Valid {
				s.CategoryID = uint64(value.Int64)
			}
		case service.FieldAuthorUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author_uuid", values[i])
			} else if value.Valid {
				s.AuthorUUID = value.String
			}
		case service.FieldCover:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cover", values[i])
			} else if value.Valid {
				s.Cover = value.String
			}
		case service.FieldDocument:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field document", values[i])
			} else if value.Valid {
				s.Document = value.String
			}
		case service.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				s.Type = uint32(value.Int64)
			}
		case service.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				s.Price = uint32(value.Int64)
			}
		case service.FieldView:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field view", values[i])
			} else if value.Valid {
				s.View = uint64(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Service.
// Note that you need to call Service.Unwrap() before calling this method if this Service
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Service) Update() *ServiceUpdateOne {
	return NewServiceClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Service entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Service) Unwrap() *Service {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Service is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Service) String() string {
	var builder strings.Builder
	builder.WriteString("Service(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(s.Title)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(s.Desc)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(s.Content)
	builder.WriteString(", ")
	builder.WriteString("category_id=")
	builder.WriteString(fmt.Sprintf("%v", s.CategoryID))
	builder.WriteString(", ")
	builder.WriteString("author_uuid=")
	builder.WriteString(s.AuthorUUID)
	builder.WriteString(", ")
	builder.WriteString("cover=")
	builder.WriteString(s.Cover)
	builder.WriteString(", ")
	builder.WriteString("document=")
	builder.WriteString(s.Document)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", s.Type))
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", s.Price))
	builder.WriteString(", ")
	builder.WriteString("view=")
	builder.WriteString(fmt.Sprintf("%v", s.View))
	builder.WriteByte(')')
	return builder.String()
}

// Services is a parsable slice of Service.
type Services []*Service
