package receivernotification

import (
	"encoding/json"
	"fmt"
	"receiver_siem/diff"
	"receiver_siem/entity/subject"
	"receiver_siem/hash"
)

type FileUpdate struct {
	FileBefore subject.File
	FileAfter  subject.File
	BaseNotification
}

func (f FileUpdate) GetInfo() string {
	answer := ""
	before := string(f.FileBefore.Content)
	after := string(f.FileAfter.Content)
	if before != after {
		answer += fmt.Sprintf("ContentDiff: %s\n", diff.Diff(before, after))
		answer += fmt.Sprintf("Size: %s -> %s\n", f.FileBefore.Size, f.FileAfter.Size)
	}
	if f.FileBefore.Mode != f.FileAfter.Mode {
		answer += fmt.Sprintf("Mode: %s -> %s\n", f.FileBefore.Mode, f.FileAfter.Mode)
	}

	return answer
}

func (f FileUpdate) GetInfoMarkdown() string {
	answer := ""
	before := string(f.FileBefore.Content)
	after := string(f.FileAfter.Content)
	if before != after {
		answer += fmt.Sprintf("*ContentDiff:* %s\n", diff.Diff(before, after))
		answer += fmt.Sprintf("*Size:* %s -> %s\n", f.FileBefore.Size, f.FileAfter.Size)
	}
	if f.FileBefore.Mode != f.FileAfter.Mode {
		answer += fmt.Sprintf("*Mode:* %s -> %s\n", f.FileBefore.Mode, f.FileAfter.Mode)
	}
	return answer
}

func (f FileUpdate) JSON() string {
	bytes, err := json.Marshal(f)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (f FileUpdate) Type() subject.SubjectType {
	return FileChangeT
}

func (f FileUpdate) Name() string {
	return fmt.Sprintf("Файл %s изменён пользователем %s(%s) в процессе %s (%s).",
		f.FileBefore.FullName,
		f.Who.Username, f.Who.Uid,
		f.WhoProcess.PID, f.WhoProcess.NameProcess)
}

func (f FileUpdate) Hash(hash hash.Hash) string {
	return hash(f.JSON())
}
