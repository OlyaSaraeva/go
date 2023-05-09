function copyValueTo() {
  // получаем форму
  let form = document.forms.form; // <form name="my"> element

  // получаем элемент
  let title = form.elements.title; // <input name="one"> element
  let subtitle = form.elements.subtitle;
  let author = form.elements.author;
  let publishDate = form.elements.publishDate;

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
}

function previewFile(photo, post, prev) {
  const preview = document.getElementById(post);
  const previewphoto = document.getElementById(prev);
  const file = document.getElementById(photo).files[0];
  const reader = new FileReader();

  if (/\.(jpe?g|png|gif)$/i.test(file.name)) {
    reader.addEventListener(
      "load",
      () => {
        // convert image file to base64 string
        preview.src = reader.result;
        previewphoto.src = preview.src;
      },
      false
    );
  }

  if (file) {
    reader.readAsDataURL(file);
  }
  if (prev == "prevArticleImage") {
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
  photo.classList.toggle("input-file__button");
  photo.classList.toggle("input-file__button-upload");
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
    reclass();
    let i = 0;
    rename(i);
  } else {
    b.classList.toggle("input-file__img-block");
    if (lablePar == "labelAticle") {
      i = 1;
      rename(i);
    }
    else {
      i = 2;
      rename(2);
    }
  }

  let remB = document.getElementById(rem);

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
    removePosts.classList.remove("preview-artical");
  } else if (two == "prevHeroImagePost") {
    removePosts.classList.remove("preview-post");
  }

  changeBlock(tree, block, rem);

  let prevPosts = document.getElementById(prevPost);
  prevPosts.src = '';

  if (rem == "removeBlockAuthor") {
    let i = 0;
    deletename(i);
  } else {
    if (rem == "removeBlock") {
      i = 1;
      deletename(i);
    }
    else {
      i = 2;
      deletename(2);
    }
  }
}

//JSON
var title = document.getElementById('title');
var subtitle = document.getElementById('subtitle');
var author = document.getElementById('author');
var publishDate = document.getElementById('start');
var authorImg = document.getElementById('postAuthorPhoto');
var postImg = document.getElementById('postImage');
var content = document.getElementById('content')

var publish = document.getElementById('publish')

publish.addEventListener(
  "click", function () {
    var data = {
      "Title": title.value,
      "Subtitle": subtitle.value,
      "Author": author.value,
      "PublishDate": publishDate.value,
      "Content": content.value,
      "authorImg": authorImg.src,
      "postImg": postImg.src
    }
    console.log(JSON.stringify(data,null, 2))
  }
);