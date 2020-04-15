package mocks

import (
	`github.com/golang/mock/gomock`

	`tviso-scrapper/pkg/tviso`
)

func HavingReadRepository(ctrl *gomock.Controller) *MockReadRepository {
	return NewMockReadRepository(ctrl)
}

func (m *MockReadRepository) ArrangeReturnEmptyCollection() *gomock.Call {
	return m.EXPECT().GetUserCollection().Times(1).Return([]tviso.Media{}, nil)
}

func (m *MockReadRepository) ArrangeReturnCollection(col []tviso.Media) *gomock.Call {
	return m.EXPECT().GetUserCollection().Times(1).Return(col, nil)
}

func (m *MockReadRepository) ArrangeGetMediaInfo(c tviso.Media) *gomock.Call {
	return m.EXPECT().GetMediaInfo(&c).Times(1).Return(nil)
}

func HavingWriteRepository(ctrl *gomock.Controller) *MockWriteRepository {
	return NewMockWriteRepository(ctrl)
}

func (m *MockWriteRepository) ArrangeStoreCollectionIsNotCalled() *gomock.Call {
	return m.EXPECT().StoreCollection(gomock.Any()).Times(0)
}

func (m *MockWriteRepository) ArrangeStoreCollectionIsCalledWith(col []tviso.Media) *gomock.Call {
	return m.EXPECT().StoreCollection(col).Times(1).Return(nil)
}