package main

import "fmt"
type Character struct {
		C Weapon
}
func (c*Character)Fight() error  {
	return c.C.UseWeapon()
}
type Weapon struct {
	WB WeaponBehaviour
}
func (w*Weapon)UseWeapon() error {
	return w.WB.useWeapon()
}
type WeaponBehaviour interface {
	useWeapon() error
}
type QueenCharacter struct {
	Character
}

func (QueenCharacter) Who() string {
	return string("I am Queen")
}
type KnightCharacter struct {
	Character
}

func (KnightCharacter)Who() string {
	return string("I am Knight")
}

type KingCharacter struct {
	Character
}

func (KingCharacter)Who() string {
	return string("I am King")
}

type TrollCharacter struct {
	Character
}

func (TrollCharacter)Who() string {
	return string("I am Troll")
}

type AxeBehaviour struct {

}

func (AxeBehaviour)useWeapon() error  {
	fmt.Println("Axe is using")
	return nil
}
type BowAndArrowBehaviour struct {

}

func (BowAndArrowBehaviour)useWeapon() error  {
	fmt.Println("Bow and Arrow are using")
	return nil
}
type KnifeBehaviour struct {

}

func (KnifeBehaviour)useWeapon() error  {
	fmt.Println("Knife is using")
	return nil
}
type SwordBehaviour struct {

}

func (SwordBehaviour)useWeapon() error  {
	fmt.Println("Sword is using")
	return nil
}

func main() {
	var king = &KingCharacter{Character{}}
	fmt.Println(KingCharacter{}.Who())//I am King
king.C=Weapon{}//Choose the weapon
king.C.WB=AxeBehaviour{}
king.Fight()//Axe is using
king.C.WB=BowAndArrowBehaviour{}
king.Fight()//Bow and Arrow are using
king.C.WB=KnifeBehaviour{}
king.Fight()//Knife is using
king.C.WB=SwordBehaviour{}
king.Fight()//Sword is using
//create new person
	var troll=&TrollCharacter{Character{}}
	fmt.Println(TrollCharacter{}.Who())//I am Troll
	troll.C=Weapon{}
	troll.C.WB=KnifeBehaviour{}//Knife is using
	troll.Fight()
	troll.C.WB=BowAndArrowBehaviour{}//Bow and Arrow are using
	troll.Fight()
	troll.C.WB=AxeBehaviour{}//Axe is using
	troll.Fight()
	troll.C.WB=SwordBehaviour{}//Sword is using
	troll.Fight()

}
