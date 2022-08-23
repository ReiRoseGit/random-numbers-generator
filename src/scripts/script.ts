import { Validation } from './modules/validation'
import { HttpGeneration } from './modules/httpGeneration'
import { WebsocketGeneration } from './modules/websocketGeneration'
import { htmlEl } from './modules/types'
import { Xml } from './modules/xml'

// Вся форма для ввода параметров
const formElem: HTMLFormElement | null = document.getElementById('formElem') as HTMLFormElement

// xml для получения истории
const xml: Xml = new Xml()

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
            ws.printDynamicNumbers(bound.value, flows.value, unsortedData, sortedData, time, xml, accordion)
            // xml.getLastGenerations(accordion)
        }
        // Ветка для статичного вывода чисел
        else if (formElem && accordion) {
            httpGen.printStaticNumbers(unsortedData, sortedData, time, formElem, xml, accordion)
            // xml.getLastGenerations(accordion)
        }
    }
}

xml.getLastGenerations(accordion)

accordion.addEventListener('click', (e: MouseEvent): void => {
    console.log(e.target)
})

// Отключает стандартное поведение кнопки
if (formElem) {
    formElem.onsubmit = (e: SubmitEvent): void => e.preventDefault()
}

document.addEventListener('DOMContentLoaded', (): void => {
    if (formElem) formElem.onsubmit = generateNumbers // todo
})
