package statelookuptable

const (
	SMALL = 0
	SUPER = 1
	CAP   = 2
	FIRE  = 3

	GotMushroom = 0
	GotCap      = 1
	GotFire     = 2
	MetMonster  = 3
)

var transitionTable = [][]int{
	{SUPER, CAP, FIRE, SMALL},
	{SMALL, CAP, FIRE, SUPER},
	{CAP, CAP, CAP, SMALL},
	{FIRE, FIRE, FIRE, SMALL},
}

var actionTable = [][]int{
	{+100, +200, +300, 0},
	{0, +200, +300, -100},
	{0, 0, 0, -200},
	{0, 0, 0, -300},
}

type MarioStateMachine struct {
	score        int
	currentState int
}

func (m *MarioStateMachine) obtainMushroom() {
	m.executeEvent(GotMushroom)
}

func (m *MarioStateMachine) executeEvent(event int) {
	stateValue := m.currentState
	eventValue := event
	m.currentState = transitionTable[stateValue][eventValue]
	m.score += actionTable[stateValue][eventValue]
}

func (m *MarioStateMachine) obtainCap() {
	m.executeEvent(GotCap)
}

func (m *MarioStateMachine) obtainFireFlower() {
	m.executeEvent(GotFire)
}

func (m *MarioStateMachine) meetMonster() {
	m.executeEvent(MetMonster)
}

func NewMarioStateMachine(score int, currentState int) *MarioStateMachine {
	return &MarioStateMachine{
		score:        score,
		currentState: currentState,
	}
}
