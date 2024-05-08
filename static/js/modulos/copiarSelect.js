
const seleccion=document.querySelector(".consulta-lista");
console.log(seleccion);

seleccion.addEventListener("change",()=>{
 let valorOption=seleccion.value;
  let inputEntrada=document.querySelector("#entrada");
  let inputHidden=document.querySelector("#id_consulta")
  console.log(valorOption);
  
  let opcionSelecionada = seleccion.options[seleccion.selectedIndex];

  console.log("opcion: ",opcionSelecionada.text);
  console.log("Valor: ",opcionSelecionada.value);
  
  inputEntrada.value=(opcionSelecionada.text);
  inputHidden.value=(opcionSelecionada.value);

})