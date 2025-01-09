package receivernotification

import (
	"encoding/json"
	"fmt"
	"receiver_siem/entity/subject"
	"receiver_siem/hash"
)

type FileNew struct {
	File subject.File
	BaseNotification
}

func (f FileNew) GetInfo() string {
	return fmt.Sprintf("%s: %s;\n%s: %s;\n%s: %d;\n%s: %s;\n%s: %s\n;",
		"FileName", f.File.FullName,
		"Content", f.File.Content,
		"Size", f.File.Size,
		"Mode", f.File.Mode,
		"Modified", f.File.Modified.Format("2006-01-02 15:04:05"))
}

func (f FileNew) GetInfoMarkdown() string {
	return fmt.Sprintf("*%s:* %s;\n*%s:* %s;\n*%s:* %d;\n*%s:* %s;\n*%s:* %s\n;",
		"FileName", f.File.FullName,
		"Content", f.File.Content,
		"Size", f.File.Size,
		"Mode", f.File.Mode,
		"Modified", f.File.Modified.Format("2006-01-02 15:04:05"))
}

func (f FileNew) JSON() string {
	bytes, err := json.Marshal(f)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (f FileNew) Type() subject.SubjectType {
	return FileNewT
}

func (f FileNew) Name() string {
	return fmt.Sprintf("Файл %s создан %s(%s) в процессе %s (%s).",
		f.File.FullName,
		f.Who.Username, f.Who.Uid,
		f.WhoProcess.PID, f.WhoProcess.NameProcess)
}

func (f FileNew) Hash(hash hash.Hash) string {
	return hash(f.JSON())
}
