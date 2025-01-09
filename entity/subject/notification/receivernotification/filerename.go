package receivernotification

import (
	"fmt"
	"receiver_siem/entity/subject"
)

type FileRename struct {
	FileUpdate
}

func (f FileRename) GetInfo() string {
	return fmt.Sprintf("%s: %s -> %s;",
		"FileName", f.FileBefore.FullName, f.FileAfter.FullName,
	)
}

func (f FileRename) GetInfoMarkdown() string {
	return fmt.Sprintf("*%s:* %s -> %s;",
		"FileName", f.FileBefore.FullName, f.FileAfter.FullName,
	)
}

func (f FileRename) Name() string {
	return fmt.Sprintf("Файл %s переименован в %s  %s(%s) в процессе %s (%s).",
		f.FileBefore.FullName,
		f.FileAfter.FullName,
		f.Who.Username, f.Who.Uid,
		f.WhoProcess.PID, f.WhoProcess.NameProcess)
}

func (f FileRename) Type() subject.SubjectType {
	return FileRenameT
}
