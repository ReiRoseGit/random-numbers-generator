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

eval("let formElem = document.getElementById(\"formElem\")\r\n\r\nformElem.onsubmit = async (e) => {\r\n    e.preventDefault();\r\n    let response = await fetch('/numbers', {\r\n      method: 'POST',\r\n      body: new FormData(formElem)\r\n    });\r\n\r\n    let result = await response.json();\r\n    let resultElements = document.getElementsByClassName('result__element')\r\n    if ('error_code' in result){\r\n      resultElements[0].getElementsByClassName('result__data')[0].innerHTML = result['error_message']\r\n      resultElements[1].getElementsByClassName('result__data')[0].innerHTML = result['error_message']\r\n      resultElements[2].getElementsByClassName('result__data')[0].innerHTML = result['error_message']\r\n    }\r\n    else{\r\n      resultElements[0].getElementsByClassName('result__data')[0].innerHTML = result['unsorted_numbers']\r\n      resultElements[1].getElementsByClassName('result__data')[0].innerHTML = result['sorted_numbers']\r\n      resultElements[2].getElementsByClassName('result__data')[0].innerHTML = result['time'] + \" ns\"\r\n    }\r\n\r\n};\n\n//# sourceURL=webpack://random-numbers/./src/script.js?");

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