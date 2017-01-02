function alreadyLogin(){
    $.get("/task/isUserAlreadyLogin",function(err,status){})
}

window.load = alreadyLogin();