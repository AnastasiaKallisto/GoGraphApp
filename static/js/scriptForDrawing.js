class Point {
    constructor(x, y, number) {
        this.X = x;
        this.Y = y;
        this.Number = number;
    }
}

function drawGraph(graph) {
  var canvas = document.getElementById('canvasForGraph');
  var ctx = canvas.getContext('2d');
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  // Ребра
  ctx.strokeStyle = 'black';
  ctx.lineWidth = 2;
  graph.edges.forEach(function(edge) {
    ctx.beginPath();
    ctx.moveTo(edge.a.x, edge.a.y);
    ctx.lineTo(edge.b.x, edge.b.y);
    ctx.stroke();
  });
  // Вершины
  ctx.fillStyle = 'blue';
  graph.vertices.forEach(function(vertex) {
    ctx.beginPath();
    ctx.arc(vertex.x, vertex.y, 5, 0, 2*Math.PI);
    ctx.fill();
  });
}
