package diff

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var dmp = diffmatchpatch.New()

func getDiffString(diffs []diffmatchpatch.Diff) string {
	answer := ""
	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffDelete:
			answer += fmt.Sprintf("-%s", diff.Text)
			continue
		case diffmatchpatch.DiffInsert:
			answer += fmt.Sprintf("+%s", diff.Text)
			continue
		case diffmatchpatch.DiffEqual:
			answer += fmt.Sprintf("{%s}", diff.Text)
			continue
		}
	}
	return answer
}

func Diff(before, after string) string {
	diffs := dmp.DiffMain(before, after, false)
	return getDiffString(diffs)
}
