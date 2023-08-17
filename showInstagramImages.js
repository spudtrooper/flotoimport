/**
 * Shows the various sized images for an instagram image.
 */
(function () {
  let imgs = [];
  document.querySelectorAll('img[style="object-fit: cover;"]').forEach(function (el) {
    const alt = el.getAttribute('alt');
    console.log(alt);
    let srcset = el.getAttribute('srcset');
    if (!srcset) {
      console.log('no srcset for alt', alt);
      return;
    }
    srcset.split(',').forEach(function (e) {
      let p = e.split(' ');
      let type = p[1];
      if (type == '640w') {
        let img = p[0];
        console.log('IMG', img);
        imgs.push(img);
      }
    });
  });
  console.log();
  console.log();
  console.log('Run the following command to download all the images:');
  let out = imgs.map((img, i) => 'curl -s \'' + img + '\' -o img' + i + '.jpg').join('; ');
  console.log(out);
})();
