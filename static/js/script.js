function copyValueTo() {
  // получаем форму
  let form = document.forms.form; // <form name="my"> element

  // получаем элемент
  let title = form.elements.title; // <input name="one"> element
  let subtitle = form.elements.subtitle;
  let author = form.elements.author;
  let publishDate = form.elements.publishDate;

  console.log(form)
  console.log(title)
  /* if (title.value != null) {
      title.classList.add("active");
  console.log();
  } 
  if (title.value == null) {
      title.classList.remove("active");
      console.log(elem)
  } */

  let titlea = form.elements.titlea;
  let subtitlea = form.elements.subtitlea;
  let authorPost = form.elements.authorPost;
  let datePost = form.elements.datePost;

  titlea.value = title.value;
  subtitlea.value = subtitle.value;
  authorPost.value = author.value;

  let titlep = form.elements.titlep;
  let subtitlep = form.elements.subtitlep;

  datePost.value = publishDate.value;


  titlep.value = title.value;
  subtitlep.value = subtitle.value;

  // console.log(publishDate.value);
}

function previewFile(photo, post, prev) {
  const preview = document.getElementById(post);
  // console.log(preview);
  const previewphoto = document.getElementById(prev);
  console.log(previewphoto);
  const file = document.getElementById(photo).files[0];
  //console.log(file);
  const reader = new FileReader();
  //console.log(reader);

  //const prevblock=document.getElementsByClassName('input-file__load-block');

  reader.addEventListener(
    "load",
    () => {
      // convert image file to base64 string
      preview.src = reader.result;
      previewphoto.src = preview.src;
      // console.log(preview.src);
    },
    false
  );

  if (file) {
    reader.readAsDataURL(file);
    // console.log(reader);
  }
  console.log(prev);
  if (prev == "prevArticleImage") {
    console.log(22222);
    previewphoto.classList.add("preview-artical");
  } else if (prev == "prevHeroImagePost") {
    previewphoto.classList.add("preview-post");
  }
}

//переименование upload
function reclass() {
  const previewphoto = document.getElementById('prevAuthor');
  let photo = document.getElementById('uploadPhoto');

  previewphoto.classList.remove("input-file__button");
  previewphoto.classList.add("input-file__button-new");
  photo.classList.remove("input-file__button");
  photo.classList.add("input-file__button-upload");
}

function rename(i) {
  let a = document.getElementsByClassName('input-file__upload')[i].textContent = "Upload New";  
}

function deletename(i) {
  let a = document.getElementsByClassName('input-file__upload')[i].textContent = "Upload";  
}

//изменение блока
function changeBlock(lablePar, block, rem) {

  let b = document.getElementById(lablePar);
  if (lablePar == "labelAuthor") {
    console.log(lablePar);
    reclass();
    let i=0;
    rename(i);
  } else {
    b.classList.toggle("input-file__img-block");
    if (lablePar == "labelAticle") {
      i=1;
      rename(i);
    }
    else {
      i=2;
      rename(2);
    }
  }
  
  let remB = document.getElementById(rem);
  console.log(11111111)
  remB.classList.toggle("form-block__remove");
  remB.classList.toggle("form-block__remove_hidden");

  let blocks = document.getElementById(block);
  blocks.classList.toggle("form-block__block");
}

//удаление превью
function removePost(two, tree, block, rem, prevPost) {
  let removePosts = document.getElementById(two);
  removePosts.src = '';

  if (two == "prevArticleImage") {
    console.log(333333);
    removePosts.classList.remove("preview-artical");
  } else if (two == "prevHeroImagePost") {
    removePosts.classList.remove("preview-post");
  }

  console.log(removePosts.value);
  changeBlock(tree, block, rem);
  console.log(rem);

  let prevPosts = document.getElementById(prevPost);
  prevPosts.src = '';

  if (rem == "removeBlockAuthor") {
    let i=0;
    deletename(i);
  } else {
    if (rem == "removeBlock") {
      i=1;
      deletename(i);
    }
    else {
      i=2;
      deletename(2);
    }
  }
}