package kugousdk

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/httplib"
	"strconv"
)

const (
	KG_PATH_SEARCH = "http://mobilecdn.kugou.com/api/v3/search/song"
	KG_PATH_DETAIL = "http://m.kugou.com/app/i/getSongInfo.php"
)

type KugouSdk struct {
}

func NewKugouSdk() (kugouSdk *KugouSdk, err error) {
	sdk := &KugouSdk{}
	return sdk, nil
}

func (this *KugouSdk) Search(keyword string, page, size uint64) ([]SearchSongItem, error) {
	req := httplib.Get(KG_PATH_SEARCH)
	req.Param("page", strconv.FormatUint(page, 10))
	req.Param("pagesize", strconv.FormatUint(size, 10))
	req.Param("keyword", keyword)
	byteData, err := req.Bytes()
	if err != nil {
		return nil, errors.New("request to kugou server error: " + err.Error())
	}
	var searchResult SearchResult
	err = json.Unmarshal(byteData, &searchResult)
	if err != nil {
		return nil, errors.New("respond parse as json error: " + err.Error())
	}
	if searchResult.Errcode != 0 {
		return nil, errors.New("kugou server errorCode " + strconv.Itoa(searchResult.Errcode) + ": " + searchResult.Error)
	}
	return searchResult.Data.Info, nil
}
