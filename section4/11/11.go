package external

import "context"

// プライベートタイプを宣言
type requestIDKey struct{}

// 外部から requestID を取得するための関数
func GetRequestId(ctx context.Context) (int, bool) {
	// 値取得と型のキャストを行い、値が存在しないか、キャストできない場合は 0, false が返される
	r, ok := ctx.Value(requestIDKey{}).(int)
	if ok {
		return r, true
	}
	return 0, false
}

// 外部から RequestID を保存するための関数
func WithRequestID(ctx context.Context, reqID int) context.Context {
	// パッケージ内で宣言したキーで値を保存
	return context.WithValue(ctx, requestIDKey{}, reqID)
}
