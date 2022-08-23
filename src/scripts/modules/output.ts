import { LastGeneratedNumbers } from './interfaces'
import { htmlEl, OldGeneration } from './types'

class Output {
    // Форматирует вывод данных
    public static outputNumbers(numbers: number[]): string {
        return numbers.join(' ')
    }
    public static printLastGenerations(accordion: HTMLElement, data: OldGeneration[]): void {
        accordion.innerHTML = ''
        for (let i: number = 0; i < data.length; i++) {
            accordion.innerHTML += `<div class="accordion__item">
            <div class="accordion__header" style="display: flex;">
                Генерация #${i + 1} : дата ${data[i].created_at}
            </div>
            <div class="accordion__body">
                <div class="previous-generations__unsorted_numbers"><strong>Неотсортированные данные:</strong> ${Output.outputNumbers(
                    data[i].unsorted_numbers
                )}</div>
                <div class="previous-generations__unsorted_numbers"><strong>Отсортированные данные:</strong> ${Output.outputNumbers(
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
            item.addEventListener('click', (e: Event): void => {
                item.classList.contains('accordion__item_show')
                    ? item.classList.remove('accordion__item_show')
                    : item.classList.add('accordion__item_show')
            })
        })
    }
    public static async getAndPrintHistory(accordion: HTMLElement, url: string): Promise<void> {
        const response: Response = await fetch(url, { method: 'GET' })
        Output.printLastGenerations(accordion, await response.json())
    }
}

export { Output }
