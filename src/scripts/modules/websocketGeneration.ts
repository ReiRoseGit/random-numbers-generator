import { GeneratedNumbers } from './interfaces'
import { Output } from './output'

// Класс, реализующий генерацию чисел по вебсокету
class WebsocketGeneration {
    websocket: WebSocket

    constructor(url: string) {
        this.websocket = new WebSocket(url)
    }

    // Получает параметры для генерации чисел, преобразует их в JSON и отправляет на сервер
    private getDynamicNumbers(boundValue: string, flowsValue: string): void {
        this.websocket.send(
            JSON.stringify({
                bound: boundValue,
                flows: flowsValue,
            })
        )
    }

    // Принимает поля для записи готовых данных, назначает обработчик, записывающий данные в эти поля
    private setFieldsAndData(unsortedData: HTMLElement, sortedData: HTMLElement, time: HTMLElement): void {
        unsortedData.innerHTML = sortedData.innerHTML = time.innerHTML = ''
        this.websocket.onmessage = (e: MessageEvent): void => {
            if (e.data.includes('time')) {
                const js: GeneratedNumbers = JSON.parse(e.data)
                sortedData.innerHTML = Output.outputNumbers(js.sorted_numbers)
                time.innerHTML = js.time + 'ns'
            } else {
                unsortedData.innerHTML += e.data + ' '
            }
        }
    }

    // Запускает генерацию и запись чисел по вебсокету
    public printDynamicNumbers(
        boundValue: string,
        flowsValue: string,
        unsortedData: HTMLElement,
        sortedData: HTMLElement,
        time: HTMLElement
    ): void {
        this.getDynamicNumbers(boundValue, flowsValue)
        this.setFieldsAndData(unsortedData, sortedData, time)
    }
}

export { WebsocketGeneration }