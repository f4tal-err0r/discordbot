package music

import (
	"fmt"
	"context"

	"google.golang.org/api/youtube/v3"
)

type YTServ *youtube.Service

func (YTServ *YTServ) Search(entry string) string {
	
}