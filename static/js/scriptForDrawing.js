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

function drawGraph(graph) {
  var canvas = document.getElementById('canvasForGraph');
  var ctx = canvas.getContext('2d');
  // Ребра
  // ctx.strokeStyle = 'black';
  // ctx.lineWidth = 2;
  // graph.edges.forEach(function(edge) {
  //   ctx.beginPath();
  //   ctx.moveTo(edge.a.x, edge.a.y);
  //   ctx.lineTo(edge.b.x, edge.b.y);
  //   ctx.stroke();
  // });
  // // Вершины
  // ctx.fillStyle = 'blue';
  // graph.vertices.forEach(function(vertex) {
  //   ctx.beginPath();
  //   ctx.arc(vertex.x, vertex.y, 5, 0, 2*Math.PI);
  //   ctx.fill();
  // });
    ctx.beginPath();
    ctx.fillStyle = 'red';
    ctx.fillRect(50, 50, 100, 100);
    ctx.stroke();
}
