
$(function () {
    // $(document).tooltip();
});

function showContent(el) {
    var item = $(el).children("p[data-JobId]");
    // console.log(item.attr("data-JobId"));

    var params = {};
    // params['module'] = $module.val(); 
    params['module'] = 'joblist';
    params['data'] = {};
    // $('#json').find('input[name]').each(function () { 
    // var k = $(this).attr('name'); 
    // var v = $(this).val(); 
    // params['data'][k] = v; 
    // }); 
    params['data']['jobtype'] = 'tool';
    params['data']['jobid'] = item.attr("data-jobid");

    console.log('REQUEST : ' + JSON.stringify(params));

    $.ajax({
        url: '/test/ajax/getjobdtl',
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
            drawContent(result);

        },
        error: function (result) {
            // $('#status').text('请求失败'); 
            console.log("请求失败");
        },
        complete: function () {
            console.log("Ajax finish");
        },
    });

}

function drawContent(result) {
    $('#JobName').val(result['retdata']['job']['JobName']);
    $('#JobRemark').text(result['retdata']['job']['JobRemark']);

    console.log(result['retdata']['jobflow'].length);

    $.each(result['retdata']['jobflow'], function (i, v) {
        console.log(i, v);
        console.log(result['retdata']['jobflow'][i]['JobflowParam'].length);
    });

    // console.log($('#JfStatus').val());
    // console.log($('#JfStatus').text());
    $('#JfStatus').val('停用');

    // $('#content').empty();
}
