package models

type FileMetadata struct {
	FileSha1 string `json:"file_sha1"`
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	FilePath string `json:"file_path"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}
