// Code generated by goctl. DO NOT EDIT.

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysApiInfoFieldNames          = builder.RawFieldNames(&SysApiInfo{})
	sysApiInfoRows                = strings.Join(sysApiInfoFieldNames, ",")
	sysApiInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(sysApiInfoFieldNames, "`id`", "`updatedTime`", "`deletedTime`", "`createdTime`"), ",")
	sysApiInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(sysApiInfoFieldNames, "`id`", "`updatedTime`", "`deletedTime`", "`createdTime`"), "=?,") + "=?"
)

type (
	sysApiInfoModel interface {
		Insert(ctx context.Context, data *SysApiInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysApiInfo, error)
		FindOneByRoute(ctx context.Context, route string) (*SysApiInfo, error)
		Update(ctx context.Context, data *SysApiInfo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysApiInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysApiInfo struct {
		Id           int64        `db:"id"`           // 编号
		Route        string       `db:"route"`        // 路由
		Method       int64        `db:"method"`       // 请求方式（1 GET 2 POST 3 HEAD 4 OPTIONS 5 PUT 6 DELETE 7 TRACE 8 CONNECT 9 其它）
		Name         string       `db:"name"`         // 请求名称
		BusinessType int64        `db:"businessType"` // 业务类型（1新增 2修改 3删除 4查询 5其它）
		Group        string       `db:"group"`        // 接口组
		Desc         string       `db:"desc"`         // 备注
		CreatedTime  time.Time    `db:"createdTime"`  // 创建时间
		UpdatedTime  time.Time    `db:"updatedTime"`  // 更新时间
		DeletedTime  sql.NullTime `db:"deletedTime"`
	}
)

func newSysApiInfoModel(conn sqlx.SqlConn) *defaultSysApiInfoModel {
	return &defaultSysApiInfoModel{
		conn:  conn,
		table: "`sys_api_info`",
	}
}

func (m *defaultSysApiInfoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysApiInfoModel) FindOne(ctx context.Context, id int64) (*SysApiInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysApiInfoRows, m.table)
	var resp SysApiInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysApiInfoModel) FindOneByRoute(ctx context.Context, route string) (*SysApiInfo, error) {
	var resp SysApiInfo
	query := fmt.Sprintf("select %s from %s where `route` = ? limit 1", sysApiInfoRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, route)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysApiInfoModel) Insert(ctx context.Context, data *SysApiInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, sysApiInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Route, data.Method, data.Name, data.BusinessType, data.Group, data.Desc)
	return ret, err
}

func (m *defaultSysApiInfoModel) Update(ctx context.Context, newData *SysApiInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysApiInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Route, newData.Method, newData.Name, newData.BusinessType, newData.Group, newData.Desc, newData.Id)
	return err
}

func (m *defaultSysApiInfoModel) tableName() string {
	return m.table
}
