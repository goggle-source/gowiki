# `gowiki` - statistical website generator on go

## Description
##### This project reads an md file with html templates and returns a ready-made html page via http

## Quick start

1. clone repository:
   `git clone <url project>`

2. installing dependencies from the project root:
   `go mod tidy`

3. You need to define config.yml in the project root.
**example** config.yml:

```yml
{
    path_for_md_file: "./content",
    path_for_ready_html_file: "./public",
    path_for_template_html: "./template",
    port: 8084
}
```

4. run the create_env command from the project root:
   `go run cmd/cli/main.go build create_env`

5. run the build command from the project root: 
   `go run cmd/cli/main.go build`

6. run the serve commandfrom the project root(The http server will be started on localhost on the port specified in the config):
   `go run cmd/cli/main.go serve` 


## Description Commands

### build

**description**
This command generates ready-made html pages.

**example**
`go run cmd/cli/main.go build`

### serve

**description**
This command starts an http server on the specified port

**example**
`go run cmd/cli/main.go serve` 

### create_env

**description**
This command creates folders for md files, templates, and ready-made html pages at the paths specified in the config.

**example**
`go run cmd/cli/main.go create_env`

### new

**description**
This command copies the specified file to the folder specified in **config.yml**. Copies only files with the .md and .html extension

**example**

example for md file:
`go run cmd/cli/main.go new --path./dir_file/name_file.md`

example for html file:
`go run cmd/cli/main.go new --path./dir_file/name_file.html`

## Requirements

The hidden go version 1.26.4

### Requirements for the design of md files and html template
1. The md file must contain the following mandatory fields in the metadata:
    **name_template** - the name of the template with the extension
    **name_ready_html** - the name of the ready html page

2. If there are other parameters in the template besides Content, they should be included in the md file's metadata. All other data that is not in the metadata will be added to the *.Content*

    example template:

    ```html
    <!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 700px; margin: 40px auto; padding: 0 20px; }
        h1 { color: #222; }
        .meta { color: #777; font-size: 0.9em; margin-bottom: 20px; }
        .content { line-height: 1.6; }
        hr { border: 0; border-top: 1px solid #eee; margin: 30px 0; }
    </style>
</head>
<body>
    <h1>{{.Title}}</h1>
    <div class="meta">
        Автор: {{.Name}} 
        Email: {{.Email}}
        Опубликовано: {{.CreatedAt}}
    </div>
    <div class="content">
        {{.Content}}
    </div>
    <hr>
</body>
</html>

    ```

    example md file for this template: 

    ```md 
    
    ---
    Title: blog for golang
    Name: goggle
    Email: it's not email
    name_template: test1.html
    name_ready_html: test1_ready_html_for_golang.html
    CreatedAt: 2026-07-20
    ---

    this test md file for example

    ```

3. md files should be located in the folder specified in the config as path_for_md_file, and templates should be located in path_for_template_html.

4. the template must contain a field *.Content*
