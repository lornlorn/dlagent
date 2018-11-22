
$(function () {

    $('#btn').click(function () {
        var params = {};
        // params['module'] = $module.val(); 
        params['from'] = 'test';
        params['data'] = {};
        // $('#json').find('input[name]').each(function () { 
        // var k = $(this).attr('name'); 
        // var v = $(this).val(); 
        // params['data'][k] = v; 
        // }); 

        /*
        params['data']['shell'] = $('#shell').val();
        params['data']['cmd'] = $('#cmd').val();
        */

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
                // $('#status').text('请求成功'); 
                console.log("请求成功");
                // $('#result').text(result['retcode'] + '|' + result['retmsg']); 
            },
            error: function (result) {
                // $('#status').text('请求失败'); 
                console.log("请求失败");
            },
            complete: function () {
                console.log("Ajax finish");
            },
        });
    });

}); 