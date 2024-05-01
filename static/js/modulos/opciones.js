export default function opciones(btnradio1, btnradio2) {
    const d= document;
    d.addEventListener("click",(e)=>{
        if (e.target.matches(btnradio1)) {
            d.querySelector(btnradio1).setAttribute("value", "opcionSpa");
            d.querySelector(btnradio2).setAttribute("value", "");
            console.log("se agrego opcion español")
            console.log("se elimino opcion Kekchi")
            console.log("-----------------------")

        }
        if (e.target.matches(btnradio2)) {
            d.querySelector(btnradio1).setAttribute("value", "");
            d.querySelector(btnradio2).setAttribute("value", "opcionKek");
            console.log("se agrego opcion Kekchi")
            console.log("se se elimino opcion español")
            console.log("-----------------------")
            }     
    })
}

