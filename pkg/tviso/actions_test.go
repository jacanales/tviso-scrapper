package tviso_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"tviso-scrapper/pkg/tviso"
	"tviso-scrapper/pkg/tviso/mocks"
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

	m := tviso.Media{
		ID:   1,
		Name: "test",
	}

	mc := havingAMediaCollectionWith(m)
	rr := mocks.HavingReadRepository(ctrl)
	rr.ArrangeReturnCollection(mc)
	rr.ArrangeGetMediaInfo(&m)

	wr := mocks.HavingWriteRepository(ctrl)
	wr.ArrangeStoreCollectionIsCalledWith(mc)

	err := tviso.GetUserCollection(rr, wr)
	assert.NoError(t, err)
}

func TestGetUserCollection_FailsReadingCollection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rr := mocks.HavingReadRepository(ctrl)
	rr.ArrangeGetUserCollectionError()

	wr := mocks.HavingWriteRepository(ctrl)

	err := tviso.GetUserCollection(rr, wr)
	assert.EqualError(t, err, fmt.Errorf("could not get user collection: %w", mocks.ErrGetUserCollectionError).Error())
}

func TestGetUserCollection_FailsGettingMediaInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := tviso.Media{ID: 1, Name: "test"}
	m2 := tviso.Media{ID: 2, Name: "test2"}
	mc := havingAMediaCollectionWith(m, m2)
	rr := mocks.HavingReadRepository(ctrl)
	rr.ArrangeReturnCollection(mc)
	rr.ArrangeGetMediaInfoError(&m)
	rr.ArrangeGetMediaInfoError(&m2)

	wr := mocks.HavingWriteRepository(ctrl)

	err := tviso.GetUserCollection(rr, wr)
	assert.EqualError(t, err, fmt.Sprintf(
		"media: %v, error: %v\nmedia: %v, error: %v\n",
		1,
		mocks.ErrGetMediaInfoError.Error(),
		2,
		mocks.ErrGetMediaInfoError.Error(),
	),
	)
}

func havingAMediaCollectionWith(media ...tviso.Media) []tviso.Media {
	var mc []tviso.Media

	return append(mc, media...)
}
