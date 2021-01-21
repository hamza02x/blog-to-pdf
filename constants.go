package main

const constSampleINI = `domain = blog.com

# supports multiple sitemap
# comma is important
sitemaps_url = https://blog.com/sitemap.xml,

# in case of wp-json type
# comma is important
# sitemaps_url = https://www.muslimmedia.info/wp-json/wp/v2/posts?per_page=50&post_type=post,https://www.muslimmedia.info/wp-json/wp/v2/posts?per_page=50&post_type=page,

get_sitemap_by_wp_json = false

# Generate pdf or not, if false then only combined-html files will be created!
generate_pdf = false

# Force Re-fetch urls from sitemap / by wget
# comment this if you have used custom sort (sorted.txt maybe?)
force_urls_fetch = true

url_file = ./urls.txt

# asc or desc, according to sitemap time or date
# only works during url grab

# better try this -
# $ cat urls.txt | sort -u | tee -a sorted.txt
# $ cat urls.txt | sort -n | tee -a sorted.txt

# Default name: <min_range>-<max_range>_your_blog.com.pdf
# If you set this then: <min_range>-<max_range>_custom.pdf
pdf_file_name = blog

# This is one of the important value, since we will merge (article_per_pdf) 10 article in a single PDF
# so, which portion of the HTML will be merged in the main Layout? Ex: 'div#content', 'div.post', 'article'
# for "id", use "$" instead of "#"
article_parent_element = article

# for "id", use "$" instead of "#"
article_title_class = h2.entry-title

# for "id", use "$" instead of "#"
elements_to_remove = footer, aside, .respond

article_per_pdf = 25

# use $ instead of ;
browser_user_agent = Mozilla/5.0 (iPhone$ CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1

# There will be in need of some REPLACES, that's why had to use JSON file,
# Make sure that's valid JSON file

append_article_url_in_title = true

append_auto_article_number_in_title = true

# only $1 will be replaced
# this is called first
# then string replaces
pattern_replaces_file = pattern_replaces.json

string_replaces_file = string_replaces.json

# Force Re-fetch htmls from server
force_html_fetch = false

# -1 => work with all url
limit_urls = -1

# for blogpost url, use: github.com/hamza02x/sort-blogspot-urls
# urls should start with 'https'
# https://x.blogspot.com/2014/11/blog-post_2.html
# https://x.blogspot.com/2014/11/blog-post.html
# https://x.blogspot.com/2014/11/blog-post_1.html

# $ sort-url-by-path-date urls.txt

post_order = desc

# Only generate non generated PDFs
skip_pdf_creation_if_exists_already = false

pdf_output_dir_path = ./pdf

# A0        =>	841 x 1189 mm
# A1        =>	594 x 841 mm
# A2        =>	420 x 594 mm
# A3        =>	297 x 420 mm
# A4        =>	210 x 297 mm, 8.26
# A5        =>	148 x 210 mm
# A6        =>	105 x 148 mm
# A7        =>	74 x 105 mm
# A8        =>	52 x 74 mm
# A9        =>	37 x 52 mm
# B0        =>	1000 x 1414 mm
# B1        =>	707 x 1000 mm
# B10       =>	31 x 44 mm
# B2        =>	500 x 707 mm
# B3        =>	353 x 500 mm
# B4        =>	250 x 353 mm
# B5        =>	176 x 250 mm, 6.93
# B6        =>	125 x 176 mm
# B7        =>	88 x 125 mm
# B8        =>	62 x 88 mm
# B9        =>	33 x 62 mm
# C5E       =>	163 x 229 mm
# Comm10E   =>	105 x 241 mm, U.S. Common 10 Envelope
# Custom    =>	Unknown, or a user defined size.
# DLE       =>	110 x 220 mm
# Executive =>	7.5 x 10 inches, 190.5 x 254 mm
# Folio     =>	210 x 330 mm
# Ledger    =>	431.8 x 279.4 mm
# Legal     =>	8.5 x 14 inches, 215.9 x 355.6 mm
# Letter    =>	8.5 x 11 inches, 215.9 x 279.4 mm
# Tabloid   =>	279.4 x 431.8 mm

pdf_size = A7

# "Landscape" or "Portrait"
pdf_orientation = Portrait

# UI
custom_css_file = custom.css

# Margin / White spaces for pdf (mm)
pdf_margin_top = 3
pdf_margin_left = 3
pdf_margin_right = 3
pdf_margin_bottom = 3
`

const constReplacesJSONStr = `
[
	{"serial": 1, "data": {"<script src=\"//stats.wp.com/w.js?60\" type=\"text/javascript\" async=\"\" defer=\"\"></script>": ""}},
	{"serial": 2, "data": {",v=\"//\"": ",v=\"https://\""}},
	{"serial": 3, "data": {"=\"//": "=\"https://"}},
	{"serial": 4, "data": {" dir=\"ltr\"": ""}},
	{"serial": 5, "data": {"<div><span>": "<p>"}},
	{"serial": 6, "data": {"<div ><span >": "<p>"}},
	{"serial": 7, "data": {"</span></div>": "</p>"}}
]
`

const constReplacesJSONPatternStr = `
[
	{"serial": 1, "data": {" (style=\".*?\")": ""}}
]
`

const constHelpStr = `
# Usage

-c string
	(required) run the config file, ex: blog-to-pdf -c config.ini
-d string
	(required, if -i is passed) initialization directory name, ex: blog-to-pdf -i -d any-blog-name
-ec
	print sample config data to console. ex: blog-to-pdf -ec
-gc
	create sample config file. ex: blog-to-pdf -gc
-i	initialize a new directory for new blog, ex: blog-to-pdf -i -d any-blog-name

`

const constCusotmCSS = `

/* Generated By blog-to-pdf
* Modify the css according to your need 
*/ 

.general-article { width: 100% !important; page-break-after: always; }
.the-page-break-class {page-break-after: always;}
.the-credit {position: relative;line-height: 2rem;color:#222222;font-family:'Noto Sans Bengali',sans-serif;font-weight:200;height:400px;margin:0}
.the-credit .full-height{height:100vh}
.the-credit .flex-center{align-items:center;display:flex;justify-content:center}
.the-credit .position-ref{position:relative}
.the-credit .top-right{position:absolute;right:10px;top:18px}
.the-credit .content{text-align:center;top: 25%;position: absolute;width: 95%;}
.the-credit .title{font-size: 17px}
.the-credit a{color:#33af7f;font-size: 15px;font-weight:600;letter-spacing:.1rem;text-decoration:none;}

.text-center { text-align: center !important; }

.general-article *, 
.general-article a, 
.general-article p, 
.general-article div {
	font-family: "Noto Sans Bengali", serif !important;
	color: #000 !important;
	margin-block-start: 0 !important;
	margin-block-end: 0 !important;	
	font-weight: 400;
	font-size: 12px !important;
}

body { background: #FFF !important; }

.general-article p, .general-article { text-align: justify; }
.general-article p { font-size: 12px !important; font-weight: 400; }

.general-article h1, 
.general-article h2, 
.general-article h3, 
.general-article h4 {
	font-weight: bolder !important;
}

.general-article h1 { font-size: 21px !important; }
.general-article h2 { font-size: 20px !important; }
.general-article h3 { font-size: 19px !important; }
.general-article h4 { font-size: 18px !important; }
.general-article h5 { font-size: 17px !important; }
.general-article h6 { font-size: 12px !important; }

.general-article h1 a, .general-article h2 a, .general-article h3 a, .general-article h4 a {
	font-weight: bolder !important;
}

.general-article .article-origin-link {	font-size: 12px; text-align: center; }
.site-url { text-decoration: none; }

img, iframe {
	max-width: 95% !important;
    object-fit: contain;
    margin: 0 auto !important;
	height: auto;
    -webkit-box-shadow: 0 0 3px 0 rgba(0, 0, 0, .12), 0 1px 2px rgba(0, 0, 0, .24);
    -moz-box-shadow: 0 0 3px 0 rgba(0, 0, 0, .12), 0 1px 2px rgba(0, 0, 0, .24);
    box-shadow: 0 0 3px 0 rgba(0, 0, 0, .12), 0 1px 2px rgba(0, 0, 0, .24);
    -webkit-border-radius: 10px;
    -moz-border-radius: 10px;
    border-radius: 10px;
}

.entry-header {
	text-align: center;
}

`

const htmlTemplate = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>blog-to-pdf</title>
</head>
<body>
	<div class="the-tool-container">
	</div>
</body>
</html>
`

const constCreditHTML = `
<article class="the-credit flex-center position-ref full-height the-page-break-class">
	<div class="content">
		<div class="title m-b-md">
			<h2>title_placeholder</h2>
			<hr>
			<h5>Auto Generated PDF</h5>
			<hr>
			<h5>by 'blog-to-pdf' tool<br></h5>
			<h5>
			<hr>
			You can also generate yourself, get the tool: <br>
				<a href="https://github.com/hamza02x/blog-to-pdf">
					https://github.com/hamza02x/blog-to-pdf
				</a>
			</h5>
		</div>
	</div>
</article>
`

// * {
//     page-break-inside: avoid;
//     page-break-after: avoid;
//     page-break-before: avoid;
//   }
