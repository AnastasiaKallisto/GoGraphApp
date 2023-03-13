function UploadGraphFromFile() {
    prompt("Button clicked!");
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

function switchPage() {
    var currentUrl = window.location.href;
    console.log(currentUrl);
    var newUrl = '';

    if (currentUrl.indexOf('/interval') !== -1) {
        newUrl = currentUrl.replace('/interval', '/exact');
    } else {
        newUrl = currentUrl.replace('/exact', '/interval');
    }
    console.log(newUrl);
    window.location.href = newUrl;
    console.log(window.location.href);
}