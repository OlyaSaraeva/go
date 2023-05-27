document.addEventListener('DOMContentLoaded', function () {
  console.log('Готов!');
  initEventsListener()
});

function initEventsListener() {
  let buttonEye = document.getElementById('eye')
  console.log(buttonEye)
  buttonEye.addEventListener('click', function () {
    showPass(buttonEye);
  })

  let login = document.getElementById('login')
  console.log(login)
  login.addEventListener(
    'click', function () {
      showJSON();
      showError()
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

function showError() {
  var mail = document.getElementById('email');
  var pass = document.getElementById('pass');
  let errorMail = document.getElementById('errorEmailLabel')
  let errorPass = document.getElementById('errorPassLabel')

  if (mail.value !== null && mail.value == '') {
    console.log("email=" + mail.value)
    errorMail.classList.remove('error')
    errorMail.classList.add('error-empty')
    errorMail.style.color = '#E86961'
    mail.style.borderColor = "#E86961";
    mail.style.marginBottom = '5px';
    allNoOk()
  }
  else {
    controlEmail(mail, errorMail)
  }

  if (pass.value !== null && pass.value == '') {
    errorPass.classList.remove('error')
    errorPass.classList.add('error-empty')
    errorPass.style.color = '#E86961'
    pass.style.borderColor = "#E86961";
    pass.style.marginBottom = '5px';
    allNoOk()
  }
  else {
    hiddenError(pass, errorPass)
  }
}

function hiddenError(field, label) {
  field.style.borderColor = '#7A7A7A';
  label.style.color = '#7A7A7A';
}

function controlEmail(mail, errorMail) {
  var format = /^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/;
  console.log(mail.value)
  if (format.test(mail.value)) {
    //console.log('yes')
    hiddenError(mail, errorMail)
    allOk()
    return true;
  } else {
    showErrorIncorrectMail(mail, errorMail)
    //console.log('no')
    return false;
  }
}

function showErrorIncorrectMail(mail, errorMail) {
  let incorr = document.getElementById('incorrect');
  let corr = document.getElementById('correct')
  corr.classList.remove('correct')
  corr.classList.add('form-block__error')
  incorr.classList.remove('form-block__error')
  incorr.classList.add('incorrect')
  incorr.classList.remove('correct')
  errorMail.classList.remove('error')
  errorMail.classList.add('error-empty')
  errorMail.style.color = '#E86961'
  mail.style.borderColor = "#E86961";
  mail.style.marginBottom = '5px';
  incorr.textContent = 'Email or password is incorrect.'
  errorMail.textContent = "Incorrect email format. Correct format is ****@**.***";
}

function allOk() {
  let incorr = document.getElementById('incorrect');
  let corr = document.getElementById('correct')
  corr.classList.add('correct')
  incorr.classList.remove('incorrect')
  incorr.classList.add('form-block__error')
  corr.textContent = "All ok";
}

function allNoOk() {
  let incorr = document.getElementById('incorrect');
  let corr = document.getElementById('correct')
  corr.classList.remove('correct')
  corr.classList.add('form-block__error')
  incorr.classList.remove('form-block__error')
  incorr.classList.add('incorrect')
  incorr.textContent = 'A-Ah! Check all fields'
}