package upload

type Upload interface {
	UpFile()
	DownFile()
	ViewFile()
}
