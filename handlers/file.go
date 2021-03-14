package handlers

import (
	"files_handler/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path/filepath"
)

type File struct {
	l *log.Logger
}

const FileIDKey = "id"
const FileNameKey = "filename"
const FileSystemRoot = "./file_system_root"
func NewFile(l *log.Logger) *File {
	return &File{l: l}
}
func (f *File) Get(writer http.ResponseWriter, request *http.Request)  {

}
func (f *File) Upload(writer http.ResponseWriter, request *http.Request)  {
	fileName := mux.Vars(request)[FileNameKey]
	fileID := mux.Vars(request)[FileIDKey]
	f.l.Printf("filename: %s, fileID: %s", fileName, fileID)
	fp := filepath.Join(fileID, fileName)
	f.l.Printf("fp: %s", fp)
	local, err := storage.NewLocal(FileSystemRoot)
	if err != nil {
		http.Error(writer, "Can't store file", http.StatusInternalServerError)
		return
	}
	local.Save(fp, request.Body)



}
