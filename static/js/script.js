function PrimAlgorithm(id) {

    document.getElementById(id).setAttribute('href',
        document.getElementById(id).getAttribute('href')
            .concat("?prim=true"));
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