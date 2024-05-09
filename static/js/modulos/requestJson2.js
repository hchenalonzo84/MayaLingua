
const d2=document;
d2.getElementById("salida").addEventListener("keyup", getPalabras2);
        function getPalabras2() {
            let inputCP = d2.getElementById("salida").value;
            let lista = d2.getElementById("lista2");            
            if (inputCP.length > 0) {
                let url = "/obtenerDatosJson2";
                let formData = new FormData();
                formData.append("salida", inputCP);
                 console.log(formData)
                fetch(url, {
                    method: "POST",
                    body: formData,
                    mode: "cors" //Default cors, no-cors, same-origin
                })
                .then(response => response.json())
                .then(data => {
                    lista.style.display = 'block';
                    lista.innerHTML = '';
                    data.forEach(item => {
                        let option = d2.createElement('option');
                        option.value = item.id; // Asigna el valor del ID como valor de la opción
                     
                        option.textContent = `${item.kek}`;
                        lista.appendChild(option);
                        // option.style.display='block';
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