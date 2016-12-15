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
      stickyArea.prepend(div.appendChild(node));
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
  database.saveSticky('', function(id){
    stickyManager.addSticky(id, new Sticky('', id));
    stickyManager.showSticky(id);
  })
};

// call this function when cursor is out of focus.
var setFocused = function(){
    $('.sticky').focusout(function(event){
    var id = event.target.id;
    var content = $('#'+id).val();
    database.updateSticky(id, content);

    })
}

$(document).ready(function(){
stickyArea = $('#stickies')[0];
    stickyManager = new StickyManager();
    database.getAllStickies(showOnScreen);
    setTimeout(setFocused, 10);
})