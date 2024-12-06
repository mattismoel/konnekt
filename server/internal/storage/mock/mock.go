package mock

type mockStorage struct{}

func NewMockStorage() *mockStorage {
	return nil
}

func removeFromSlice[T any](slice []T, idx int) []T {
	slice = append(slice[:idx], slice[idx+1:]...)
	return slice
}
