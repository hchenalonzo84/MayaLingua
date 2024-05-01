import hamburgerMenu from "../modulos/menu-hamburguesa.js";
import opciones from "../modulos/opciones.js";
const d=document;

d.addEventListener("DOMContentLoaded",(e)=>{
    hamburgerMenu(".panel-btnPerso", ".panelperso",".menuPerso a");  
    
    opciones("#btnradio1","#btnradio2");

});

d.addEventListener("keydown",(e)=>{
    
console.info("JS del main")
    
})