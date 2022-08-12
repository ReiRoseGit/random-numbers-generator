/*
 * ATTENTION: The "eval" devtool has been used (maybe by default in mode: "development").
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
/******/ (() => { // webpackBootstrap
/******/ 	var __webpack_modules__ = ({

/***/ "./src/script.js":
/*!***********************!*\
  !*** ./src/script.js ***!
  \***********************/
/***/ (() => {

eval("let formElem = document.getElementById(\"formElem\")\r\n\r\nformElem.onsubmit = async (e) => {\r\n    e.preventDefault();\r\n    let bound = document.getElementById(\"bound\")\r\n    let flows = document.getElementById(\"flows\")\r\n    bound.classList.remove(\"is-invalid\")\r\n    flows.classList.remove(\"is-invalid\")\r\n    let errorCode = cleanParams(bound.value, flows.value)\r\n    if (!errorCode){\r\n        let response = await fetch('/numbers', {\r\n            method: 'POST',\r\n            body: new FormData(formElem)\r\n        });\r\n        let result = await response.json();\r\n        let resultElements = document.getElementsByClassName('result__element')\r\n        resultElements[0].getElementsByClassName('result__data')[0].innerHTML = result['unsorted_numbers']\r\n        resultElements[1].getElementsByClassName('result__data')[0].innerHTML = result['sorted_numbers']\r\n        resultElements[2].getElementsByClassName('result__data')[0].innerHTML = result['time'] + \" ns\"\r\n    }\r\n    else{\r\n        if (errorCode == 300){\r\n            bound.classList.add(\"is-invalid\")\r\n            flows.classList.add(\"is-invalid\")\r\n        }\r\n        else if (errorCode == 200){\r\n            flows.classList.add(\"is-invalid\")\r\n        }\r\n        else{\r\n            bound.classList.add(\"is-invalid\")\r\n        }\r\n    }\r\n};\r\n\r\nfunction cleanParams(bound, flows){\r\n    if (bound > 0 && flows > 0) return 0\r\n    if (bound < 1 && flows < 1) return 300\r\n    if (bound < 1 || isNaN(parseInt(bound))) return 100\r\n    if (flows < 1 || isNaN(parseInt(flows))) return 200\r\n    return 300\r\n}\n\n//# sourceURL=webpack://random-numbers/./src/script.js?");

/***/ })

/******/ 	});
/************************************************************************/
/******/ 	
/******/ 	// startup
/******/ 	// Load entry module and return exports
/******/ 	// This entry module can't be inlined because the eval devtool is used.
/******/ 	var __webpack_exports__ = {};
/******/ 	__webpack_modules__["./src/script.js"]();
/******/ 	
/******/ })()
;