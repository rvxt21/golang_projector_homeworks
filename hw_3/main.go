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

type EndSituation struct {
	Description string
}

func (s Situation) PrintSituationWithOptions() {
	fmt.Printf("%s.\nВаріант 1: %s\nВаріант 2: %s\nВаріант 3: %s.\nОбери свій шлях!\n", s.Description, s.Option1, s.Option2, s.Option3)
}

func StartTheGame() {
	var name string
	fmt.Println("Введи своє ім'я, Гравець!")
	fmt.Scan(&name)
	character := Character{name, "факел та ніж"}
	character.PrintCharacterInfo()
}

func Game() {
	startSituation := Situation{
		Description: "Ти прокидаєшся в старовинній кімнаті замку. Ти бачиш три виходи: двері ліворуч, двері праворуч та таємний хід прямо.",
		Option1:     "Зайти в двері ліворуч.",
		Option2:     "Зайти в двері праворуч",
		Option3:     "Зайти в таємний хід прямо.",
	}

	libraryWay := &Situation{
		Description: "Ви потряпляєте у бібліотеку. У бібліотеці мудрець пропонує вам три книги: про магію, про історію замку та про монстрів.",
		Option1:     "Взяти про магію",
		Option2:     "Взяти книгу про історію замку",
		Option3:     "Взяти книгу про монстрів",
	}

	paintingRoomWay := &Situation{
		Description: "Ви потрапляєте у залу з картинами. У залі з картинами одна з картин починає рухатися. Ви можете дослідити її, шукати прихований прохід або повернутися назад.",
		Option1:     "Досліджувати картину.",
		Option2:     "Шукати прихований прохід.",
		Option3:     "Повернутися назад.",
	}
	dungeonWay := &Situation{
		Description: "Ви потрапляєте у підземелля. У підземеллі ви зустрічаєте трьох монстрів: дракона, гобліна і привида. Ви повинні вибрати, з ким битися.",
		Option1:     "Битися з драконом",
		Option2:     "Битися з гобліном",
		Option3:     "Битися з привидом",
	}

	startSituation.Next1 = libraryWay
	startSituation.Next2 = paintingRoomWay
	startSituation.Next3 = dungeonWay
	magicBook := &Situation{
		Description: "Ви дізнаєтеся закляття, яке допомагає вам вийти з замку. Ви успішно втекли! Кінець гри.",
	}
	historyBook := &Situation{
		Description: "Ви знаходите підказку про прихований вихід і успішно втекли! Кінець гри.",
	}
	monsterBook := &Situation{
		Description: "Ви дізнаєтеся слабкі місця монстрів і використовуєте цю інформацію, щоб втекти з замку! Кінець гри.",
	}

	examinePainting := &Situation{
		Description: "Ви потрапляєте в секретну кімнату з скарбами. Ви знаходите вихід і успішно тікаєте! Кінець гри.",
	}
	hiddenPassage := &Situation{
		Description: "Ви знаходите таємний тунель, що веде на свободу. Ви успішно втекли! Кінець гри.",
	}
	returnToStart := &Situation{
		Description: "Ви повертаєтеся назад до початку гри.",
		Option1:     "ліворуч",
		Option2:     "праворуч",
		Option3:     "прямо",
		Next1:       libraryWay,
		Next2:       paintingRoomWay,
		Next3:       dungeonWay,
	}

	dragonFight := &Situation{
		Description: "Ви знайшли спосіб перемогти дракона і втекли з замку! Кінець гри.",
	}
	goblinFight := &Situation{
		Description: "Ви вирішили головоломку і перемогли гобліна! Ви втекли з замку! Кінець гри.",
	}
	ghostFight := &Situation{
		Description: "Ви знайшли свічку і вигнали привида! Ви втекли з замку! Кінець гри.",
	}

	libraryWay.Next1 = magicBook
	libraryWay.Next2 = historyBook
	libraryWay.Next3 = monsterBook

	paintingRoomWay.Next1 = examinePainting
	paintingRoomWay.Next2 = hiddenPassage
	paintingRoomWay.Next3 = returnToStart

	dungeonWay.Next1 = dragonFight
	dungeonWay.Next2 = goblinFight
	dungeonWay.Next3 = ghostFight

	currentSituation := startSituation

	for true {
		currentSituation.PrintSituationWithOptions()
		var userInputChoice string
		fmt.Scan(&userInputChoice)

		switch userInputChoice {
		case "1":
			if currentSituation.Next1 != nil {
				currentSituation = *currentSituation.Next1
			}
		case "2":
			if currentSituation.Next2 != nil {
				currentSituation = *currentSituation.Next2
			}
		case "3":
			if currentSituation.Next3 != nil {
				currentSituation = *currentSituation.Next3
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
	StartTheGame()
	Game()
}
