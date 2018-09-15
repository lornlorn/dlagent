
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
        var html =
            "<div class=\"jobflow\" data-JfId=\"" + result['retdata']['jobflow'][i]['JfId'] + "\" data-JfSeq=\"" + result['retdata']['jobflow'][i]['JfSeq'] + "\">" +
            "<h6 class=\"title is-6\">序号:" + result['retdata']['jobflow'][i]['JfSeq'] + "</h6>" +
            "<div class=\"columns\">" +
            "<div class=\"column is-four-fifths\">" +
            "<div class=\"field has-addons\">" +
            "<div class=\"control\">" +
            "<a class=\"button is-static\">" +
            "名称" +
            "</a>" +
            "</div>" +
            "<div class=\"control is-expanded\">" +
            "<input class=\"input\" type=\"text\" placeholder=\"\" id=\"JfName\" value=\"" + result['retdata']['jobflow'][i]['JfName'] + "\">" +
            "</div>" +
            "</div>" +
            "</div>" +
            "<div class=\"column\">" +
            "<div class=\"field has-addons\">" +
            "<div class=\"control\">" +
            "<a class=\"button is-static\">" +
            "状态" +
            "</a>" +
            "</div>" +
            "<div class=\"control is-expanded\">" +
            "<div class=\"select\">" +
            "<select id=\"JfStatus\">" +
            "<option>启用</option>" +
            "<option>停用</option>" +
            "</select>" +
            "</div>" +
            "</div>" +
            "</div>" +
            "</div>" +
            "</div>" +
            "<div class=\"columns\">" +
            "<div class=\"column is-one-fifth\">" +
            "<div class=\"field has-addons\">" +
            "<div class=\"control\">" +
            "<a class=\"button is-static\">" +
            "SHELL路径" +
            "</a>" +
            "</div>" +
            "<div class=\"control is-expanded\">" +
            "<input class=\"input\" type=\"text\" placeholder=\"\" id=\"JfSh\" value=\"" + result['retdata']['jobflow'][i]['JfSh'] + "\">" +
            "</div>" +
            "</div>" +
            "</div>" +
            "<div class=\"column\">" +
            "<div class=\"field has-addons\">" +
            "<div class=\"control\">" +
            "<a class=\"button is-static\">" +
            "命令" +
            "</a>" +
            "</div>" +
            "<div class=\"control is-expanded\">" +
            "<input class=\"input\" type=\"text\" placeholder=\"\" id=\"JfCmd\" value=\"" + result['retdata']['jobflow'][i]['JfCmd'] + "\">" +
            "</div>" +
            "</div>" +
            "</div>" +
            "</div>" +
            "<div class=\"columns parameters\">" +
            "<div class=\"column is-one-fifth\">" +
            "<div class=\"field has-addons\">" +
            "<div class=\"control\">" +
            "<a class=\"button is-static\">" +
            "参数1" +
            "</a>" +
            "</div>" +
            "<div class=\"control\">" +
            "<input class=\"input\" type=\"text\" placeholder=\"\" id=\"JfpParameter1\">" +
            "</div>" +
            "</div>" +
            "</div>" +
            "<div class=\"column is-one-fifth\">" +
            "<div class=\"field has-addons\">" +
            "<div class=\"control\">" +
            "<a class=\"button is-static\">" +
            "参数2" +
            "</a>" +
            "</div>" +
            "<div class=\"control\">" +
            "<input class=\"input\" type=\"text\" placeholder=\"\" id=\"JfpParameter2\">" +
            "</div>" +
            "</div>" +
            "</div>" +
            "<div class=\"column is-one-fifth\">" +
            "<div class=\"field has-addons\">" +
            "<div class=\"control\">" +
            "<a class=\"button is-static\">" +
            "参数3" +
            "</a>" +
            "</div>" +
            "<div class=\"control\">" +
            "<input class=\"input\" type=\"text\" placeholder=\"\" id=\"JfpParameter3\">" +
            "</div>" +
            "</div>" +
            "</div>" +
            "<div class=\"column is-one-fifth\">" +
            "<div class=\"field has-addons\">" +
            "<div class=\"control\">" +
            "<a class=\"button is-static\">" +
            "参数4" +
            "</a>" +
            "</div>" +
            "<div class=\"control\">" +
            "<input class=\"input\" type=\"text\" placeholder=\"\" id=\"JfpParameter4\">" +
            "</div>" +
            "</div>" +
            "</div>" +
            "<div class=\"column is-one-fifth\">" +
            "<div class=\"field has-addons\">" +
            "<div class=\"control\">" +
            "<a class=\"button is-static\">" +
            "参数5" +
            "</a>" +
            "</div>" +
            "<div class=\"control\">" +
            "<input class=\"input\" type=\"text\" placeholder=\"\" id=\"JfpParameter5\">" +
            "</div>" +
            "</div>" +
            "</div>" +
            "</div>" +
            "</div>" +
            "<hr class=\"hr\">";

        $('#jobflows').append(html);
    });

    // console.log($('#JfStatus').val());
    // console.log($('#JfStatus').text());
    $('#JfStatus').val('停用');

    // $('#content').empty();


}
