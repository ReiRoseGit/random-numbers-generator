var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
var _this = this;
var formElem = document.getElementById("formElem");
var resultDataElements = document.getElementsByClassName('result__data');
var ws = new WebSocket("ws://localhost:3000/ws");
// Форматирует вывод данных
function outputNumbers(numbers) {
    return numbers.join(" ");
}
// Слушатель на сообщение от сервера, ожидает динамические данные и JSON
ws.onmessage = function (e) {
    if (e.data.includes("time")) {
        var js = JSON.parse(e.data);
        resultDataElements[1].innerHTML = outputNumbers(js['sorted_numbers']);
        resultDataElements[2].innerHTML = js['time'] + "ns";
    }
    else {
        resultDataElements[0].innerHTML += e.data + " ";
    }
};
// Обрабатывает отправку формы на генерацию случайных чисел.
// Если параметры конкретные, то добавляет содержимое в контейнеры
// иначе вызывает функцию проверки и изменения некорректных значений
formElem.onsubmit = function (e) { return __awaiter(_this, void 0, void 0, function () {
    var checkBox, bound, flows, response, result;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                e.preventDefault();
                checkBox = document.getElementById("exampleCheck1");
                bound = document.getElementById("bound");
                flows = document.getElementById("flows");
                bound.classList.remove("is-invalid");
                flows.classList.remove("is-invalid");
                if (!checkBox.checked) return [3 /*break*/, 1];
                if (Number(bound.value) > 0 && Number(flows.value) > 0) {
                    // Очистка значений для данных 
                    Array.from(resultDataElements).forEach(function (element) { element.innerHTML = ""; });
                    // Отправка параметров для получения последовательности случайных чисел
                    ws.send(JSON.stringify({ bound: bound.value, flows: flows.value }));
                }
                else {
                    cleanParams(Number(bound.value), Number(flows.value), bound, flows);
                }
                return [3 /*break*/, 5];
            case 1:
                if (!(Number(bound.value) > 0 && Number(flows.value) > 0)) return [3 /*break*/, 4];
                return [4 /*yield*/, fetch('/numbers', {
                        method: 'POST',
                        body: new FormData(formElem)
                    })];
            case 2:
                response = _a.sent();
                return [4 /*yield*/, response.json()];
            case 3:
                result = _a.sent();
                resultDataElements[0].innerHTML = outputNumbers(result['unsorted_numbers']);
                resultDataElements[1].innerHTML = outputNumbers(result['sorted_numbers']);
                resultDataElements[2].innerHTML = result['time'] + " ns";
                return [3 /*break*/, 5];
            case 4:
                cleanParams(Number(bound.value), Number(flows.value), bound, flows);
                _a.label = 5;
            case 5: return [2 /*return*/];
        }
    });
}); };
// Выполняет проверку недопустимых параметров
function cleanParams(boundValue, flowsValue, bound, flows) {
    if (boundValue < 1 && flowsValue < 1) {
        bound.classList.add("is-invalid");
        flows.classList.add("is-invalid");
    }
    else if (boundValue < 1 || isNaN(boundValue)) {
        bound.classList.add("is-invalid");
    }
    else {
        flows.classList.add("is-invalid");
    }
}
