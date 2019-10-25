


## INSTALL

```
brew cask install wkhtmltopdf
go get gitlab.com/thejini3/blog-to-pdf
cd $GOPATH/src/gitlab.com/thejini3/blog-to-pdf
go install
```
## Check `constants.go` file according to your needs

## USAGE
```
Usage of blog-to-pdf:
  -article-per-pdf int
    	The number of articles per pdf (default 10)

  -domain string
    	(Required) Domain of the site, Ex: alorpothe.wordpress.com

  -force-html-fetch
    	Re-fetch htmls from server if it's not already fetched in local directory

  -force-sitemap-fetch
    	Re-fetch htmls from server if it's not already fetched in local directory

  -https
    	https or not (default true)
```

## EXAMPLE
```
$ blog-to-pdf -domain=alorpothe.wordpress.com -article-per-pdf=7
$ blog-to-pdf -domain=bibijaan.com -sitemap-slug=sitemap-posts.xml -generate-pdf=false -article-parent-div=".inner" -force-html-fetch=true
```