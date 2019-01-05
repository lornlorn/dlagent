$(function () {

    window.onunload = function () {
        window.opener.location.reload();
        // window.close();
    }

    $('#submit').click(function () {
        var params = {};
        params['from'] = 'ParameterAdd';
        params['data'] = {};
        params['data']['CompId'] = GetQuery('CompId');
        params['data']['ParamSeq'] = $('#ParamSeq').val();
        params['data']['ParamName'] = $('#ParamName').val();
        params['data']['ParamDefault'] = $('#ParamDefault').val();

        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/ajax/NewParameter',
            type: 'POST',
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(params),
            async: 'true',
            dataType: 'json',
            success: function (result) {
                console.log('RESPONSE : ' + JSON.stringify(result));
                console.log("请求成功");
                alert('成功');
                window.close();
            },
            error: function (result) {
                console.log("请求失败");
            },
            complete: function () {
                console.log("Ajax finish");
            },
        });
    });

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