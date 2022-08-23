class Validation {
    // Выполняет проверку недопустимых параметров и назначает классы недопустимым значениям инпутов
    public static cleanParams(bound: HTMLInputElement, flows: HTMLInputElement): boolean {
        this.cleanOutputFields(bound, flows)
        try {
            this.isCorrectValues(bound, flows)
            this.cleanOutputFields(bound, flows)
            return true
        } catch (e) {
            return false
        }
    }

    // Проверяет значения на допустимость, выбрасывает ошибку, если параметры не соответствуют логике
    private static isCorrectValues(bound: HTMLInputElement, flows: HTMLInputElement): void | Error {
        const boundValue: number = Number(bound.value)
        const flowsValue: number = Number(flows.value)
        if (boundValue < 1 && flowsValue < 1) {
            bound.classList.add('is-invalid')
            flows.classList.add('is-invalid')
            throw new Error('Недопустимые параметры')
        }
        if (boundValue < 1) {
            bound.classList.add('is-invalid')
            throw new Error('Недопустимые параметры')
        }
        if (flowsValue < 1 || flowsValue > 500) {
            flows.classList.add('is-invalid')
            throw new Error('Недопустимые параметры')
        }
    }

    // Очистка полей
    private static cleanOutputFields(...fields: HTMLInputElement[]): void {
        fields.forEach((element: HTMLInputElement): void => element.classList.remove('is-invalid'))
    }
}

export { Validation }
