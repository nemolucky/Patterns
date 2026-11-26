package memento

type editor struct {
	text string
}

func NewEditor() *editor {
	return &editor{}
}

func (e *editor) SetText(text string) {
	e.text = text
}

func (e *editor) Text() string {
	return e.text
}

func (e *editor) Save() *snapshot {
	return newSnapShot(e.text)
}

func (e *editor) Restore(s *snapshot) {
	if s == nil {
		return
	}

	e.text = s.text
}

type snapshot struct {
	text string
}

func newSnapShot(text string) *snapshot {
	return &snapshot{text: text}
}

type history struct {
	snapshots []*snapshot
}

func NewHistory() *history {
	return &history{}
}

func (h *history) Push(s *snapshot) {
	if s == nil {
		return
	}

	h.snapshots = append(h.snapshots, s)
}

func (h *history) Pop() (*snapshot, bool) {
	if len(h.snapshots) == 0 {
		return nil, false
	}

	pop := h.snapshots[len(h.snapshots)-1]
	h.snapshots = h.snapshots[:len(h.snapshots)-1]

	return pop, true
}
