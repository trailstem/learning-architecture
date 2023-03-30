package input

import "myapp/entities"

// ユースケースへの入力境界インターフェースを定義
type GetUserInputBoundary interface {
	GetUser(input GetUserInputData) (entities.User, error)
}
