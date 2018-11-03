
$(function () {
    // $(document).tooltip();

    // 添加指令 
    $("#steplist").on("click", ".stepdtl-add", function () {
        var seq = $(this).parent().parent().attr("data-seq");
        var html =
            "<div class=\"stepdtl\">" +
            " <input data-type=\"check\" data-name=\"stepdtl-check\" type=\"checkbox\">" +
            " <label>IP:</label>" +
            " <input data-type=\"data\" data-name=\"ip\" value=\"192.168.100.101\">" +
            " <label>端口:</label>" +
            " <input data-type=\"data\" data-name=\"port\" value=\"22\">" +
            " <label>用户名:</label>" +
            " <input data-type=\"data\" data-name=\"username\" value=\"test\">" +
            " <label>密码:</label>" +
            " <input data-type=\"data\" data-name=\"password\" value=\"test\">" +
            " <label>命令:</label>" +
            " <input data-type=\"data\" data-name=\"command\" value=\"ls -lrt\">" +
            "</div>";

        $("div.step[data-seq=" + seq + "]").append(html);
    });
});

function showContent(el) {
    var item = $(el).children("p[data-JobId]");
    // console.log(item.attr("data-JobId"));

    var params = {};
    // params['module'] = $module.val(); 
    params['from'] = 'joblist';
    params['data'] = {};
    // $('#json').find('input[name]').each(function () { 
    // var k = $(this).attr('name'); 
    // var v = $(this).val(); 
    // params['data'][k] = v; 
    // }); 
    params['data']['jobtype'] = 'tool';
    params['data']['jobid'] = item.attr("data-JobId");

    $('.content').find('#run').attr("data-JobId", item.attr("data-JobId"));

    console.log('REQUEST : ' + JSON.stringify(params));

    $.ajax({
        url: '/ajax/getjobdtl',
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
    $('#jobflows').empty();

    $('#JobName').val(result['retdata']['job']['JobName']);
    $('#JobRemark').text(result['retdata']['job']['JobRemark']);

    // console.log(result['retdata']['jobflow'].length);
    var jobFlowTitleHtml =
        "<div class=\"field is-grouped\">" +
        "<div class=\"control\">" +
        "<h5 class=\"title is-5\">作业流</h5>" +
        "</div>" +
        "<div class=\"control\">" +
        "<a class=\"button is-primary is-small is-outlined\" id=\"add\" onclick=\"addJobFlow();\">" +
        "<span class=\"icon\">" +
        "<i class=\"fas fa-plus\"></i>" +
        "</span>" +
        "<span>增加作业流</span>" +
        "</a>" +
        "</div>" +
        "</div>";
    $('#jobflows').append(jobFlowTitleHtml);

    $.each(result['retdata']['jobflow'], function (i, v) {
        // console.log(i, v);
        // console.log(result['retdata']['jobflow'][i]['JobflowParam'].length);

        var jobFlowParamHtml = "";
        $.each(result['retdata']['jobflow'][i]['JobflowParam'], function (i, v) {
            // console.log(i, v);
            jobFlowParamHtml = jobFlowParamHtml +
                "<div class=\"column is-one-fifth\">" +
                "<div class=\"parameter\" data-JfpId=\"" + v.JfpId + "\" data-JfpSeq=\"" + v.JfpSeq + "\">" +
                "<div class=\"field has-addons\">" +
                "<div class=\"control\">" +
                "<a class=\"button is-static\">" +
                v.JfpParameter +
                "</a>" +
                "</div>" +
                "<div class=\"control\">" +
                "<input class=\"input\" type=\"text\" placeholder=\"\" id=\"parameter-" + v.JfpId + "\" value=\"" + v.JfpDefault + "\">" +
                "</div>" +
                "</div>" +
                "</div>" +
                "</div>";
        });

        var jobFlowHtml =
            "<div class=\"jobflow\" data-JfId=\"" + v.JfId + "\" data-JfSeq=\"" + v.JfSeq + "\">" +
            "<h6 class=\"title is-6\">序号:" + v.JfSeq + "</h6>" +
            "<div class=\"columns\">" +
            "<div class=\"column is-four-fifths\">" +
            "<div class=\"field has-addons\">" +
            "<div class=\"control\">" +
            "<a class=\"button is-static\">" +
            "名称" +
            "</a>" +
            "</div>" +
            "<div class=\"control is-expanded\">" +
            "<input class=\"input\" type=\"text\" placeholder=\"\" id=\"JfName\" value=\"" + v.JfName + "\">" +
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
            "SHELL" +
            "</a>" +
            "</div>" +
            "<div class=\"control is-expanded\">" +
            "<input class=\"input\" type=\"text\" placeholder=\"\" id=\"JfSh\" value=\"" + v.JfSh + "\">" +
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
            "<input class=\"input\" type=\"text\" placeholder=\"\" id=\"JfCmd\" value=\"" + v.JfCmd + "\">" +
            "</div>" +
            "</div>" +
            "</div>" +
            "</div>" +
            "<div class=\"columns parameters\">" +
            jobFlowParamHtml +
            "</div>" +
            "</div>" +
            "<hr class=\"hr\">";

        $('#jobflows').append(jobFlowHtml);

        $("div.jobflow[data-JfId=" + v.JfId + "]").find("#JfStatus").val(v.JfStatus);
    });
}

function addJobFlow() {
    var len = Number($("#jobflows").find(".jobflow").length);
    var newSeq = len + 1;
    var newFlowHtml =
        "<div class=\"jobflow\" data-JfSeq=\"" + newSeq + "\">" +
        "<h6 class=\"title is-6\">序号:" + newSeq + "</h6>" +
        "<div class=\"columns\">" +
        "<div class=\"column is-four-fifths\">" +
        "<div class=\"field has-addons\">" +
        "<div class=\"control\">" +
        "<a class=\"button is-static\">" +
        "名称" +
        "</a>" +
        "</div>" +
        "<div class=\"control is-expanded\">" +
        "<input class=\"input\" type=\"text\" placeholder=\"\" id=\"JfName\" value=\"\">" +
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
        "<input class=\"input\" type=\"text\" placeholder=\"\" id=\"JfSh\" value=\"\">" +
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
        "<input class=\"input\" type=\"text\" placeholder=\"\" id=\"JfCmd\" value=\"\">" +
        "</div>" +
        "</div>" +
        "</div>" +
        "</div>" +
        "<div class=\"columns parameters\">" +
        // jobFlowParamHtml +
        "</div>" +
        "</div>" +
        "<hr class=\"hr\">";

    $('#jobflows').append(newFlowHtml);
}
