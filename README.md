# flotoimport

This will import photos from http://instagram/spudtrooper to http://jeffpalm.com/floto in a round-about way.

[floto](http://jeffpalm.com/floto) started as a J2ME app in 2006 that uploaded photos directly to this server starting with [this photo](https://jeffpalm.com/floto/20060602141910.jpg), then uploaded photos by attaching to an email, then from instagram as of [this photo](https://jeffpalm.com/floto/20131226024955.jpg), now via instagram, but not directly since they don't support that API anymore. Along the way it would also send photos to dropbox and other cloud places, but it doesn't anymore.

Caveat (a pretty big one): This won't do anything for anyone else but me.

## Importing photos

Pre:

```bash
mkdir - ~/Desktop/raw
```

	1. Navigate to http://instagram.com/spudtrooper
	2. Run [showInstagramImages.js](showInstagramImages.js)
	3. Run the `curl` command line at the end of the output from `~/Desktop/raw`
	4. Remove the images you don't want to import
	5. Import all images by running the following from this directory:

	```bash
	go run main.go --dir ~/Desktop/raw
	```
	
	6. Upload.
	   a. cd ~/Desktop/floto
	   a. ftp jeffpalm.com / cd public_html/floto
	   b. Username: jeffpalm
	   c. Password: ********
	   d. prompt
	   e. binary
	   f. mput *
	   g. quit
	   
