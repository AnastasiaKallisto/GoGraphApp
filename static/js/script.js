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
        if(alert("Вы действительно хотите очистить рабочее пространство?")){
            // вызов функции из go, которая сбрасывает все массивы с графами
            // все переменные, касающиеся того, какой сейчас алгоритм выполняется
            // не трогаем только переменную, которая касается весов графа
            //
        }
    }
}

function ExactWeights() {
    prompt("Вы действительно хотите переключиться на работу с точными весами?");
}

function IntervalWeights() {
    prompt("Вы действительно хотите переключиться на работу с интервальными весами?");
}

function openGenerateGraphForm() {
    document.getElementById("generateGraphForm").showModal();
}

function closeGenerateGraphForm() {
    document.getElementById("generateGraphForm").close();
}