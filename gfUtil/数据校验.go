package gfUtil

import (
	"context"

	"github.com/gogf/gf/util/gvalid"
)

func E数据校验_检查Map(ctx context.Context, 等待验证的数据 interface{}, 验证规则 interface{}, 提示消息 ...gvalid.CustomMsg) gvalid.Error {
	return gvalid.CheckMap(ctx, 等待验证的数据, 验证规则, 提示消息...)
}
func E数据校验_检查Struct(ctx context.Context, 等待验证的数据 interface{}, 验证规则 interface{}, 提示消息 ...gvalid.CustomMsg) gvalid.Error {
	return gvalid.CheckStruct(ctx, 等待验证的数据, 验证规则, 提示消息...)
}
func E数据校验_检查(ctx context.Context, 等待验证的数据 interface{}, 验证规则 string, 提示消息 interface{}, 参数 ...interface{}) gvalid.Error {
	return gvalid.CheckValue(ctx, 等待验证的数据, 验证规则, 提示消息, 参数...)
}
