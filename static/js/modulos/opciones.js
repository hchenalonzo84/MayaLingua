export default function opciones(btnradio1, btnradio2, entrada, salida) {
    const d = document;
   let entradaN= d.getElementById(entrada);
   let salidaN=d.getElementById(salida);
   let selectLista=d.getElementById("lista");
   let selectLista2=d.getElementById("lista2");
    d.addEventListener("click", (e) => {
        if (e.target.matches(btnradio1)) {
            d.querySelector(btnradio1).setAttribute("value", "opcionSpa");
            d.querySelector(btnradio2).setAttribute("value", "");
            entradaN.readOnly = false;
            salidaN.readOnly = true;
            selectLista2.style.display='none';
            limpiarCampos(entradaN,salidaN);
            console.log("Se agregó la opción español");
            console.log("Se eliminó la opción Kekchi");
            console.log("-----------------------");
        }
        if (e.target.matches(btnradio2)) {
            d.querySelector(btnradio1).setAttribute("value", "");
            d.querySelector(btnradio2).setAttribute("value", "opcionKek");
            salidaN.readOnly = false;
            entradaN.readOnly = true;
            selectLista.style.display='none';
            limpiarCampos(entradaN,salidaN);
            console.log("Se agregó la opción Kekchi");
            console.log("Se eliminó la opción español");
            console.log("-----------------------");
        }     
    });
}

function limpiarCampos(entradaN,salidaN) {
    entradaN.value='';
    salidaN.value='';
}
    

