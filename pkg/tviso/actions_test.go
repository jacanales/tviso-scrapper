package tviso_test

import (
	`testing`

	`github.com/golang/mock/gomock`
	`github.com/stretchr/testify/assert`

	`tviso-scrapper/pkg/tviso`
	`tviso-scrapper/pkg/tviso/mocks`
)

func TestGetUserCollection_EmptyCollection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rr := mocks.HavingReadRepository(ctrl)
	rr.ArrangeReturnEmptyCollection()

	wr := mocks.HavingWriteRepository(ctrl)
	wr.ArrangeStoreCollectionIsNotCalled()

	err := tviso.GetUserCollection(rr, wr)
	assert.NoError(t, err)
}

func TestGetUserCollection_ReturnCollection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := tviso.Media{ID: 1, Name:"test",}
	mc := havingAMediaCollectionWith(m)
	rr := mocks.HavingReadRepository(ctrl)
	rr.ArrangeReturnCollection(mc)
	rr.ArrangeGetMediaInfo(m)

	wr := mocks.HavingWriteRepository(ctrl)
	wr.ArrangeStoreCollectionIsCalledWith(mc)

	err := tviso.GetUserCollection(rr, wr)
	assert.NoError(t, err)
}

func havingAMediaCollectionWith(media tviso.Media) []tviso.Media {
	return []tviso.Media{media}
}
