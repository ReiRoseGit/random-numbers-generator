(()=>{let e=document.getElementById("formElem");e.onsubmit=async s=>{s.preventDefault();let t=document.getElementById("bound"),a=document.getElementById("flows");t.classList.remove("is-invalid"),a.classList.remove("is-invalid");let n=function(e,s){return e>0&&s>0?0:e<1&&s<1?300:e<1||isNaN(parseInt(e))?100:s<1||isNaN(parseInt(s))?200:300}(t.value,a.value);if(n)300==n?(t.classList.add("is-invalid"),a.classList.add("is-invalid")):200==n?a.classList.add("is-invalid"):t.classList.add("is-invalid");else{let s=await fetch("/numbers",{method:"POST",body:new FormData(e)}),t=await s.json(),a=document.getElementsByClassName("result__element");a[0].getElementsByClassName("result__data")[0].innerHTML=t.unsorted_numbers,a[1].getElementsByClassName("result__data")[0].innerHTML=t.sorted_numbers,a[2].getElementsByClassName("result__data")[0].innerHTML=t.time+" ns"}}})();