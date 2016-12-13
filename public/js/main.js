function save(content) {
  $.post('/tasks', {task: content}, function (res, err) {
  })
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

function createStickies(contentList, id) {
  var stickyContainer = $('#stickies')[0];
  contentList.forEach(function (content, index = id || index) {
    var div = document.createElement('div');
    div.classList.add('stickyContainer');
    stickyContainer.prepend(div.appendChild(textArea(content, index)));
  });
}

function allTask() {
  $.get('/tasks', function (res, err) {
    res = res.split('<br/>');
    res.length -= 1;
    createStickies(res);
  })
}

function addSticky() {
  $('#add').hide();
  var newBlock = $('#new');
  var listOfStickies = $('.sticky');
  var id = listOfStickies.length != 0 ? Number(listOfStickies.first()[0].id) + 1 : 1;
  var node = textArea('', id);
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

function saveSticky() {
  var newSticky = $('.new').children();
  var content = newSticky.val();
  createStickies([content], newSticky.id);
  debugger;
  save(content);
  reset();
}

window.load = allTask();

