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
            accordion.innerHTML += `<div class="accordion__item accordion__item_show">
            <div class="accordion__header">
                Генерация #${i + 1} : дата ${data[i].created_at}
            </div>
            <div class="accordion__body">
                <div class="previous-generations__unsorted_numbers">Неотсортированные данные: ${Output.outputNumbers(
                    data[i].unsorted_numbers
                )}</div>
                <div class="previous-generations__unsorted_numbers">Отсортированные данные: ${Output.outputNumbers(
                    data[i].sorted_numbers
                )}</div>
                <div class="previous-generations__unsorted_numbers">Время генерации: ${data[i].time} ns</div>
            </div>
          </div>`
        }
    }
}

export { Output }
