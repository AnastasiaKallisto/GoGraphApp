const GraphType = { Exact: 'Граф с точными весами', Interval: 'Граф с интервальными весами'};

class GraphicGraph {
    constructor(id, shortDescription, longDescription, imageSrc, type) {
        this.id = id;
        this.shortDescription = shortDescription;
        this.longDescription = longDescription;
        this.imageSrc = imageSrc;
        this.type = type;
    }

}



function openIntervalGraphInfo(id, graphNumber){
    uploadIntervalGraphInfoOnForm(document.getElementById(id), graphNumber);
    document.getElementById(id).showModal();
}

function uploadIntervalGraphInfoOnForm(id, graphNumber){

    document.getElementById(id).setAttribute("graphNumber", graphNumber);
}