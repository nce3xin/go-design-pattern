package finitestatemachine

type IMario interface {
	GetName() string
	ObtainMushroom()
	ObtainCap()
	ObtainFireFlower()
	MeetMonster()
}

type SmallMario struct {
	stateMachine *MarioStateMachine
}

func (s *SmallMario) GetName() string {
	return "small"
}

func (s *SmallMario) ObtainMushroom() {
	s.stateMachine.currentState = &SuperMario{stateMachine: s.stateMachine}
	s.stateMachine.score += 100
}

func (s *SmallMario) ObtainCap() {
	s.stateMachine.currentState = &CapMario{stateMachine: s.stateMachine}
	s.stateMachine.score += 200
}

func (s *SmallMario) ObtainFireFlower() {
	s.stateMachine.currentState = &FireMario{stateMachine: s.stateMachine}
	s.stateMachine.score += 300
}

func (s *SmallMario) MeetMonster() {
	// do nothing
}

type MarioStateMachine struct {
	score        int
	currentState IMario
}

func (m *MarioStateMachine) ObtainMushroom() {
	m.currentState.ObtainMushroom()
}

func (m *MarioStateMachine) ObtainCap() {
	m.currentState.ObtainCap()
}

func (m *MarioStateMachine) ObtainFireFlower() {
	m.currentState.ObtainFireFlower()
}

func (m *MarioStateMachine) MeetMonster() {
	m.currentState.MeetMonster()
}

type SuperMario struct {
	stateMachine *MarioStateMachine
}

func (s *SuperMario) GetName() string {
	return "super"
}

func (s *SuperMario) ObtainMushroom() {
	//do nothing
}

func (s *SuperMario) ObtainCap() {
	s.stateMachine.currentState = &CapMario{stateMachine: s.stateMachine}
	s.stateMachine.score += 200
}

func (s *SuperMario) ObtainFireFlower() {
	s.stateMachine.currentState = &FireMario{stateMachine: s.stateMachine}
	s.stateMachine.score += 300
}

func (s *SuperMario) MeetMonster() {
	s.stateMachine.currentState = &SmallMario{stateMachine: s.stateMachine}
	s.stateMachine.score -= 100
}

type CapMario struct {
	stateMachine *MarioStateMachine
}

func (c *CapMario) GetName() string {
	return "cap"
}

func (c *CapMario) ObtainMushroom() {
	//do nothing
}

func (c *CapMario) ObtainCap() {
	//do nothing
}

func (c *CapMario) ObtainFireFlower() {
	//do nothing
}

func (c *CapMario) MeetMonster() {
	c.stateMachine.currentState = &SmallMario{stateMachine: c.stateMachine}
	c.stateMachine.score -= 200
}

type FireMario struct {
	stateMachine *MarioStateMachine
}

func (f *FireMario) GetName() string {
	return "fire"
}

func (f *FireMario) ObtainMushroom() {
	//do nothing
}

func (f *FireMario) ObtainCap() {
	//do nothing
}

func (f *FireMario) ObtainFireFlower() {
	//do nothing
}

func (f *FireMario) MeetMonster() {
	f.stateMachine.currentState = &SmallMario{stateMachine: f.stateMachine}
	f.stateMachine.score -= 300
}

func NewMarioStateMachine(score int) *MarioStateMachine {
	// MarioStateMachine ????????????????????????????????????????????????
	// MarioStateMachine ????????????????????????????????????????????????????????????????????????????????????????????????MarioStateMachine ??????
	// ??????????????????????????????????????????MarioStateMachine ??? score ??? currentState ??????
	m := &MarioStateMachine{
		score: score,
	}
	m.currentState = &SmallMario{stateMachine: m}
	return m
}
