var gulp = require('gulp');
var sass = require('gulp-sass');
var jade = require('gulp-jade');
var minifyCss = require('gulp-minify-css');
var uglify = require('gulp-uglify');

gulp.task('styles', function() {
    gulp.src('sass/**/*.scss')
        .pipe(sass({
            errLogToConsole: true
        }))
        .pipe(gulp.dest('./compiled_full/css/'));
});

gulp.task('templates', function() {
    gulp.src('./jade/*.jade')
        .pipe(jade())
        .pipe(gulp.dest('./views/'));
});

gulp.task('uglify', function() {
    return gulp.src('compiled_full/js/*.js')
        .pipe(uglify())
        .pipe(gulp.dest('js/'));
});

gulp.task('minify-css', function() {
    return gulp.src('compiled_full/css/index.css')
        .pipe(minifyCss())
        .pipe(gulp.dest('css/'));
});

gulp.task('default',function() {
    gulp.watch('sass/**/*.scss',['styles', 'minify-css']);
    gulp.watch('jade/**/*.jade',['templates']);
    gulp.watch('compiled_full/js/**/*.js',['uglify']);
});
