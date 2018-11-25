
$(function () {

    $('#btn').click(function () {
        var params = {};
        params['from'] = 'test';
        params['data'] = {};

        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/test/ajax/test_ajax_req',
            type: 'POST',
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(params),
            async: 'true',
            dataType: 'json',
            success: function (result) {
                console.log('RESPONSE : ' + JSON.stringify(result));
                console.log("请求成功");
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