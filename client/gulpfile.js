var gulp = require('gulp');
var sass = require('gulp-sass');
var jade = require('gulp-jade');

gulp.task('styles', function() {
    gulp.src('sass/**/*.scss')
        .pipe(sass({
            errLogToConsole: true
        }))
        .pipe(gulp.dest('./css/'));
});

gulp.task('templates', function() {
    gulp.src('./jade/*.jade')
        .pipe(jade())
        .pipe(gulp.dest('./views/'));
});

gulp.task('default',function() {
    gulp.watch('sass/**/*.scss',['styles']);
    gulp.watch('jade/**/*.jade',['templates']);
});
