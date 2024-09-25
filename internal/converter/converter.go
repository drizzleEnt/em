package converter

import (
	"errors"

	"github.com/drizzleent/em/internal/model"
	"github.com/gin-gonic/gin"
)

func FromReqToSong(ctx *gin.Context) (*model.Music, error) {
	var music model.Music
	err := ctx.BindJSON(&music)
	if err != nil {
		return nil, err
	}

	if len(music.Group) == 0 {
		return nil, errors.New("group field is requared")
	}

	if len(music.Song) == 0 {
		return nil, errors.New("song field is requared")
	}
	return &music, nil
}
