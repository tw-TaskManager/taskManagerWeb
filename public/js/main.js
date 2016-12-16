var stickyArea;
var stickyManager;

var database = {
  getAllStickies(callBack) {
    var stickies;

    $.get('/tasks', function (data, status) {
      callBack(JSON.parse(data));
    });
  },

  saveSticky(content, callBack) {
    $.post('/tasks/save', {content: content}, function (res, err) {
        callBack(res)
    });
  },

  removeSticky(id) {
    $.post('/task/delete', {id: id}, function (res, err) {
    });
  },

  updateSticky(id, content) {
    $.post('/task/update', {id: id, content: content}, function (res, err) {
    });
  }
};

var dom = {

  stickyDom(content, id) {
    var node = document.createElement('textarea');
    node.id = id;
    node.classList.add('sticky');
    node.style.width = '60%';
    node.style.height = '40px';
    node.addEventListener("focusout", setFocused);
    var textNode = document.createTextNode(content);
    node.appendChild(textNode);
    return node;
  }
};


function Sticky(content, id){
    this.id = id;
    this.content = content;
}


Sticky.prototype={
  show:function() {
      var div = document.createElement('div');
      var node = dom.stickyDom(this.content, this.id)
      div.appendChild(node)
      div.appendChild(this.createDeleteButton())
      stickyArea.prepend(div);

   },

   createDeleteButton(){
         var deleteButton = document.createElement('button');
         deleteButton.id = this.id;
         deleteButton.classList.add('delete')
         deleteButton.innerHTML = 'Delete Task'
         deleteButton.onclick = stickyManager.deleteSticky.bind(null, database.removeSticky)
         return deleteButton;
   }
}

function StickyManager (){
    this.stickies = []
}

StickyManager.prototype = {
  addSticky:function(id, sticky) {
      this.stickies[id] = sticky;
  },
   showSticky(id) {
       var sticky = this.stickies[id];
       sticky.show();
   },
   deleteSticky(callBack, event) {
    var id = Number(event.target.id);
    $('#'+id).remove();
    event.target.remove();
    callBack(id)
   }
}


var showOnScreen = function (stickiesList) {
  stickiesList.forEach(function (stickyInfo) {
    var id = stickyInfo.Id;
    var sticky = new Sticky(stickyInfo.Task, id);
    stickyManager.addSticky(id, sticky);
    stickyManager.showSticky(id);
  });
};

var insertNewSticky = function () {
  database.saveSticky('', function(id){
    var sticky = new Sticky('', id);
    stickyManager.addSticky(id, sticky);
    stickyManager.showSticky(id);
  })
};

var setFocused = function(event) {
    var id = event.target.id;
    var content = $('#'+id).val();
    database.updateSticky(id, content)
}

$(document).ready(function(){
stickyArea = $('#stickies')[0];
    stickyManager = new StickyManager();
    database.getAllStickies(showOnScreen);
})

