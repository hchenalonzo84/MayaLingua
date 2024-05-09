
const d=document;
d.getElementById("entrada").addEventListener("keyup", getPalabras);
        function getPalabras() {
            let inputCP = d.getElementById("entrada").value;
            let lista = d.getElementById("lista"); 
            // let opLIsta = d.getElementById("opLIsta");           
            if (inputCP.length > 0) {
                let url = "/obtenerDatosJson";
                let formData = new FormData();
                formData.append("entrada", inputCP);
                 console.log(formData)
                fetch(url, {
                    method: "POST",
                    body: formData,
                    mode: "cors" //Default cors, no-cors, same-origin
                })
                .then(response => response.json())
                .then(data => {
                    lista.style.display = 'block';
                    // opLIsta.style.display='block';
                    lista.innerHTML = '';
                    data.forEach(item => {
                        let option = d.createElement('option');
                        option.value = item.id; // Asigna el valor del ID como valor de la opción
                     
                        option.textContent = `${item.spa}`;
                        lista.appendChild(option);
                        // console.log(option);
                    });
                })
                .catch(err => console.error('Error al obtener datos:', err));
            } else {
                lista.style.display = 'none';
            }
        }
       // ID: ${item.id},
        // , KEK: ${item.kek}