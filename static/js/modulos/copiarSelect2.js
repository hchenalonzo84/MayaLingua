
const seleccion2=document.querySelector(".consulta-lista2");
console.log(seleccion2);

seleccion2.addEventListener("change",()=>{
 let valorOption2=seleccion2.value;
  let inputSalida=document.querySelector("#salida");
  let inputHidden2=document.querySelector("#id_consulta2")
  console.log(valorOption2);
  
  let opcionSelecionada = seleccion2.options[seleccion2.selectedIndex];

  console.log("opcion2: ",opcionSelecionada.text);
  console.log("Valor2: ",opcionSelecionada.value);
  
  inputSalida.value=(opcionSelecionada.text);
  inputHidden2.value=(opcionSelecionada.value);

})