// Интерфейс для JSON данных, приходящих с сервера
interface GeneratedNumbers {
    unsorted_numbers: number[]
    sorted_numbers: number[]
    time: string
}

interface LastGeneratedNumbers extends GeneratedNumbers {
    created_at: string
}

export { GeneratedNumbers, LastGeneratedNumbers }
