# Floto

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
	   
