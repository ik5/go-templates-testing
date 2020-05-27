package main

const layout = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title> {{ .Title }}</title>
	</head>
	<body>
	  <section class="container">
			{{template "content" .}}
		</section>
    <footer class="footer">
      <div class="content has-text-centered">
        <p>
					Footer
        </p>
      </div>
    </footer>
	</body>
</html>`

const content = `<p>Inside content</p>`
