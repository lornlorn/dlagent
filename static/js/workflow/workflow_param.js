$(function () {

    window.onunload = function () {
        window.opener.location.reload();
        // window.close();
    }

    // $('#WfiStatus').val("{{ .WfiStatus }}");

});

function GetQuery(param) {
    var url = location.search; //获取url中"?"符后的字串 
    var query = new Object();
    if (url.indexOf("?") != -1) {
        var str = url.substr(1);
        strs = str.split("&");
        for (var i = 0; i < strs.length; i++) {
            query[strs[i].split("=")[0]] = unescape(strs[i].split("=")[1]);
        }
    }
    return query[param];
}
