package main

import "fmt"

type ZookeeperName string
type Age uint8
type AnimalName string

type Zookeeper struct {
	Name            ZookeeperName
	Age             Age
	AnimalsCaughted uint8
}

func (z *Zookeeper) caughtAnimal() {
	z.AnimalsCaughted += 1
}

func (z Zookeeper) printZookeeperInfo() {
	fmt.Printf("Zookeeper name is %s, age is %d, there were animals caught: %d.\n", z.Name, z.Age, z.AnimalsCaughted)
}

type Animal struct {
	Species         string //вид
	Habitat         string //середа проживання
	SocialStructure string //соц.структура solitary, pack, herd
	Diet            string //carnivore, herbivore, omnivore
	Weight          float32
}

func (a Animal) printAnimalInfo() {
	fmt.Printf("Species are %s, habitat is %s, social structure is %s, diet is %s, weight is %.2f.\n", a.Species, a.Habitat, a.SocialStructure, a.Diet, a.Weight)
}

type Lion struct {
	Animal
	Name AnimalName
}

func (l Lion) printLionInfo() {
	l.printAnimalInfo()
	fmt.Printf("And name is %s.\n", l.Name)
}

type Panda struct {
	Animal
	Name AnimalName
}

func (p Panda) printPandaInfo() {
	p.printAnimalInfo()
	fmt.Printf("And name is %s.\n", p.Name)
}

type Zebra struct {
	Animal
	Name AnimalName
}

func (z Zebra) printZebraInfo() {
	z.printAnimalInfo()
	fmt.Printf("And name is %s.\n", z.Name)
}

type Gorilla struct {
	Animal
	Name AnimalName
}

func (g Gorilla) printGorillaInfo() {
	g.printAnimalInfo()
	fmt.Printf("And name is %s.\n", g.Name)
}

type Cage struct {
	Number              uint8
	AnimalThatLivesHere *Animal
	AnimalsCount        uint8
}

func (c Cage) printCageInfo() {
	if c.AnimalThatLivesHere != nil {
		fmt.Printf("Cage number is %d, type of animal that lives here: %s, count of animals %d.\n", c.Number, c.AnimalThatLivesHere.Species, c.AnimalsCount)
	} else {
		fmt.Println("Error, we have 0 Animals")
	}
}

func (c *Cage) addAnimalToCage(animalCount uint8) {
	c.AnimalsCount += animalCount
	fmt.Printf("Added %d animals to cage.\n", animalCount)
}

func main() {
	zookeeper := Zookeeper{"Anastasiia", 25, 0}
	zookeeper.caughtAnimal()
	zookeeper.printZookeeperInfo()
	fmt.Println("***************************************************************")
	lion := Lion{Animal: Animal{"Lion", "Savanna", "Pride", "Carnivore", 450.1}, Name: "PiPi"}
	lion.printLionInfo()
	fmt.Println("***************************************************************")
	zebra := Zebra{Animal: Animal{"Zebra", "Grasslands", "Herd", "Herbivore", 1000.5}, Name: "Zizi"}
	zebra.printZebraInfo()
	fmt.Println("***************************************************************")
	panda := Panda{Animal: Animal{"Panda", "Bamboo Forests", "Solitary", "Herbivore", 2300.0}, Name: "Papa"}
	panda.printPandaInfo()
	fmt.Println("***************************************************************")
	cage1 := Cage{1, &lion.Animal, 1}
	cage1.printCageInfo()
	cage1.addAnimalToCage(2)
	cage1.printCageInfo()
}
