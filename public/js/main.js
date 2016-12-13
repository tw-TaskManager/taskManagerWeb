function save() {
    var toSave = $("#save input").val();
    if (toSave != "") {
        $.post("/tasks", {task: toSave}, function (res, err) {
            $('#tasks').append(toSave + "<br>");
        })
    }
}

function allTask() {
    $.get("/tasks", function (res, err) {
        $('#tasks').html(res);
    })
}
window.load = allTask()

