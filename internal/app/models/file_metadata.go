package models

type FileMetadata struct {
	FileSha1 string `json:"file_sha1"`
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	FilePath string `json:"file_path"`
	CreateAt string `json:"create_at"`
}

var fileMetadataCollections map[string]FileMetadata

func init() {
	fileMetadataCollections = make(map[string]FileMetadata)
}

func UpdateFileMetadata(metadata FileMetadata) {
	fileMetadataCollections[metadata.FileSha1] = metadata
}

func GetFileMetadata(fileSha1 string) FileMetadata {
	return fileMetadataCollections[fileSha1]
}
