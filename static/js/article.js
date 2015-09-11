/**
 * Created by elvizlai on 15/8/25.
 */
"use strict";
define([], function () {
    var deps = [
        "jquery",
        "editormd",
        "validate",
        "semantic",
        "/static/js/console.js",
        "/static/plugins/link-dialog/link-dialog.js",
        "/static/plugins/reference-link-dialog/reference-link-dialog.js",
        "/static/plugins/image-dialog/image-dialog.js",
        "/static/plugins/code-block-dialog/code-block-dialog.js",
        "/static/plugins/table-dialog/table-dialog.js",
        "/static/plugins/emoji-dialog/emoji-dialog.js",
        "/static/plugins/goto-line-dialog/goto-line-dialog.js",
        "/static/plugins/help-dialog/help-dialog.js",
        "/static/plugins/html-entities-dialog/html-entities-dialog.js",
        "/static/plugins/preformatted-text-dialog/preformatted-text-dialog.js"
    ];

    require(deps, function ($, editormd) {
        //设置401接收
        $.ajaxSetup({
            statusCode: {
                401: function () {
                    $('.popLogin').modal('show');
                }
            }
        });

        //发布表单验证
        $(".validate").validate({
            rules: {
                title: {
                    required: true
                }
            },
            messages: {
                title: {
                    required: "Title should not be null"
                }
            },
            errorPlacement: function (error, element) {
                element.attr("placeholder", "Title should not be null");
            }
        });

        if ($("#editor").length > 0) {
            editormd.loadCSS("/static/codemirror/addon/fold/foldgutter");
            var editor = editormd("editor", {
                width: "90%",
                height: 720,
                syncScrolling: "single",
                path: "/static/editormd/",
                toolbarIcons: function () {
                    return ["undo", "redo", "|", "bold", "del", "italic", "quote", "ucwords", "|", "h1", "h2", "h3", "h4", "h5", "h6", "|", "list-ul", "list-ol", "hr", "|", "link", "image", "code", "code-block", "table", "datetime", "emoji", "html-entities", "pagebreak", "|", "preview", "search"]
                },
                imageUpload: true,
                imageFormats: ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
                imageUploadURL: "/upload",
                codeFold: true,//代码折叠
                searchReplace: true,
                gfm: true,
                emoji: true,
                taskList: true,
                flowChart: true,
                tex: true,
                sequenceDiagram: true,
                htmlDecode: "style,script,iframe,sub,sup|on*"  // Filter tags, and all on* attributes
            });

            editor.setToolbarAutoFixed(true);

            //发布
            $("#pubBtn").on("click", function () {
                if (!$(".validate").valid()) {
                    return
                }

                var title = $("#title");
                var tags = $("#tags");
                var markdown = editor.getMarkdown();
                var htmlContent = editor.getPreviewedHTML();
                var params = {title: title.val(), tags: tags.val(), markdown: markdown, htmlContent: htmlContent};

                $.post("/newArticle", JSON.stringify(params), function (data) {
                    if (data.code === 0) {
                        location.href = "/"
                    } else {
                        alert(data.code + " " + data.msg)
                    }
                });
            });

            //修改
            $("#modifyBtn").on("click", function () {
                if (!$(".validate").valid()) {
                    return
                }

                var title = $("#title");
                var tags = $("#tags");
                var markdown = editor.getMarkdown();
                var htmlContent = editor.getPreviewedHTML();
                var params = {
                    method: "update",
                    title: title.val(),
                    tags: tags.val(),
                    markdown: markdown,
                    htmlContent: htmlContent
                };

                $.post(window.location.pathname, JSON.stringify(params), function (data) {
                    if (data.code === 0) {
                        location.href = "/article/" + window.location.pathname.substr(15);
                    } else {
                        alert(data.code + " " + data.msg);
                    }
                });
            });

            //删除
            $("#deleteBtn").on("click", function () {
                $('.delConfirm').modal({
                    closable: false,
                    onApprove: function () {
                        var params = {method: "delete"};
                        $.post(window.location.pathname, JSON.stringify(params), function (data) {
                            if (data.code === 0) {
                                location.href = "/trash";//回收站
                            } else {
                                alert(data.code + " " + data.msg);
                            }
                        });
                    }
                }).modal('show');
            });
        }

        //登录
        $("#login").on("click", function () {
            //先判断输入是否合法
            if (!$("#loginForm").valid()) {
                return
            }
            var email = $("#email"),
                password = $("#password"),
                params = {email: email.val(), password: password.val()};
            $.post("/api/login", JSON.stringify(params), function (data) {
                if (data.code === 0) {
                    $('.popLogin').modal('toggle');
                } else {
                    showError(data.msg);
                }
            });
        });

        //删除文章恢复
        $(".rollback").on("click", function () {
            var $url = $(this).parent().parent().find("a").attr("href");
            $('.rollbackConfirm').modal({
                closable: false,
                onApprove: function () {
                    $.post($url, "", function (data) {
                        if (data.code === 0) {
                            location.href = "/trash";//回收站
                        } else {
                            alert(data.code + " " + data.msg);
                        }
                    });
                }
            }).modal('show');
        });

        function showError(msg) {
            var $Error = $("#Error");
            $Error.text(msg);
            $Error.parent().show();
        }

    });
});