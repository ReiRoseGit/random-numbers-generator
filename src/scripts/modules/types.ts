// Псевдоним, описывающий типы полей для вывода данных
type htmlEl = HTMLElement | null

// Псевдоним, описывающий ранее сгенерированую последовательность
type OldGeneration = {
    unsorted_numbers: number[]
    sorted_numbers: number[]
    time: string
    created_at: string
}

export { htmlEl, OldGeneration }
