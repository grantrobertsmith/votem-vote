/**
	SETUP AND INITIALIZATION
*/

var gulp = require('gulp');
var util = require('gulp-util');
var es   = require('event-stream');
var path = require('path');

var config = {
	// paths to the assets.
	paths: {
		src:        __dirname + '/assets/',
		build:      __dirname + '/public/assets/',
		publics:    __dirname + '/public/',
		migrations: __dirname + '/db/migrations/',
		templates:  __dirname + '/templates/'
	},

	// The manifest file holds the mappings for the fingerprinted files,
	// example: {"js/bootstrap.js": "js/bootstrap-82e1ab4c0.js"}
	manifestBase: 'public/assets',
	// app/routes.go NotFound handler loads this manifest file
	// to serve compiled assets in production mode
	manifestPath: 'public/assets/manifest.json',

	// if undefined, set to false. set production mode to true by
	// passing in --production, i.e: gulp compile --production
	// If production is set to true then livereload.js will be ignored.
	// production is automatically set to true when running "build" task.
	production: !!util.env.production,

	// buildTarget is initialized by the setup function.
	// buildTarget is set to build or VOTEM_VOTE_SERVER_PUBLIC_PATH if present. 
	// All tasks get their build output path from this variable.
	buildTarget: "",
};

// update the buildTarget if VOTEM_VOTE_SERVER_PUBLIC_PATH env var is set.
// This will be set to a /tmp folder when running "abcweb dev", so that
// your development assets compile to tmp instead of ./public
// which is usually reserved for full builds.
// setup also sets config.production to true if executing the build task.
(function setup() {
	if (process.env.VOTEM_VOTE_SERVER_PUBLIC_PATH) {
		config.buildTarget = path.join(process.env.VOTEM_VOTE_SERVER_PUBLIC_PATH, 'assets')
	} else {
		config.buildTarget = config.paths.build
	}

	// set config.production to true if executing the gulp build task
	if (process.argv.slice(2).indexOf("build") != -1) {
		config.production = true
	}
})();

/**
	COMPILE AND MOVE TASKS
*/

gulp.task('compile-css', function() {
	var less    = require('gulp-less');
	var sass    = require('gulp-sass');
	var postcss = require('gulp-postcss');

	return es.merge(
		gulp.src(path.join(config.paths.src, 'css/**/*.less')).pipe(less()),
		gulp.src([
			path.join(config.paths.src, 'css/**/*.scss'), 
			path.join(config.paths.src, 'css/**/*.sass')
		]).pipe(sass()),
		gulp.src([
			path.join(config.paths.src, 'css/**/*.css'), 
			path.join(config.paths.src, 'vendor/css/**/*.css')
		])
	).pipe(postcss([require('autoprefixer', 'postcss-flexbug-fixes')]))
		.pipe(gulp.dest(path.join(config.buildTarget, 'css')));
});

gulp.task('compile-js', function() {
	var glob = [
		path.join(config.paths.src, 'js/**/*.js'), 
		path.join(config.paths.src, 'vendor/js/**/*.js')
	]
	// ignore specific files when in production mode
	if (config.production) {
		glob.push('!' + path.join(config.paths.src, 'vendor/js/livereload.js'));
	}

	return gulp.src(glob).pipe(gulp.dest(path.join(config.buildTarget, 'js')));
});

gulp.task('compile-img', function() {
	return gulp.src([
		path.join(config.paths.src, 'img/**/*'), 
		path.join(config.paths.src, 'vendor/img/**/*')
	]).pipe(gulp.dest(path.join(config.buildTarget, 'img')));
});

gulp.task('compile-fonts', function() {
	return gulp.src([
		path.join(config.paths.src, 'fonts/**/*'), 
		path.join(config.paths.src, 'vendor/fonts/**/*')
	]).pipe(gulp.dest(path.join(config.buildTarget, 'fonts')));
});

gulp.task('compile-video', function() {
	return gulp.src([
		path.join(config.paths.src, 'video/**/*'), 
		path.join(config.paths.src, 'vendor/video/**/*')
	]).pipe(gulp.dest(path.join(config.buildTarget, 'video')));
});

gulp.task('compile-audio', function() {
	return gulp.src([
		path.join(config.paths.src, 'audio/**/*'),
		path.join(config.paths.src, 'vendor/audio/**/*')
	]).pipe(gulp.dest(path.join(config.buildTarget, 'audio')));
});

/**
	MINIFY TASKS	
*/

gulp.task('minify-css', function() {
	var minifyCSS = require('gulp-clean-css');
	return gulp.src(path.join(config.buildTarget, 'css/**/*.css'))
		.pipe(minifyCSS())
		.pipe(gulp.dest(path.join(config.buildTarget, 'css')));
});

gulp.task('minify-js', function() {
	var minifyJS = require('gulp-minify');

	return gulp.src(path.join(config.buildTarget, 'js/**/*.js'))
		.pipe(minifyJS({noSource: true, ext: {min: '.js'}}))
		.pipe(gulp.dest(path.join(config.buildTarget, 'js')));
});

/**
	GZIP TASK - GZIP COMPRESS ALL ASSETS INTO ACCOMPANYING .GZ FILES
*/

gulp.task('gzip', function() {
	var gzip = require('gulp-gzip');

	return gulp.src(path.join(config.buildTarget, '**/*'))
		.pipe(gzip())
		.pipe(gulp.dest(config.buildTarget));
});

/**
	MANIFEST TASK - FINGERPRINT ASSETS AND GENERATE MANIFEST FILE	
*/

gulp.task('manifest', function() {
	var rev    = require('gulp-rev');
	var revdel = require('gulp-rev-delete-original');

	// Fingerprint all assets, delete old assets and then
	// create the manifest file out of the new asset filenames.
	return gulp.src(path.join(config.buildTarget, '**/*'))
		.pipe(rev())
		.pipe(revdel())
		.pipe(gulp.dest(function(file) {
			return file.base
		}))
		.pipe(rev.manifest(config.manifestPath, {base: config.manifestBase}))
		.pipe(gulp.dest(config.buildTarget));
});

/**
	CLEAN TASK - CLEAN BUILD ASSETS DIRECTORY TO PREVENT LEFT-OVER FILES
*/

gulp.task('clean', function() {
	var del = require('del');

	// Safety catch.
	if (config.buildTarget == "" || config.buildTarget == "/") {
		return
	}

	return del([config.buildTarget], {force: true})
});

/**
	COPY TASK - COPY NON-COMPILED ASSETS INTO VOTEM_VOTE_SERVER_PUBLIC_PATH IF SET 
*/

gulp.task('copy', function(cb) {
	if (process.env.VOTEM_VOTE_SERVER_PUBLIC_PATH) {
		return gulp.src([
			path.join(config.paths.publics, '**/*'), 
			'!' + config.paths.build
		]).pipe(gulp.dest(process.env.VOTEM_VOTE_SERVER_PUBLIC_PATH))
	}

	return cb();
});

/**
	WRAPPER TASKS
*/

// Compile task executes all compile and move tasks.
gulp.task('compile', gulp.parallel('compile-js', 'compile-css', 'compile-fonts', 'compile-img', 'compile-video', 'compile-audio'));

// Minify task executes all minify tasks.
gulp.task('minify', gulp.parallel('minify-js', 'minify-css'));

// Build task executes compile and move tasks,
// then minify tasks,
// then finally the fingerprint and manifest generation task.
gulp.task('build', gulp.series('clean', 'copy', 'compile', 'minify', 'gzip', 'manifest'));

// Default task executes all compile and move tasks.
gulp.task('default', gulp.series('compile'));

/**
	WATCH TASK
*/

// Watch task compiles all assets on start, then watches assets and html templates
// for change, recompiles them and signals livereload to reload the browser.
gulp.task('watch', gulp.series('clean', 'copy', 'compile', function() {
	var watch      = require('gulp-watch');
	var less       = require('gulp-less');
	var sass       = require('gulp-sass');
	var postcss    = require('gulp-postcss');
	var plumber    = require('gulp-plumber');
	var livereload = require('gulp-livereload');

	livereload.listen();
	
	// stylesheets
	es.merge(
		watch([
			path.join(config.paths.src, 'css/**/*.scss'), 
			path.join(config.paths.src, 'css/**/*.sass')
		]).pipe(plumber()).pipe(sass()),
		watch(path.join(config.paths.src, 'css/**/*.less'), { ignoreInitial: true })
			.pipe(plumber()).pipe(less()),
		watch([
			path.join(config.paths.src, 'css/**/*.css'),
			path.join(config.paths.src, 'vendor/css/**/*.css')
		])
	).pipe(postcss([require('autoprefixer', 'postcss-flexbug-fixes')]))
		.pipe(gulp.dest(path.join(config.buildTarget, 'css')))
		.pipe(livereload())

	// javascript
	watch([
		path.join(config.paths.src, 'js/**/*.js'), 
		path.join(config.paths.src, 'vendor/js/**/*.js')
	]).pipe(plumber())
		.pipe(gulp.dest(path.join(config.buildTarget, 'js')))
		.pipe(livereload())

	// images
	watch([
		path.join(config.paths.src, 'img/**/*'),
		path.join(config.paths.src, 'vendor/img/**/*')
	]).pipe(plumber())
		.pipe(gulp.dest(path.join(config.buildTarget, 'img')))
		.pipe(livereload())

	// fonts
	watch([
		path.join(config.paths.src, 'fonts/**/*'), 
		path.join(config.paths.src, 'vendor/fonts/**/*')
	]).pipe(plumber())
		.pipe(gulp.dest(path.join(config.buildTarget, 'fonts')))
		.pipe(livereload())

	// videos
	watch([
		path.join(config.paths.src, 'video/**/*'),
		path.join(config.paths.src, 'vendor/video/**/*')
	]).pipe(plumber())
		.pipe(gulp.dest(path.join(config.buildTarget, 'video')))
		.pipe(livereload())

	// audio
	watch([
		path.join(config.paths.src, 'audio/**/*'),
		path.join(config.paths.src, 'vendor/audio/**/*')
	]).pipe(plumber())
		.pipe(gulp.dest(path.join(config.buildTarget, 'audio')))
		.pipe(livereload())
	
	// non-compiled assets copy in public folder
	if (process.env.VOTEM_VOTE_SERVER_PUBLIC_PATH) {
		watch([
			path.join(config.paths.publics, '**/*'), 
			'!' + config.paths.build
		]).pipe(plumber())
			.pipe(gulp.dest(process.env.VOTEM_VOTE_SERVER_PUBLIC_PATH))
			.pipe(livereload())
	}

	// Watch html templates
	return watch(path.join(config.paths.templates, '**/*'))
		.pipe(livereload())
}));
