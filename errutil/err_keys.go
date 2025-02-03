package errutil

type ErrKey string

const (
	FilenameErrKey    ErrKey = "filename"
	FilepathErrKey    ErrKey = "filepath"
	FileSizeErrKey    ErrKey = "file_size"
	FileModeErrKey    ErrKey = "file_mode"
	FileModTimeErrKey ErrKey = "file_mod_time"
	FileInfoErrKey    ErrKey = "file_info"
	BytesReadErrKey   ErrKey = "bytes_read"
	ErrorObjErrKey    ErrKey = "error_obj"
)
