// TorrentParser包用于解析torrent文件
package TorrentParser

import (
	"github.com/jackpal/bencode-go"
	"io"
	"time"
)

type fileInfo interface {
}

type fileInMultiple struct {
	length int64
	path   []string
}

type multipleFileInfo struct {
	name        string
	pieceLength int64
	pieces      []byte
	private     bool
	files       []fileInMultiple
}

type singleFileInfo struct {
	name        string
	pieceLength int64
	pieces      []byte
	private     bool
	length      int64
}

// 代表Torrent信息的结构体
type TorrentInfo struct {
	Announce     string    // tracker地址
	CreatedBy    string    // 创建者
	CreationDate time.Time // 创建时间
	Encoding     string    // 编码
	Info         fileInfo  // 文件信息（单个或多个文件均可）
	UrlList      []string  // url列表
}

func newFileInMultiple(fileInfo map[string]interface{}) fileInMultiple {
	var result fileInMultiple
	result.length = fileInfo["length"].(int64)
	paths := fileInfo["path"].([]interface{})
	result.path = make([]string, 0)
	for _, pathPart := range paths {
		result.path = append(result.path, pathPart.(string))
	}
	return result
}

func newMultipleFileInfo(info map[string]interface{}) multipleFileInfo {
	var result multipleFileInfo
	result.name = info["name"].(string)
	result.pieceLength = info["piece length"].(int64)
	result.pieces = []byte(info["pieces"].(string))
	result.files = make([]fileInMultiple, 0)
	private, hasPrivateProperty := info["private"].(bool)
	if !hasPrivateProperty {
		private = false
	}
	result.private = private
	files := info["files"].([]interface{})
	for _, file := range files {
		result.files = append(result.files, newFileInMultiple(file.(map[string]interface{})))
	}
	return result
}

func newSingleFileInfo(info map[string]interface{}) singleFileInfo {
	var result singleFileInfo
	result.name = info["name"].(string)
	result.pieceLength = info["piece length"].(int64)
	result.pieces = []byte(info["pieces"].(string))
	result.length = info["length"].(int64)
	private, hasPrivateProperty := info["private"].(bool)
	if !hasPrivateProperty {
		private = false
	}
	result.private = private
	return result
}

func newFileInfo(info map[string]interface{}) fileInfo {
	_, hasKeyFiles := info["files"]
	if hasKeyFiles {
		return newMultipleFileInfo(info)
	} else {
		return newSingleFileInfo(info)
	}
}

// 从代表torrent文件的io Reader中构建TorrentInfo对象
func NewTorrentInfo(reader io.Reader) (result TorrentInfo, err error) {
	decoded, err := bencode.Decode(reader)
	if err != nil {
		return
	}
	decodedMap, _ := decoded.(map[string]interface{})
	result.UrlList = make([]string, 0)
	for key := range decodedMap {
		switch key {
		case "announce":
			result.Announce = decodedMap["announce"].(string)
		case "created by":
			result.CreatedBy = decodedMap["created by"].(string)
		case "creation date":
			result.CreationDate = time.Unix(decodedMap["creation date"].(int64), 0)
		case "encoding":
			result.Encoding = decodedMap["encoding"].(string)
		case "info":
			result.Info = newFileInfo(decodedMap["info"].(map[string]interface{}))
		case "url-list":
			result.UrlList = make([]string, 0)
			urlList := decodedMap["url-list"].([]interface{})
			for _, url := range urlList {
				result.UrlList = append(result.UrlList, url.(string))
			}
		}
	}
	return
}

func marshalFileInfo(info fileInfo) map[string]interface{} {
	result := make(map[string]interface{})
	switch info.(type) {
	case singleFileInfo:
		infoWithType := info.(singleFileInfo)
		result["name"] = infoWithType.name
		result["piece length"] = infoWithType.pieceLength
		result["pieces"] = infoWithType.pieces
		if infoWithType.private {
			result["private"] = 0
		} else {
			result["private"] = 1
		}
		result["length"] = infoWithType.length
	case multipleFileInfo:
		infoWithType := info.(multipleFileInfo)
		result["name"] = infoWithType.name
		result["piece length"] = infoWithType.pieceLength
		result["pieces"] = infoWithType.pieces
		if infoWithType.private {
			result["private"] = 0
		} else {
			result["private"] = 1
		}
		result["files"] = infoWithType.files
	}
	return result
}

// 将TorrentInfo对象序列化到io Writer代表的torrent文件中
func (theTorrentInfo TorrentInfo) Marshal(writer io.Writer) error {
	data := make(map[string]interface{})
	data["announce"] = theTorrentInfo.Announce
	data["created by"] = theTorrentInfo.CreatedBy
	data["creation date"] = theTorrentInfo.CreationDate.Unix()
	data["encoding"] = theTorrentInfo.Encoding
	data["url-list"] = theTorrentInfo.UrlList
	data["info"] = marshalFileInfo(theTorrentInfo.Info)
	return bencode.Marshal(writer, data)
}
