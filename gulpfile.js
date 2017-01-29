"use strict";

const gulp = require("gulp");
const sass = require("gulp-sass");
const concat = require("gulp-concat");
const rename = require("gulp-rename");
const uglify = require("gulp-uglify");
const minifyCss = require("gulp-minify-css");


// TODO: SCSS, JS, and HTML Linters
// TODO: Generate a CSS map

var path = {
	// HTML
	HTML: "./src/frontman/client/*.html",
	HTML_DEST: "./src/frontman/build/",

	// Javascript
	JS: ["!/src/frontman/client/*.js", "./src/frontman/client/**/*.js"],
	JS_DEST: "./src/frontman/build/js",
	JS_OUT: "bundle.min.js",

	// SCSS 
	SCSS: "./src/frontman/client/scss/*.scss",
	CSS_DEST: "./src/frontman/build/css",
	CSS_OUT: "main.min.css",
};

gulp.task("html", () => {
	return gulp.src(path.HTML)
		.pipe(gulp.dest(path.HTML_DEST));
});

// TODO: If tests are included in this, they will need to be removed.
gulp.task("bundle", () => {
	return gulp.src(path.JS)
		.pipe(concat("bundle.js"))
		.pipe(rename(path.JS_OUT))
		.pipe(uglify())
		.pipe(gulp.dest(path.JS_DEST));
});

gulp.task("sass", () => {
	return gulp.src(path.SCSS)
		.pipe(sass().on("error", sass.logError))
		.pipe(gulp.dest(path.CSS_DEST))
		.pipe(rename(path.CSS_OUT))
		.pipe(minifyCss())
		.pipe(gulp.dest(path.CSS_DEST));
});

gulp.task("build", () => {
	gulp.start("html");
	gulp.start("bundle");
	gulp.start("sass");
});

gulp.task("watch", () => {
	gulp.start("build");
	gulp.watch(path.HTML, ["html"]);
	gulp.watch(path.JS, ["bundle"]);
	gulp.watch(path.SCSS, ["sass"]);
});

gulp.task("default", ["watch"]);
