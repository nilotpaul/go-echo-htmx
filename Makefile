start: 
	@npx tailwindcss -i ./web/styles/globals.css -o ./public/bundle/styles.css --minify
	@node ./esbuild.config.js
	@ENVIRONMENT=PRODUCTION go build -tags '!dev' -o bin/build
	@ENVIRONMENT=PRODUCTION ./bin/build

dev: build
	@ENVIRONMENT=DEVELOPMENT ./bin/build

build:
	@npx tailwindcss -i ./web/styles/globals.css -o ./public/bundle/styles.css
	@node ./esbuild.config.js
	@ENVIRONMENT=DEVELOPMENT go build -tags 'dev' -o bin/build

watch:
	@wgo \
    -exit \
    -file=.go \
    -file=.html \
    node ./esbuild.config.js \
    :: npx tailwindcss -i ./web/styles/globals.css -o ./public/bundle/styles.css \
    :: ENVIRONMENT=DEVELOPMENT go build -tags 'dev' -o bin/build main.go \
    :: ENVIRONMENT=DEVELOPMENT ./bin/build \
    :: wgo -debounce 30ms -dir=public npx livereload public