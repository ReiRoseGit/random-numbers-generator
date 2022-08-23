import { OldGeneration } from './types'

class Output {
    constructor() {}
    // Форматирует вывод данных
    public outputNumbers(numbers: number[]): string {
        return numbers.join(' ')
    }

    // Выводит последние генерации в аккордеон, если последних генераций нет, то вывод заголовок с просьбой выполнить генерацию
    public printLastGenerations(accordion: HTMLElement, data: OldGeneration[]): void {
        accordion.innerHTML = ''
        if (data.length > 0) {
            for (let i: number = 0; i < data.length; i++) {
                accordion.innerHTML += `<div class="accordion__item">
                <div class="accordion__header" style="display: flex;">
                    Генерация #${i + 1} : дата ${data[i].created_at}
                </div>
                <div class="accordion__body">
                    <div class="previous-generations__unsorted_numbers"><strong>Неотсортированные данные:</strong> ${this.outputNumbers(
                        data[i].unsorted_numbers
                    )}</div>
                    <div class="previous-generations__unsorted_numbers"><strong>Отсортированные данные:</strong> ${this.outputNumbers(
                        data[i].sorted_numbers
                    )}</div>
                    <div class="previous-generations__unsorted_numbers"><strong>Время генерации:</strong> ${
                        data[i].time
                    } ns</div>
                </div>
              </div>`
            }
            const list: NodeListOf<Element> = accordion.querySelectorAll('.accordion__item')
            list.forEach((item: Element): void => {
                item.addEventListener('click', (): void => {
                    item.classList.contains('accordion__item_show')
                        ? item.classList.remove('accordion__item_show')
                        : item.classList.add('accordion__item_show')
                })
            })
        } else {
            accordion.innerHTML =
                '<h6 style="text-align: center;">Для отображения истории необходимо выполнить хотя бы одну генерацию!</h6>'
        }
    }

    // Отправляет Get запрос на адрес /history и вызывает функцию вывода информации о последних генерациях
    public async getAndPrintHistory(accordion: HTMLElement, url: string): Promise<void> {
        const response: Response = await fetch(url, { method: 'GET' })
        this.printLastGenerations(accordion, await response.json())
    }

    // Удаляет все данные о последних генерациях из базы данных
    public async clearDataBase(url: string): Promise<void> {
        await fetch(url, { method: 'DELETE' })
    }
}

export { Output }
