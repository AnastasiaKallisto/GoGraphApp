class Vertex {
    constructor(x, y, number) {
        this.x = x;
        this.y = y;
        this.number = number;
    }
}

class ExactEdge {
    constructor(a, b, weight) {
        this.a = a;
        this.b = b;
        this.weight = weight;
    }
}

class ExactGraph {

    constructor(edges, vertices) {
        this.vertices = [];
        this.edges = [];

        for (let i = 0; i < vertices.length; ++i) {
            this.vertices[i] = vertices[i];
        }
        for (let i = 0; i < edges.length; ++i) {
            this.edges[i] = edges[i];
        }
    }

    addVertex(vertex) {
        this.vertices.push(vertex);
    }

    addEdge(edge) {
        this.edges.push(edge);
    }
}

function generateVertices(n, sizeFrameX, sizeFrameY, graph) {
    console.log(sizeFrameX);
    console.log(sizeFrameY);
    let centerX = sizeFrameX / 2;
    let centerY = sizeFrameY / 2;
    let radius = Math.floor(centerY * 0.9);

    for (let i = 0; i < n; i++) {
        graph.vertices[i] = new Vertex(
            Math.floor(centerX + radius * Math.cos(i * Math.PI * 2 / n)),
            Math.floor(centerY + radius * Math.sin(i * Math.PI * 2 / n)),
            i
        );
    }
    let number = 0;
    for (let i = 0; i < graph.edges.length; i++) {
        number = graph.edges[i].a.number;
        graph.edges[i].a.x = graph.vertices[number].x;
        graph.edges[i].a.y = graph.vertices[number].y;
        number = graph.edges[i].b.number;
        graph.edges[i].b.x = graph.vertices[number].x;
        graph.edges[i].b.y = graph.vertices[number].y;
    }
}

function drawGraph(graph) {
  let canvas = document.getElementById('canvasForGraph');
  let ctx = canvas.getContext('2d');
    generateVertices(graph.vertices.length, canvas.clientWidth, canvas.clientHeight, graph)
    // граф
    // Ребра
    ctx.strokeStyle = 'black';
    ctx.lineWidth = 1;
    ctx.font = "12px serif";
    ctx.fillStyle = 'gray';
    graph.edges.forEach(function(edge) {
      ctx.beginPath();
      ctx.moveTo(edge.a.x, edge.a.y);
      ctx.lineTo(edge.b.x, edge.b.y);
      ctx.stroke();
        ctx.fillText( edge.weight, Math.round((edge.a.x+edge.b.x)*0.5)-10, Math.round((edge.a.y+edge.b.y)*0.5)-10);
    });
    // Вершины
    ctx.fillStyle = 'black';
    graph.vertices.forEach(function(vertex) {
      ctx.beginPath();
      ctx.arc(vertex.x, vertex.y, 2, 0, 2*Math.PI);
      ctx.fillText("V" + vertex.number, vertex.x+5, vertex.y+15);
      ctx.fill();
    });
}

function drawPrimTree(graph) {
    document.getElementById("textGraph").value = "kekeke";
    let canvas = document.getElementById('canvasForGraph');
    let ctx = canvas.getContext('2d');
    generateVertices(graph.vertices.length, canvas.clientWidth, canvas.clientHeight, graph)
    // граф
    // Ребра
    ctx.strokeStyle = 'blue';
    ctx.lineWidth = 1;
    ctx.font = "12px serif";
    ctx.fillStyle = 'blue';
    graph.edges.forEach(function(edge) {
        ctx.beginPath();
        ctx.moveTo(edge.a.x +3, edge.a.y+3);
        ctx.lineTo(edge.b.x+3, edge.b.y+3);
        ctx.stroke();
    });
    // Вершины
    ctx.fillStyle = 'blue';
    graph.vertices.forEach(function(vertex) {
        ctx.beginPath();
        ctx.arc(vertex.x+3, vertex.y+3, 2, 0, 2*Math.PI);
        ctx.fill();
    });
}

function drawCruscalTree(graph) {
    let canvas = document.getElementById('canvasForGraph');
    let ctx = canvas.getContext('2d');
    generateVertices(graph.vertices.length, canvas.clientWidth, canvas.clientHeight, graph)
    // граф
    // Ребра
    ctx.strokeStyle = 'red';
    ctx.lineWidth = 1;
    ctx.font = "12px serif";
    ctx.fillStyle = 'red';
    graph.edges.forEach(function(edge) {
        ctx.beginPath();
        ctx.moveTo(edge.a.x -3, edge.a.y-3);
        ctx.lineTo(edge.b.x-3, edge.b.y-3);
        ctx.stroke();
    });
    // Вершины
    ctx.fillStyle = 'blue';
    graph.vertices.forEach(function(vertex) {
        ctx.beginPath();
        ctx.arc(vertex.x-3, vertex.y-3, 2, 0, 2*Math.PI);
        ctx.fill();
    });
}


