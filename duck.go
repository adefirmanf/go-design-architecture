package main

import "fmt"

/**
 ! Do you think all Duck can swim? If yes, put that method to Duck interface, it make all duck can swim.
 ! But, all ducks can't do Fly or Quack. I.e Duck Rubber. So, let's keep varying that method to each behaviour
 */
type Duck interface {
	Swim()
	DoFlyWithSomeBehaviour() FlyBehaviour
	DoQuackWithSomeBehaviour() QuackBehaviour
}

/**
 ! FlyBehaviour, what tool / part that used to fly? How about rocket
 */
type FlyBehaviour interface {
	Fly()
}
/**
 ! Still remember Duck Rubber, the quack is different with the real duck right?
 */
type QuackBehaviour interface{
	Quack()
}

type Wings struct{
	Speed float64
}
func (s *Wings) Fly() {
	fmt.Println("Yeay, I'm flying with my wings")
}
func NewFlyWithWings() FlyBehaviour{
	return &Wings{}
}

type Rocket struct{
	Speed float64
}
func (s *Rocket) Fly() {
	fmt.Println("WTF, It's so fast but my butt is hurt")
}

type Squeek struct{}
func (s *Squeek) Quack() {
	fmt.Println("Qua-, I'm mean Squeek!!")
}
func NewQuackWithSqueek() QuackBehaviour{
	return &Squeek{}
}

type QuackQuack struct{}
func (s *QuackQuack) Quack(){
	fmt.Println("Quack Quack... Finally, I'm real duck")
}
func NewQuackQuack() QuackBehaviour{
	return &QuackQuack{}
}
func NewFlyWithRocket() FlyBehaviour{
	return &Rocket{}
}

type RubberDuck struct{}
func (s *RubberDuck) Swim() {
	fmt.Println("Rubber duck can swim")
}
func (s *RubberDuck) DoFlyWithSomeBehaviour() FlyBehaviour{
	return NewFlyWithRocket()
}
func (s *RubberDuck) DoQuackWithSomeBehaviour() QuackBehaviour{
	return NewQuackWithSqueek()
}
func newRubberDuck() Duck{
	return &RubberDuck{}
}

type RealDuck struct{}
func (s *RealDuck) Swim(){
	fmt.Println("Real duck absolutely can swim")
}
func (s *RealDuck) DoFlyWithSomeBehaviour() FlyBehaviour{
	return NewFlyWithWings()
}
func (s *RealDuck) DoQuackWithSomeBehaviour() QuackBehaviour{
	return NewQuackQuack()
}
func NewRealDuck() Duck{
	return &RealDuck{}
}

func playRubberDuck(){
	RubberDuck := newRubberDuck()
	RubberDuck.Swim()

	PerformQuack := RubberDuck.DoQuackWithSomeBehaviour()
	PerformFly := RubberDuck.DoFlyWithSomeBehaviour()

	PerformQuack.Quack()
	PerformFly.Fly()
}

func playRealDuck(){
	RealDuck := NewRealDuck()
	RealDuck.Swim()
	RealDuck.DoFlyWithSomeBehaviour().Fly()
	RealDuck.DoQuackWithSomeBehaviour().Quack()
}

func main(){
	playRubberDuck()
	playRealDuck()
}
