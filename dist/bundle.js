(()=>{"use strict";({464:function(){var e=this&&this.__awaiter||function(e,n,t,i){return new(t||(t=Promise))((function(s,o){function u(e){try{r(i.next(e))}catch(e){o(e)}}function a(e){try{r(i.throw(e))}catch(e){o(e)}}function r(e){var n;e.done?s(e.value):(n=e.value,n instanceof t?n:new t((function(e){e(n)}))).then(u,a)}r((i=i.apply(e,n||[])).next())}))};const n=document.forms[0],t=document.getElementById("data__unsorted"),i=document.getElementById("data__sorted"),s=document.getElementById("data__time"),o=new WebSocket("ws://localhost:3000/ws");function u(e){return e.join(" ")}function a(e,n,t,i){e<1&&n<1?(t.classList.add("is-invalid"),i.classList.add("is-invalid")):e<1||isNaN(e)?t.classList.add("is-invalid"):i.classList.add("is-invalid")}function r(r){return e(this,void 0,void 0,(function*(){r.preventDefault();const d=document.querySelector("#exampleCheck1"),c=document.querySelector("#bound"),l=document.querySelector("#flows");c&&l&&d&&(c.classList.remove("is-invalid"),l.classList.remove("is-invalid"),d.checked?function(e,n){Number(e.value)>0&&Number(n.value)>0?(t&&i&&s&&(t.innerHTML="",i.innerHTML="",s.innerHTML=""),o.send(JSON.stringify({bound:e.value,flows:n.value}))):a(Number(e.value),Number(n.value),e,n)}(c,l):function(o,r){e(this,void 0,void 0,(function*(){if(Number(o.value)>0&&Number(r.value)>0&&n){const e=yield fetch("/numbers",{method:"POST",body:new FormData(n)}),o=yield e.json();t&&i&&s&&(t.innerHTML=u(o.unsorted_numbers),i.innerHTML=u(o.sorted_numbers),s.innerHTML=o.time+" ns")}else a(Number(o.value),Number(r.value),o,r)}))}(c,l))}))}o.onmessage=e=>{if(e.data.includes("time")&&i&&s){const n=JSON.parse(e.data);i.innerHTML=u(n.sorted_numbers),s.innerHTML=n.time+"ns"}else t&&(t.innerHTML+=e.data+" ")},n&&(n.onsubmit=e=>e.preventDefault()),document.addEventListener("DOMContentLoaded",(()=>{n.onsubmit=r}))}})[464]()})();