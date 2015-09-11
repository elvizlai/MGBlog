/**
 * Created by elvizlai on 15/7/28.
 */
"use strict";
define([], function () {
    require(["jquery", "semantic", "validate"], function ($) {
        //注册表单验证
        $("#registerForm").validate({
            rules: {
                email: {
                    required: true,
                    email: true
                },
                nickname: {
                    required: true,
                    minlength: 5
                },
                password: {
                    required: true,
                    minlength: 6
                },
                repeatpassword: {
                    required: true,
                    equalTo: "#password"
                }
            },
            messages: {
                email: {
                    required: "Email should not be null",
                    email: "Must be an email"
                },
                password: {
                    required: "Password should not be null",
                    minlength: "Password too short"
                }
            },
            errorPlacement: function (error, element) {
                error.appendTo(element.parent().parent());
            }
        });

        //登录表单验证
        $("#loginForm").validate({
            rules: {
                email: {
                    required: true,
                    email: true
                },
                password: {
                    required: true,
                    minlength: 6
                }
            },
            messages: {
                email: {
                    required: "Email should not be null",
                    email: "Must be an email"
                },
                password: {
                    required: "Password should not be null",
                    minlength: "Password too short"
                }
            },
            errorPlacement: function (error, element) {
                error.appendTo(element.parent().parent());
            }
        });
        //注册表单

        //登陆
        $("#loginSubmit").on("click", function () {
            //先判断输入是否合法
            if (!$("#loginForm").valid()) {
                return
            }
            var email = $("#email"),
                password = $("#password"),
                params = {email: email.val(), password: password.val()};
            $.post("/api/login", JSON.stringify(params), function (data) {
                if (data.code === 0) {
                    window.location.href = data.result.url;
                } else {
                    showError(data.msg);
                }
            });
        });

        //注册
        $("#registerSubmit").on("click", function () {
            if (!$("#registerForm").valid()) {
                return
            }
            var params = {
                email: $("#email").val(),
                nickName: $("#nickname").val(),
                password: $("#password").val()
            };
            $.post("/api/register", JSON.stringify(params), function (data) {
                if (data.code === 0) {
                    window.location.href = "/login"
                } else {
                    showError(data.msg);
                }
            });

        });

        //找回密码

        function showError(msg) {
            var $Error = $("#Error");
            $Error.text(msg);
            $Error.parent().show();
        }

        $(".message .close").on("click", function () {
            $(this).closest(".message").fadeOut();
        });
    });
});