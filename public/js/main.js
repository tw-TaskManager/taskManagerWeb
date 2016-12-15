// /*
// function save(content) {
//   var id;
//   $.post('/tasks', {task: content}, function (res, err) {
//     id = res.Id;
//   });
//   return id;
// }
//
// function textArea(content, id) {
//   var node = document.createElement('textarea');
//   node.id = id;
//   node.classList.add('sticky');
//   node.style.width = '29%';
//   node.style.height = '150px';
//   var textNode = document.createTextNode(content);
//   node.appendChild(textNode);
//
//   return node;
// }
//
// // function createRemoveButton(id) {
// //   var deleteButton = document.createElement('button');
// //   deleteButton.id = id;
// //   deleteButton.classList.add('deleteSticky');
// //   return deleteButton;
// // }
//
// function createSticky(data, stickyArea) {
//   var div = document.createElement('div');
//   div.classList.add('stickyArea');
//   stickyArea.prepend(div.appendChild(textArea(data.Task, data.Id)));
// }
//
// function showStickies(dataList) {
//   var stickyArea = $('#stickies')[0];
//   dataList.forEach(function (data) {
//     createSticky(data, stickyArea);
//   });
// }
//
// function getAll() {
//   var data;
//   $.get('/tasks', function (data, status) {
//     data = JSON.parse(data);
//   });
//   return data;
// }
//
// function addSticky() {
//   $('#add').hide();
//   var newBlock = $('#new');
//   var listOfStickies = $('.sticky');
//   var node = textArea('', 'new');
//   var div = document.createElement('div');
//   div.classList.add('new');
//
//   node.style.width = '95%';
//   node.style.height = '300px';
//
//   div.appendChild(node);
//
//   newBlock.append(div);
//   $('#save').show();
// }
//
// function reset() {
//   var newSticky = $('.new').first();
//   newSticky.remove();
//   $('#add').show();
//   $('#save').hide();
//
// }
//
// function filterSticky() {
//   var newSticky;
//   getAll().filter(function () {
//
//   });
//   return newSticky;
// }
//
// function update() {
//   createSticky()
// }
//
// function saveSticky() {
//   var newSticky = $('.new').children();
//   var content = newSticky.val();
//   save(content);
//   update();
// //  showStickies([{Task:content, id: id}]);
//   reset();
// }
//
// function onSave() {
//   saveSticky();
// }
//
// window.load = function () {
//   showStickies(getAll());
// };
//
// */

var stickyArea;
var stickyManager;

var database = {
  getAllStickies(callBack) {
    var stickies;

    $.get('/tasks', function (data, status) {
      callBack(JSON.parse(data));
    });
  },

  saveSticky(content) {
    $.post('/tasks', {content: content}, function (res, err) {
    });
  },

  removeSticky(id) {
    $.post('/task/remove', {id: id}, function (res, err) {
    });
  },

  updateSticky(id) {
    $.post('/task/update', {id: id}, function (res, err) {
    });
  }
};

var dom = {

  stickyDom(content, id) {
    var node = document.createElement('textarea');
    node.id = id;
    node.classList.add('sticky');
    node.style.width = '80%';
    node.style.height = '50px';
    var textNode = document.createTextNode(content);
    node.appendChild(textNode);

    return node;
  }
};

class Sticky {
  constructor(content, id) {
    this.id = id;
    this.content = content;
  }

  show() {
    var div = document.createElement('div');
    div.classList.add('stickyContainer');
    var node = dom.stickyDom(this.content, this.id);
    stickyArea.prepend(div.appendChild(node));
  }
}

class StickyManager {
  constructor() {
    this.stickies = []
  }

  addSticky(id, sticky) {
    this.stickies[id] = sticky;
  }

  addNewSticky() {
    var id = 'new';
    var sticky = new Sticky('', id);
    this.addSticky(id, sticky);
  }

  showSticky(id) {
    var sticky = this.stickies[id];
    sticky.show();
  }
}

var showOnScreen = function (stickiesList) {
  stickiesList.forEach(function (stickyInfo) {
    var id = stickyInfo.Id;
    stickyManager.addSticky(id, new Sticky(stickyInfo.Task, id));
    stickyManager.showSticky(id);
  });
};

var insertNewSticky = function () {
  stickyManager.addNewSticky();
  stickyManager.showSticky('new');
  database.saveSticky();
};

window.onload = function () {
  stickyArea = $('#stickies')[0];
  stickyManager = new StickyManager();

  var stickiesList = database.getAllStickies(showOnScreen);
};