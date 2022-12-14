import { GeneratedNumbers } from './interfaces'
import { Output } from './output'

class HttpGeneration {
    // Адрес, на который необходимо делать пост запрос
    private url: string

    constructor(url: string) {
        this.url = url
    }

    // Генерирует и записывает данные, принимает форму и извлекает из нее параметры
    public async printStaticNumbers(
        unsortedData: HTMLElement,
        sortedData: HTMLElement,
        time: HTMLElement,
        formElem: HTMLFormElement,
        accordion: HTMLElement,
        output: Output
    ): Promise<void> {
        const js: GeneratedNumbers = await this.getStaticJSON(formElem)
        this.writeData(unsortedData, sortedData, time, js, output)
        output.getAndPrintHistory(accordion, '/history')
    }

    // Отправляет post запрос и возвращает json
    private async getStaticJSON(formElem: HTMLFormElement): Promise<GeneratedNumbers> {
        const response: Response = await fetch(this.url, {
            method: 'POST',
            body: new FormData(formElem),
        })
        return response.json()
    }

    // Записывает данные в поля
    private writeData(
        unsortedData: HTMLElement,
        sortedData: HTMLElement,
        time: HTMLElement,
        js: GeneratedNumbers,
        output: Output
    ): void {
        unsortedData.innerHTML = output.outputNumbers(js.unsorted_numbers)
        sortedData.innerHTML = output.outputNumbers(js.sorted_numbers)
        time.innerHTML = js.time + ' ns'
    }
}

export { HttpGeneration }
