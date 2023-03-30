package output

import "myapp/entities"

// ユースケースからの出力境界インターフェースを定義
type GetUserOutputBoundary interface {
	Present(user entities.User) (GetUserOutputData, error)
}
