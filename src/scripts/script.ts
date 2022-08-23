import { Validation } from './modules/validation'
import { HttpGeneration } from './modules/httpGeneration'
import { WebsocketGeneration } from './modules/websocketGeneration'
import { htmlEl } from './modules/types'
import { Output } from './modules/output'
import '../public/css/style.css'

function content(): void {
    // Вся форма для ввода параметров
    const formElem: HTMLFormElement | null = document.getElementById('formElem') as HTMLFormElement

    // Поля для вывода данных
    const unsortedData: htmlEl = document.getElementById('data__unsorted')
    const sortedData: htmlEl = document.getElementById('data__sorted')
    const time: htmlEl = document.getElementById('data__time')

    // Поле для вывода истории генерации
    const accordion: htmlEl = document.getElementById('previousGenerations__accordion') as HTMLElement

    // Создание websocket
    const ws: WebsocketGeneration = new WebsocketGeneration('ws://localhost:3000/ws')

    // Создание объекта класса для статического вывода данных
    const httpGen: HttpGeneration = new HttpGeneration('/numbers')

    // Кнопка для удаления списка генераций
    const deleteButton: htmlEl = document.getElementById('db_delete') as HTMLElement

    // Форматированный вывод данных
    const output: Output = new Output()

    // Модальное окно (по умолчанию скрыто)
    const modal: htmlEl = document.getElementById('modal') as HTMLElement

    // Кнопка закрытия модального окна
    const modalBtn: htmlEl = document.getElementById('modal_close__btn') as HTMLElement

    // Обрабатывает отправку формы на генерацию случайных чисел.
    // Если параметры конкретные, то добавляет содержимое в контейнеры
    // иначе вызывает функцию проверки и изменения некорректных значений
    async function generateNumbers(e: SubmitEvent): Promise<void> {
        e.preventDefault()
        const bound: HTMLInputElement | null = document.querySelector('#bound')
        const flows: HTMLInputElement | null = document.querySelector('#flows')
        if (bound && flows && unsortedData && sortedData && time && Validation.cleanParams(bound, flows)) {
            const checkBox: HTMLInputElement | null = document.querySelector('#exampleCheck1')
            // Ветка для динамического вывода чисел
            if (checkBox && checkBox.checked && accordion) {
                ws.printDynamicNumbers(bound.value, flows.value, unsortedData, sortedData, time, accordion, output)
            }
            // Ветка для статичного вывода чисел
            else if (formElem && accordion) {
                httpGen.printStaticNumbers(unsortedData, sortedData, time, formElem, accordion, output)
            }
        }
    }

    // Обработчик на отправку формы
    formElem.addEventListener('submit', generateNumbers)

    // Выводит данные о прошлых генерациях в аккордеон
    output.getAndPrintHistory(accordion, '/history')

    // Обработчик событий на button удаления истории генераций
    deleteButton.addEventListener('click', (): void => {
        output.clearDataBase('/history')
        output.getAndPrintHistory(accordion, '/history')
    })

    // Обработчик на закрытие соединения
    ws.websocket.addEventListener('close', (): void => {
        document.body.style.overflow = 'hidden'
        modal.style.display = 'block'
    })

    // Обработчик закрытия кнопки
    modalBtn.addEventListener('click', (): void => {
        modal.style.display = 'none'
        document.body.style.overflow = 'auto'
    })
}

document.addEventListener('DOMContentLoaded', content)
