document.addEventListener('DOMContentLoaded', function () {
  console.log('Готов!');
  initEventsListener();
});

function initEventsListener() {
  let form = document.forms.form;

  let title = form.elements.title;
  title.addEventListener('keyup', function () {
    copyValueTitle();
  })
  let subtitle = form.elements.subtitle;
  subtitle.addEventListener('keyup', function () {
    copyValueSubtitle();
  })
  let author = form.elements.author;
  author.addEventListener('keyup', function () {
    copyValueAuthorName();
  })
  let publishDate = form.elements.publishDate;
  publishDate.addEventListener('change', function () {
    copyValuePublishDate();
  })

  let uploadAuthorPhoto = document.getElementById('labelAuthor');
  let photo = document.getElementById('photo');
  let postAuthorPhoto = document.getElementById('postAuthorPhoto');
  let prevAuthor = document.getElementById('prevAuthor');
  let blockAuthor = document.getElementById('blockAuthor');
  let removeBlockAuthor = document.getElementById('removeBlockAuthor');
  uploadAuthorPhoto.addEventListener('change', function () {
    previewFile(photo, postAuthorPhoto, prevAuthor, imgData);
    changeBlock(uploadAuthorPhoto, blockAuthor, removeBlockAuthor);
  })

  let uploadAticleImg = document.getElementById('labelAticle');
  let heroImage = document.getElementById('heroImage');
  let articleImage = document.getElementById('articleImage');
  let prevArticleImage = document.getElementById('prevArticleImage');
  let block = document.getElementById('block');
  let removeBlock = document.getElementById('removeBlock');
  uploadAticleImg.addEventListener('change', function () {
    previewFile(heroImage, articleImage, prevArticleImage, imgData);
    changeBlock(uploadAticleImg, block, removeBlock);
  })

  let uploadPostImg = document.getElementById('labelPost');
  let heroImagePost = document.getElementById('heroImagePost');
  let postImage = document.getElementById('postImage');
  let prevHeroImagePost = document.getElementById('prevHeroImagePost');
  let blockPost = document.getElementById('blockPost');
  let removeBlockPost = document.getElementById('removeBlockPost');
  uploadPostImg.addEventListener('change', function () {
    previewFile(heroImagePost, postImage, prevHeroImagePost, imgData)
    changeBlock(uploadPostImg, blockPost, removeBlockPost);
  })

  let authorRemove = document.getElementById('authorRemove');
  authorRemove.addEventListener('click', function () {
    removePost(prevAuthor, uploadAuthorPhoto, blockAuthor, removeBlockAuthor, postAuthorPhoto)
  })

  let articleRemove = document.getElementById('articleRemove');
  articleRemove.addEventListener('click', function () {

    removePost(prevArticleImage, uploadAticleImg, block, removeBlock, articleImage)
  })

  let postRemove = document.getElementById('postRemove');
  postRemove.addEventListener('click', function () {

    removePost(prevHeroImagePost, uploadPostImg, blockPost, removeBlockPost, postImage)
  })


  let errorTitle = document.getElementById('errorTitle');
  let errorSubtitle = document.getElementById('errorSubtitle');
  let errorAuthorName = document.getElementById('errorAuthorName');
  let errorPublishDate = document.getElementById('errorPublishDate');
  let datePostValue = document.getElementById('start')

  //JSON
  let publish = document.getElementById('publish')
  let content = document.getElementById('content')
  const imgData = {
    fileAuthorName: '',
    filePostName: ''
  }

  const CreateNewPost = async function () {
    const response = await fetch('api/post', {
      method: 'POST',
      body: JSON.stringify({
        title: title.value,
        subtitle: subtitle.value,
        author: author.value,
        authorImg: postAuthorPhoto.src,
        publishDate: publishDate.value,
        postImg: postImage.src,
        content: content.value,
        authorImgName: imgData.fileAuthorName,
        postImgName: imgData.filePostName,
      }
      )
    })
    console.log(response.ok)
  }

  publish.addEventListener(
    "click", function () {
      dateValue = valueEmpty(publishDate, errorPublishDate)
     if (checkAllValue()){
      CreateNewPost()
     }
    }
  )


  /* publish.addEventListener(
    "click", function () {
      var data = {
        "Title": title.value,
        "Subtitle": subtitle.value,
        "Author": author.value,
        "PublishDate": publishDate.value,
        "Content": content.value,
        "authorImg": postAuthorPhoto.src,
        "postImg": postImage.src,
        "authorImgName": imgData.fileAuthorName,
        "postImgName": imgData.filePostName,
      }
      console.log(JSON.stringify(data, null, 2))
    }
  ); */
}

function copyValueTitle() {
  let titleArticle = form.elements.titleArticle;
  let titlePost = form.elements.titlePost;

  titleArticle.value = title.value;
  titlePost.value = title.value;
}


function copyValueSubtitle() {
  let subtitleArticle = form.elements.subtitleArticle;
  let subtitlePost = form.elements.subtitlePost;

  subtitleArticle.value = subtitle.value;
  subtitlePost.value = subtitle.value;
}

function copyValueAuthorName() {
  let authorPost = form.elements.authorPost;

  authorPost.value = author.value;
}

function copyValuePublishDate() {
  let publishDate = form.elements.publishDate;
  let datePost = form.elements.datePost;
  datePost.value = publishDate.value;
}

function previewFile(photo, post, prev, imgData) {
  const preview = post;
  const previewphoto = prev;
  const file = photo.files[0];
  const reader = new FileReader();

  if (/\.(jpeg|png|gif|jpg)$/i.test(file.name)) {
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

  if (prev.id == "prevArticleImage") {
    previewphoto.classList.add("preview-artical");
  } else if (prev.id == "prevHeroImagePost") {
    previewphoto.classList.add("preview-post");
  }

  switch (post.id) {
    case "postAuthorPhoto": imgData.fileAuthorName = file.name;
    case "postImage": imgData.filePostName = file.name;
  }
}

//отображение загруженного фото
function reclass() {
  const previewphoto = document.getElementById('prevAuthor');
  let photo = document.getElementById('uploadPhoto');

  previewphoto.classList.remove("input-file__button");
  previewphoto.classList.add("input-file__button-new");
  photo.classList.toggle("input-file__button");
  photo.classList.toggle("input-file__button-upload");
}

function rename(i) {
  let uploadNew = document.getElementsByClassName('input-file__upload')[i].textContent = "Upload New";
}

function deleteName(i) {
  let upload = document.getElementsByClassName('input-file__upload')[i].textContent = "Upload";
}

//изменение блока на появление кнопки remove
function changeBlock(label, blockAdd, removeBlock) {
  if (label.id == "labelAuthor") {
    reclass();
    rename(0);
  }
  else {
    label.classList.toggle("input-file__img-block");
    if (label.id == "labelAticle") {
      rename(1);
      let articleImgExtension = document.getElementById('articleImgExtension');
      articleImgExtension.style.display = "none";
    }
    else {
      rename(2);
      let postImgExtension = document.getElementById('postImgExtension');
      postImgExtension.style.display = "none";
    }
  }

  removeBlock.classList.toggle("form-block__remove");
  removeBlock.classList.toggle("form-block__remove_hidden");

  blockAdd.classList.toggle("form-block__block");
}

//удаление превью
function removePost(previewImg, uploadImg, block, removeBlock, prevPost) {
  let removePost = previewImg;
  removePost.src = '';

  if (previewImg.id == "prevArticleImage") {
    removePost.classList.remove("preview-artical");
  } else if (previewImg.id == "prevHeroImagePost") {
    removePost.classList.remove("preview-post");
  }
  changeBlock(uploadImg, block, removeBlock);

  prevPost.src = '';

  if (removeBlock.id == "removeBlockAuthor") {
    deleteName(0);
  } else {
    if (removeBlock.id == "removeBlock") {
      deleteName(1);
    }
    else {
      deleteName(2);
    }
  }
}

function checkAllValue() {
  titleValue = valueEmpty(title, errorTitle)
  subtitleValue = valueEmpty(subtitle, errorSubtitle)
  authorValue = valueEmpty(author, errorAuthorName)
  contentValue = valueEmptyContent(content)


  let articleImgExtension = document.getElementById('articleImgExtension');
  let errorAuthorPhoto = document.getElementById('errorAuthorPhoto');
  let postImgExtension = document.getElementById('postImgExtension');

  authorPhoto = valueImg(photo, errorAuthorPhoto)
  articlePhoto = valueImg(heroImage, articleImgExtension)
  postPhoto = valueImg(heroImagePost, postImgExtension)

  if (titleValue || 
    subtitleValue || 
    authorValue || 
    contentValue || 
    dateValue || 
    !authorPhoto || 
    !articlePhoto || 
    !postPhoto) {
    showError()
    return false
  }
  else {
    title.value = ''
    subtitle.value=''
    author.value=''
    content.value=''
    photo.value=''
    heroImage.value=''
    heroImagePost.value=''
    prevHeroImagePost.src=''
    prevArticleImage.src=''
    prevAuthor.src=''
    showOk()
    return true
  }
}

function valueEmpty(field, label) {
  if (field.value !== null && field.value == '') {
    label.classList.remove('form-block__error')
    label.classList.add('error')
    label.style.color = '#E86961'
    field.style.borderColor = "#E86961";
    field.style.marginBottom = '5px';
    return true
  }
  else {
    label.style.color = '#7A7A7A'
    field.style.borderColor = '#7A7A7A';
    return false
  }
}

function valueEmptyContent(content) {
  if (content.value !== null && content.value == '') {
    content.style.borderColor = '#E86961'
    return true
  }
  else {

    content.style.borderColor = '#7A7A7A'
    return false
  }
}

function valueImg(uploadImage, label) {
  uploadImage = photo.files[0]
  if (uploadImage) {
    label.style.color = '#7A7A7A'
    return true
  }
  else {
    label.style.color = '#E86961'
    label.classList.remove('form-block__error')
    label.classList.add('error')
    label.textContent = 'Image is required.'
    return false
  }
}

function showError() {
  let incorrect = document.getElementById('incorrect')
  let correct = document.getElementById('correct')
  incorrect.classList.remove('main__block__error')
  incorrect.classList.add('error-form')
  correct.classList.remove('correct')
}

function showOk() {
  let incorrect = document.getElementById('incorrect')
  let correct = document.getElementById('correct')
  incorrect.classList.add('main__block__error')
  incorrect.classList.remove('error-form')
  correct.classList.add('correct')
}