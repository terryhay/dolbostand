package about_strings

import (
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	fmt.Printf("Демонстрация строк\n\n")

	fmt.Println("Создадим строку и посмотрим на адреса при преобразовании ее в массив байт и рун")
	str := "string"

	fmt.Printf("%p: []byte[0]\t%p []rune[0]\n", &[]byte(str)[0], &[]rune(str)[0])
	fmt.Printf("%p: []byte[0]\t%p []rune[0]\n", &[]byte(str)[0], &[]rune(str)[0])

	fmt.Printf("\nАдреса разные. Всякий раз при преобразовании типов создается новый объект\n")

	fmt.Printf("\nПри конкатинации строка ведет себя как срез: адрес внешней части не меняется, но буфер реалоцируется\n")

	fmt.Printf("%p: %s", &str, str)
	for i := 0; i < 10; i++ {
		str += "1"
		fmt.Printf("%p: %s\n", &str, str)
	}

	fmt.Printf("\nСтрока хранится как массив байт, а не символов\n")
	fmt.Printf("\nОтрежем от латинской строки \"string\" один байт и посмотрим, что будет\n")

	str = "string"
	fmt.Printf("%s: str[1:]\n", str[1:])

	fmt.Printf("\nА теперь от строки \"строка\"...\n")

	str = "строка"
	fmt.Printf("%s: str[1:]\n", str[1:])

	fmt.Printf("А теперь преобразуем строку в руны, вырежем первую, и выведем обрубок\n")

	str = "строка"
	fmt.Printf("%s: string([]rune(str)[1:])\n", string([]rune(str)[1:]))
	fmt.Println("Заработало!")
}
