package dto

type NodeDto struct {
	IsDirectory bool
	IsFile      bool
	Name        string
	Template    []byte
	Parent      *[]NodeDto
}
