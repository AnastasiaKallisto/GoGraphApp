let graphIsDrawn = false;

function UploadGraphFromFile() {
    prompt("Button clicked!");
}

function PrimAlgorithm() {
    alert("Button clicked!");
}

function CruscalAlgorithm() {
    alert("Button clicked!");
}

function Clear() {
    if (graphIsDrawn) {
        if(confirm("Вы действительно хотите очистить рабочее пространство?")){
            // вызов функции из go, которая сбрасывает все массивы с графами
            // все переменные, касающиеся того, какой сейчас алгоритм выполняется
            // не трогаем только переменную, которая касается весов графа
            //
        }
    }
}

function ExactWeights() {
    confirm("Вы действительно хотите переключиться на работу с точными весами?");
}

function IntervalWeights() {
    confirm("Вы действительно хотите переключиться на работу с интервальными весами?");
}

function openForm(id) {
    document.getElementById(id).showModal();
}

function closeForm(id) {
    document.getElementById(id).close();
}