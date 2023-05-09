//показать пароль
function showPass(target) {
    var input = document.getElementById('pass');
    if (input.getAttribute('type') == 'password') {
        target.classList.add('view');
        input.setAttribute('type', 'text');
    } else {
        target.classList.remove('view');
        input.setAttribute('type', 'password');
    }
    return false;
}

//json
var mail = document.getElementById('email');
var pass = document.getElementById('pass');
var login = document.getElementById('login');

function ka() {
    console.log(JSON.stringify(
        {
            "email": mail.value,
            "password": pass.value
        }))
}
