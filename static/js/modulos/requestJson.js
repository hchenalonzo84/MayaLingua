document.getElementById("campo").addEventListener("keyup", getPalabras)

function getPalabras() {

    let inputCP = document.getElementById("campo").value
    let lista = document.getElementById("lista")

    if (inputCP.length > 0) {

        let url = "db/getPalabras.php"
        let formData = new FormData()
        formData.append("campo", inputCP)

        fetch(url, {
            method: "POST",
            body: formData,
            mode: "cors" //Default cors, no-cors, same-origin
        }).then(response => response.json())
            .then(data => {
                lista.style.display = 'block'
                lista.innerHTML = data
            })
            .catch(err => console.log(err))
    } else {
        lista.style.display = 'none'
    }
}