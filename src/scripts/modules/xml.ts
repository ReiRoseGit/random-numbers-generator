import { Output } from './output'

class Xml {
    private xml: XMLHttpRequest
    constructor() {
        this.xml = new XMLHttpRequest()
    }
    public getLastGenerations(accordion: HTMLElement): void {
        this.xml.open('GET', '/history')
        this.xml.send()
        this.xml.onload = function (): void {
            Output.printLastGenerations(accordion, JSON.parse(this.response))
        }
    }
}

export { Xml }
