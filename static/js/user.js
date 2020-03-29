function SignUp() {
    var pwd = document.getElementById("SignUpPasswordSure").value;
    if(pwd != document.getElementById("SignUpPassword").value){
        document.getElementById("HintModalMsg").innerHTML = "两次输入的密码不一致";
        $('#HintModal').modal('show');
        return
    }
    $.ajax({
        url: '/user/signup',
        type: 'post',
        data: {
            Password: pwd,
            Username: document.getElementById("SignUpUsername").value,
            Email: document.getElementById("SignUpEmail").value
        },
        success: function(arg){
            document.getElementById("HintModalMsg").innerHTML = "注册成功";
            $('#SignUpModal').modal('hide');
            $('#HintModal').modal('show');
        },error : function() {
            document.getElementById("HintModalMsg").innerHTML = "不知道为啥，注册失败";
            $('#HintModal').modal('show');
        }
    });
}

function ClearHintModalMsg() {
    document.getElementById("HintModalMsg").value = "";
}

function Login() {
    $.ajax({
        url: '/user/login',
        type: 'post',
        data: {
            Password: document.getElementById("LoginPassword").value,
            Email: document.getElementById("LoginEmail").value
        },
        success: function(arg){
            $('#LoginModal').modal('hide');
            SetCookie("Username=" + arg["Username"]);
            ShowUserMenuAfterLogin();
        },error : function() {
            document.getElementById("HintModalMsg").innerHTML = "账号或密码错误";
            $('#HintModal').modal('show');
        }
    });
}

function SetCookie(str) {
    document.cookie = str;
}

function ShowUserMenuAfterLogin(){
    if(document.cookie.Username != ""){
        new Vue({
            el: '#ShowUserMenu',
            data: {IsLogin: true, NotLogin: false, Username: document.cookie.Username}
        })
    }
}

function Init(){
    // debugger
    new Vue({
        el: '#ShowUserMenu',
        data: {IsLogin: false, NotLogin: true}
    });
    // ShowUserMenuAfterLogin();
}

Init();