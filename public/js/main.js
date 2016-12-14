function save(content) {
  var id;
  $.post('/tasks', {task: content}, function (res, err) {
    id = res.Id;
  })
  return id;
}

function textArea(content, id) {
  var node = document.createElement('textarea');
  node.id = id;
  node.classList.add('sticky');
  node.style.width = '29%';
  node.style.height = '150px';
  var textNode = document.createTextNode(content);
  node.appendChild(textNode);

  return node;
}


// function createRemoveButton(id) {
//   var deleteButton = document.createElement('button');
//   deleteButton.id = id;
//   deleteButton.classList.add('deleteSticky');
//   return deleteButton;
// }


function createSticky(data, stickyContainer){
    var div = document.createElement('div');
    div.classList.add('stickyContainer');
    stickyContainer.prepend(div.appendChild(textArea(data.Task, data.Id)));
}

function showStickies(dataList) {
  var stickyContainer = $('#stickies')[0];
  dataList.forEach(function (data) {
    createSticky(data, stickyContainer);
  });
}

function getAll() {
  $.get('/tasks', function (data, status) {
    data=JSON.parse(data)
    showStickies(data);
  })
}

function addSticky() {
  $('#add').hide();
  var newBlock = $('#new');
  var listOfStickies = $('.sticky');
  var node = textArea('', 'new');
  var div = document.createElement('div');
  div.classList.add('new');

  node.style.width = '95%';
  node.style.height = '300px';

  div.appendChild(node);

  newBlock.append(div);
  $('#save').show();
}

function reset() {
  var newSticky = $('.new').first();
  newSticky.remove();
  $('#add').show();
  $('#save').hide();

}

function filterSticky() {
    var newSticky;
    getAll().filter(function(){

    })
    return newSticky;
}

function update() {
    createSticky()
}

function saveSticky() {
  var newSticky = $('.new').children();
  var content = newSticky.val();
  save(content);
  update()
//  showStickies([{Task:content, id: id}]);
  reset();
}

window.load = allTask();

