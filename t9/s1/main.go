package t9s1

import (
	"log"
	"sync"
)

func main() {
	numbers := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// канал для чисел
	c1 := make(chan int)
	// канал для чисел умноженных на 2
	c2 := make(chan int)
	wg := sync.WaitGroup{}

	// первая горутина записывает числа в канал c1
	wg.Add(1)
	go func() {
		// отмечаем, что горутина завершилась
		defer wg.Done()
		// после цикла (когда все числа запишутся) закрываем канал
		// что означает, что туда больше ничего не будет записываться
		defer close(c1)

		for _, n := range numbers {
			// т.к. канал не буферизованный, то тут произойдет блокировка
			// пока из канала не считается значение
			c1 <- n
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(c2)

		// считываем числа из канала c1
		// может быть блокировка, если в c1 нет чисел
		// если оно появляется выполнение продолжается
		// считывание разблокирует предыдущую горутину
		for n := range c1 {
			c2 <- n * 2
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		// выводим  x2 числа по мере поступления в c2
		for n := range c2 {
			log.Println(n)
		}
	}()

	// ожидаем завершения всех горутин
	wg.Wait()
}
