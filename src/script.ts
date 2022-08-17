// Интерфейс для JSON данных, приходящих с сервера
interface GeneratedNumbers{
    unsorted_numbers : number[]
    sorted_numbers : number[]
    time : string
}

// Псевдоним, описывающий типы полей для вывода данных
type htmlEl = HTMLElement | null;

// Вся форма для ввода параметров
const formElem : HTMLFormElement = document.forms[0];

// Поля для вывода данных
const unsortedData : htmlEl = document.getElementById('data__unsorted');
const sortedData : htmlEl = document.getElementById('data__sorted');
const time : htmlEl = document.getElementById('data__time');

// Websocket
const ws : WebSocket = new WebSocket("ws://localhost:3000/ws");

// Форматирует вывод данных
function outputNumbers(numbers : number[]):string{
    return numbers.join(" ");
}

// Динамический вывод чисел с помощью вебсокета
function getDynamicNumbers(bound : HTMLInputElement, flows : HTMLInputElement):void{
    if (Number(bound.value) > 0 && Number(flows.value) > 0) {
        // Очистка значений для данных 
        if (unsortedData && sortedData && time){
            unsortedData.innerHTML = "";
            sortedData.innerHTML = "";
            time.innerHTML = "";
        }
        // Отправка параметров для получения последовательности случайных чисел
        ws.send(JSON.stringify({bound:bound.value, flows:flows.value}));
    }
    else{
        cleanParams(Number(bound.value), Number(flows.value), bound, flows);
    }
}

// Статичный вывод чисел с помощью http
async function getStaticNumbers(bound : HTMLInputElement, flows : HTMLInputElement):Promise<void>{
        if (Number(bound.value) > 0 && Number(flows.value) > 0 && formElem) {
            const response : Response = await fetch('/numbers', {
                method: 'POST',
                body: new FormData(formElem)
            });
            const result : GeneratedNumbers = await response.json();
            if (unsortedData && sortedData && time){
                unsortedData.innerHTML = outputNumbers(result.unsorted_numbers);
                sortedData.innerHTML = outputNumbers(result.sorted_numbers);
                time.innerHTML = result.time + " ns"
            }
        }
        else{
            cleanParams(Number(bound.value), Number(flows.value), bound, flows);
        }
}

// Выполняет проверку недопустимых параметров и назначает классы недопустимым значениям инпутов
function cleanParams(boundValue : number, flowsValue : number, bound : HTMLInputElement, flows : HTMLInputElement) : void{
    if (boundValue < 1 && flowsValue < 1){
        bound.classList.add("is-invalid");
            flows.classList.add("is-invalid") ;
    }
    else if (boundValue < 1 || isNaN(boundValue)){
        bound.classList.add("is-invalid");
    }
    else{
        flows.classList.add("is-invalid");
    }
}

// Слушатель на сообщение от сервера, ожидает динамические данные и JSON
ws.onmessage = (e : MessageEvent) : void => {
    if (e.data.includes("time") && sortedData && time){
        const js : GeneratedNumbers =  JSON.parse(e.data);
        sortedData.innerHTML = outputNumbers(js.sorted_numbers);
        time.innerHTML = js.time + "ns";
    }
    else{
        if (unsortedData) unsortedData.innerHTML += e.data + " ";
    }
}

// Отключает стандартное поведение кнопки
if (formElem){
    formElem.onsubmit = (e : SubmitEvent) : void => e.preventDefault();
}

// Обрабатывает отправку формы на генерацию случайных чисел.
// Если параметры конкретные, то добавляет содержимое в контейнеры
// иначе вызывает функцию проверки и изменения некорректных значений
async function generateNumbers(e : SubmitEvent): Promise<void>{
    e.preventDefault();
    const checkBox : HTMLInputElement | null = document.querySelector('#exampleCheck1');
    const bound : HTMLInputElement | null = document.querySelector('#bound');
    const flows : HTMLInputElement | null = document.querySelector('#flows');
    if (bound && flows && checkBox){
        bound.classList.remove("is-invalid");
        flows.classList.remove("is-invalid");
        // Ветка для динамического вывода чисел
        if (checkBox.checked){
            getDynamicNumbers(bound, flows);
        }
        // Ветка для статичного вывода чисел
        else{
            getStaticNumbers(bound, flows);
        }
    }
}

document.addEventListener("DOMContentLoaded", () : void => {
    formElem.onsubmit = generateNumbers;
})