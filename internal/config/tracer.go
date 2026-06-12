package config

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/zeromicro/go-zero/core/logx"
)

type SqlTracer struct{}

func (t *SqlTracer) TraceQueryStart(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	logx.WithContext(ctx).Infof("[SQL] %s | args=%v", data.SQL, data.Args)
	return context.WithValue(ctx, startTimeKey{}, time.Now())
}

func (t *SqlTracer) TraceQueryEnd(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryEndData) {
	start, _ := ctx.Value(startTimeKey{}).(time.Time)
	cost := time.Since(start)

	if data.Err != nil {
		logx.WithContext(ctx).Errorf("[SQL] failed | cost=%s | err=%v", cost, data.Err)
		return
	}
	logx.WithContext(ctx).Infof("[SQL] done | cost=%s | rows=%d", cost, data.CommandTag.RowsAffected())
}

type startTimeKey struct{}
