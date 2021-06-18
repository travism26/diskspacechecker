package scanobjects

import "os"

// keep it simple and just using this for illustration dont mean to over engineer it //
type ScanObject struct {
	Path     string
	FileInfo os.FileInfo
}

func (so ScanObject) getPath() string {
	return so.Path
}
func (so ScanObject) getFileInfo() os.FileInfo {
	return so.FileInfo
}
