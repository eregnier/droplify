# Droplify

## Netlify drop like tool, self hosted

This tool is made to be simple to deploy. 
it is a single binary that open a connexion on port 8020
and display app where it is possible top drop static files.

It looks like below. just select the sub domain where your files will live,
then drop files and here you are.

All static content is stored in a single folder and served as subdomain root. 
This should let webapps and html document run properly.

Folder where static content live is `./uploads` by default. it can be changed like
`UPLOAD_FOLDER=/tmp/my-static-contents ./droplify`

The binary app contains the webapp and will 
unpack and serve it when the app is running
from the upload folder.