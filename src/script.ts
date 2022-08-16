let formElem = document.getElementById("formElem") as HTMLFormElement
let resultDataElements = document.getElementsByClassName('result__data') as HTMLCollectionOf<Element>

let ws : WebSocket = new WebSocket("ws://localhost:3000/ws");

// Форматирует вывод данных
function outputNumbers(numbers : number[]):string{
    return numbers.join(" ")
}

// Слушатель на сообщение от сервера, ожидает динамические данные и JSON
ws.onmessage = function(e):void{
    if (e.data.includes("time")){
        let js : any =  JSON.parse(e.data)
        resultDataElements[1].innerHTML = outputNumbers(js['sorted_numbers'])
        resultDataElements[2].innerHTML = js['time'] + "ns"
    }
    else{
        resultDataElements[0].innerHTML += e.data + " "
    }
}

// Обрабатывает отправку формы на генерацию случайных чисел.
// Если параметры конкретные, то добавляет содержимое в контейнеры
// иначе вызывает функцию проверки и изменения некорректных значений
formElem.onsubmit = async (e) => {
    e.preventDefault();
    let checkBox = document.getElementById("exampleCheck1") as HTMLInputElement
    let bound = document.getElementById("bound") as HTMLInputElement
    let flows = document.getElementById("flows") as HTMLInputElement
    bound.classList.remove("is-invalid")
    flows.classList.remove("is-invalid")
    // Ветка для динамического вывода чисел
    if (checkBox.checked){
        if (Number(bound.value) > 0 && Number(flows.value) > 0) {
            // Очистка значений для данных 
            Array.from(resultDataElements).forEach(element => {element.innerHTML = ""})
            // Отправка параметров для получения последовательности случайных чисел
            ws.send(JSON.stringify({bound:bound.value, flows:flows.value}))
        }
        else{
            cleanParams(Number(bound.value), Number(flows.value), bound, flows)
        }
    }
    // Ветка для вывода чисел по завершении генерации
    else{
        if (Number(bound.value) > 0 && Number(flows.value) > 0) {
            let response = await fetch('/numbers', {
                method: 'POST',
                body: new FormData(formElem)
            });
            let result : Promise<any> = await response.json();
            resultDataElements[0].innerHTML = outputNumbers(result['unsorted_numbers'])
            resultDataElements[1].innerHTML = outputNumbers(result['sorted_numbers'])
            resultDataElements[2].innerHTML = result['time'] + " ns"
        }
        else{
            cleanParams(Number(bound.value), Number(flows.value), bound, flows)
        }
    }
};

// Выполняет проверку недопустимых параметров
function cleanParams(boundValue : number, flowsValue : number, bound : HTMLInputElement, flows : HTMLInputElement) : void{
    if (boundValue < 1 && flowsValue < 1){
        bound.classList.add("is-invalid")
        flows.classList.add("is-invalid") 
    }
    else if (boundValue < 1 || isNaN(boundValue)){
        bound.classList.add("is-invalid")
    }
    else{
        flows.classList.add("is-invalid")
    }
}
