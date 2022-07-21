/**
 * Shows the various sized images for an instagram image.
 *
 * Usage:
 *   1. Click on an instagram thumbnail.
 *   2. Run this.
 */
(function () {
  let imgs = [];
  document.querySelectorAll('img[class="_aagt"]').forEach(function (el) {
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
