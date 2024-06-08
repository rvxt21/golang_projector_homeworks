package main

import "fmt"

type Character struct {
	Name      string
	Inventory string
}

func (c Character) PrintCharacterInfo() {
	fmt.Printf("Твій персонаж має ім'я %s, твій стартовий інвентар: %s.\n", c.Name, c.Inventory)
}

type Situation struct {
	Description string
	Option1     string
	Next1       *Situation
	Option2     string
	Next2       *Situation
	Option3     string
	Next3       *Situation
}

func NewSituation(description, option1, option2, option3 string, next1, next2, next3 *Situation) *Situation {
	return &Situation{
		Description: description,
		Option1:     option1,
		Next1:       next1,
		Option2:     option2,
		Next2:       next2,
		Option3:     option3,
		Next3:       next3,
	}
}

func NewSituationWithoutOptions(description string) *Situation {
	return &Situation{
		Description: description,
	}
}

type EndSituation struct {
	Description string
}

func (s Situation) PrintSituationWithOptions() {
	fmt.Printf("%s.\nВаріант 1: %s\nВаріант 2: %s\nВаріант 3: %s.\nОбери свій шлях!\n", s.Description, s.Option1, s.Option2, s.Option3)
}

func StartTheGame() Character {
	var name string
	fmt.Println("Введи своє ім'я, Гравець!")
	fmt.Scan(&name)
	character := Character{name, "факел та ніж"}
	character.PrintCharacterInfo()
	return character
}

func GenerateStartSituation(ch Character) *Situation {
	return NewSituation(
		fmt.Sprintf("Ти прокидаєшся в старовинній кімнаті замку. Твій інвентар: %s. Ти бачиш три виходи: двері ліворуч, двері праворуч та таємний хід прямо", ch.Inventory),
		"Зайти в двері ліворуч.",
		"Зайти в двері праворуч",
		"Зайти в таємний хід прямо.",
		nil, nil, nil,
	)
}

func LibraryWayOption() *Situation {
	return NewSituation(
		"Ви потрапляєте у бібліотеку. У бібліотеці мудрець пропонує вам три книги: про магію, про історію замку та про монстрів.",
		"Взяти про магію",
		"Взяти книгу про історію замку",
		"Взяти книгу про монстрів",
		nil, nil, nil,
	)
}

func PaintingRoomOption() *Situation {
	return NewSituation(
		"Ви потрапляєте у залу з картинами. У залі з картинами одна з картин починає рухатися. Ви можете дослідити її, шукати прихований прохід або повернутися назад.",
		"Досліджувати картину.",
		"Шукати прихований прохід.",
		"Повернутися назад.",
		nil, nil, nil,
	)
}

func DungeonWayOption() *Situation {
	return NewSituation(
		"Ви потрапляєте у підземелля. У підземеллі ви зустрічаєте трьох монстрів: дракона, гобліна і привида. Ви повинні вибрати, з ким битися.",
		"Битися з драконом",
		"Битися з гобліном",
		"Битися з привидом",
		nil, nil, nil,
	)
}

func ConnectSituations(startSituation *Situation) {
	libraryWay := LibraryWayOption()
	paintingRoomWay := PaintingRoomOption()
	dungeonWay := DungeonWayOption()

	magicBook := NewSituationWithoutOptions("Ви дізнаєтеся закляття, яке допомагає вам вийти з замку. Ви успішно втекли! Кінець гри.")
	historyBook := NewSituationWithoutOptions("Ви знаходите підказку про прихований вихід і успішно втекли! Кінець гри.")
	monsterBook := NewSituationWithoutOptions("Ви дізнаєтеся слабкі місця монстрів і використовуєте цю інформацію, щоб втекти з замку! Кінець гри.")

	examinePainting := NewSituationWithoutOptions("Ви потрапляєте в секретну кімнату з скарбами. Ви знаходите вихід і успішно тікаєте! Кінець гри.")
	hiddenPassage := NewSituationWithoutOptions("Ви знаходите таємний тунель, що веде на свободу. Ви успішно втекли! Кінець гри.")

	returnToStart := NewSituation(
		"Ви повертаєтеся назад до початку гри.",
		"ліворуч",
		"праворуч",
		"прямо",
		libraryWay, paintingRoomWay, dungeonWay,
	)

	dragonFight := NewSituationWithoutOptions("Ви знайшли спосіб перемогти дракона і втекли з замку! Кінець гри.")
	goblinFight := NewSituationWithoutOptions("Ви вирішили головоломку і перемогли гобліна! Ви втекли з замку! Кінець гри.")
	ghostFight := NewSituationWithoutOptions("Ви знайшли свічку і вигнали привида! Ви втекли з замку! Кінець гри.")

	libraryWay.Next1 = magicBook
	libraryWay.Next2 = historyBook
	libraryWay.Next3 = monsterBook

	paintingRoomWay.Next1 = examinePainting
	paintingRoomWay.Next2 = hiddenPassage
	paintingRoomWay.Next3 = returnToStart

	dungeonWay.Next1 = dragonFight
	dungeonWay.Next2 = goblinFight
	dungeonWay.Next3 = ghostFight

	startSituation.Next1 = libraryWay
	startSituation.Next2 = paintingRoomWay
	startSituation.Next3 = dungeonWay
}

func Game(ch Character) {
	startSituation := GenerateStartSituation(ch)
	ConnectSituations(startSituation)

	currentSituation := startSituation

	for {
		currentSituation.PrintSituationWithOptions()
		var userInputChoice string
		fmt.Scan(&userInputChoice)

		switch userInputChoice {
		case "1":
			if currentSituation.Next1 != nil {
				currentSituation = currentSituation.Next1
			}
		case "2":
			if currentSituation.Next2 != nil {
				currentSituation = currentSituation.Next2
			}
		case "3":
			if currentSituation.Next3 != nil {
				currentSituation = currentSituation.Next3
			}
		default:
			fmt.Println("Ваш вибір не відомий, оберіть інший!")
		}
		if currentSituation.Next1 == nil && currentSituation.Next2 == nil && currentSituation.Next3 == nil {
			fmt.Println(currentSituation.Description)
			fmt.Println("Гра завершена!")
			break
		}
	}
}

func main() {
	character := StartTheGame()
	fmt.Println(character.Name, "починає гру!")
	Game(character)
}
