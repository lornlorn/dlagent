
$(function () {
    // $(document).tooltip();
});

function showContent(el) {
    // alert($(this).attr("data-id"));
    var title = $(el).children("p[data-jobid]");
    console.log(title.attr("data-jobid"));
}
