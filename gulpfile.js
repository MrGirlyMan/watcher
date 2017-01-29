"use strict";

const gulp = require("gulp");
const concat = require("gulp-concat");
const rename = require("gulp-rename");
const uglify = require("gulp-uglify");
const minifyCss = require("gulp-minify-css");

var path = {
	ALL: ["!./src/frontman/client/*.js", "./src/frontman/client/**/*.js",
			"./src/frontman/client/css/*.css", "./src/frontman/client/**/*.css"],
	JS: ["!/src/frontman./client/*.js", "./src/frontman/client/**/*.js"],
	CSS: ["./src/frontman/client/css/*.css", "./src/frontman/client/**/*.css"],
	JS_OUT: "bundle.min.js",
	JS_DEST: "./src/frontman/public/js",
	CSS_OUT: "main.min.css",
	CSS_DEST: "./src/frontman/public/css",
	DEST: "public"
};

gulp.task("bundle", () => {
	return gulp.src(path.JS)
		.pipe(concat("bundle.js"))
		.pipe(rename(path.JS_OUT))
		.pipe(uglify())
		.pipe(gulp.dest(path.JS_DEST));
});

gulp.task("styles", () => {
	return gulp.src(path.CSS)
	.pipe(concat("bundle.css"))
	.pipe(minifyCss())
	.pipe(rename(path.CSS_OUT))
	.pipe(gulp.dest(path.CSS_DEST));
});

gulp.task("build", () => {
	gulp.start("styles");
	gulp.start("bundle");
});

gulp.task("watch", () => {
	gulp.start("build");
	gulp.watch(path.ALL, ["bundle", "styles"])
});

gulp.task("default", ["watch"]);
