package memo

type InputText struct {
	text string
}

func (i *InputText) Append(input string) {
	i.text += input
}

func (i *InputText) GetText() string {
	return i.text
}

func (i *InputText) CreateSnapshot() *Snapshot {
	return &Snapshot{text: i.text}
}

func (i *InputText) RestoreSnapshot(snapshot *Snapshot) {
	i.text = snapshot.text
}

type Snapshot struct {
	text string
}

func (s *Snapshot) GetText() string {
	return s.text
}

type SnapshotHolder struct {
	snapshots []*Snapshot
}

func (s *SnapshotHolder) Push(snapshot *Snapshot) {
	s.snapshots = append(s.snapshots, snapshot)
}

func (s *SnapshotHolder) Pop() *Snapshot {
	if len(s.snapshots) == 0 {
		return nil
	}
	snapshot := s.snapshots[len(s.snapshots)-1]
	s.snapshots = s.snapshots[:len(s.snapshots)-1]
	return snapshot
}
