package memo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInputText_RestoreSnapshot(t *testing.T) {
	inputText := &InputText{}
	snapshotHolder := &SnapshotHolder{snapshots: make([]*Snapshot, 0)}
	commands := []string{"hello", ":list", "world", ":undo", ":list"}
	for _, command := range commands {
		if command == ":list" {
			t.Log(inputText.GetText())
			assert.Equal(t, "hello", inputText.GetText())
		} else if command == ":undo" {
			snapshot := snapshotHolder.Pop()
			if snapshot != nil {
				inputText.RestoreSnapshot(snapshot)
			}
		} else {
			snapshot := inputText.CreateSnapshot()
			snapshotHolder.Push(snapshot)
			inputText.Append(command)
		}
	}
}
