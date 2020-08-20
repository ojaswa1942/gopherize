# Gopherize

Some micro-programs & packages to exercise GoLang

## Contents 
In order of time of implementation:

### Quiz Game
A program to run timed quizes via the command line 
- with unit tests
- Inspiration: Go Coding Exercises by [Jon Calhoun](https://courses.calhoun.io/courses/cor_gophercises).

### URL Shortener
A program (package) that forwards paths to other URLs (similar to Bitly).
- package `main` depicts a sample usage of the created package `urlshortener`
- contains unit tests
- Inspiration: Go Coding Exercises by [Jon Calhoun](https://courses.calhoun.io/courses/cor_gophercises).

### Choose Your Own Adventure
A package to create a web application to render "choose your own adventure" books 
- `cmd/web` depicts a sample usage of the created package `adventure`
- with some unit tests
- Inspiration: Go Coding Exercises by [Jon Calhoun](https://courses.calhoun.io/courses/cor_gophercises).

### Link Parser
A package to parse and extract all links off a HTML file 
- `cmd/parse` depicts a sample usage of the created package `linkparser`
- contains unit & integration tests
- Inspiration: Go Coding Exercises by [Jon Calhoun](https://courses.calhoun.io/courses/cor_gophercises).

### Sitemap Builder
A program to generate sitemap for an URL 
- uses the created [`linkparser`](#link-parser) package to generate links
- contains some unit tests
- Inspiration: Go Coding Exercises by [Jon Calhoun](https://courses.calhoun.io/courses/cor_gophercises).
