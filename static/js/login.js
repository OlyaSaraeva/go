document.addEventListener('DOMContentLoaded', function(){
	console.log('Готов!');
    initEventsListener()
});

function initEventsListener()
{
  let buttonEye = document.getElementById('eye')
  buttonEye.addEventListener('click', function() {
    showPass(buttonEye);
  })

  let login = document.getElementById('login')
  login.addEventListener('click', function() {
    showJSON();
  })
}


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

function showJSON() {
    var mail = document.getElementById('email');
    var pass = document.getElementById('pass');
    console.log(JSON.stringify(
        {
            "email": mail.value,
            "password": pass.value
        }))
}
