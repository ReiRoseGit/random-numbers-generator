import { GeneratedNumbers } from './interfaces'
import { Output } from './output'
import { Xml } from './xml'

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
        xml: Xml,
        accordion: HTMLElement
    ): Promise<void> {
        const js: GeneratedNumbers = await this.getStaticJSON(formElem)
        this.writeData(unsortedData, sortedData, time, js)
        xml.getLastGenerations(accordion)
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
        js: GeneratedNumbers
    ): void {
        unsortedData.innerHTML = Output.outputNumbers(js.unsorted_numbers)
        sortedData.innerHTML = Output.outputNumbers(js.sorted_numbers)
        time.innerHTML = js.time + ' ns'
    }
}

export { HttpGeneration }
