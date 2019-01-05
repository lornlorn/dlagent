$(function () {

    window.onunload = function () {
        window.opener.location.reload();
        // window.close();
    }

    $('#submit').click(function () {
        var params = {};
        params['from'] = 'ComponentAdd';
        params['data'] = {};
        params['data']['CompName'] = $('#CompName').val();
        params['data']['CompCmd'] = $('#CompCmd').val();

        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/ajax/NewComponent',
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
