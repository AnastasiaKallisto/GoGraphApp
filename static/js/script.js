function UploadGraphFromFile(id) {
    openForm(id);
}

function PrimAlgorithm() {
    alert("Button clicked!");
}

function CruscalAlgorithm() {
    alert("Button clicked!");
}

function openForm(id) {
    document.getElementById(id).showModal();
}

function closeForm(id) {
    document.getElementById(id).close();
}

function uploadFile() {
    let fileInput = document.getElementById("txtFile");
    let file = fileInput.files[0];
    let formData = new FormData();
    formData.append("file", file);

    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/upload", true);
    xhr.onload = function () {
        if (xhr.readyState === xhr.DONE) {
            if (xhr.status === 200) {
                console.log(xhr.responseText);
            }
        }
    };
    xhr.send(formData);
}