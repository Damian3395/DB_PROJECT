var gulp = require('gulp');
var source = require('vinyl-source-stream');
var browserify = require('browserify');
var reactify = require('reactify');

var path = {
    HTML: 'src/index.html',
    MINIFIED_OUT: 'application.min.js',
    OUT: 'application.js',
    DEST: 'build',
    DEST_SRC: 'src',
    ALL: ['./src/**/*.jsx'],
    ENTRY_POINT: './src/App.js'
}

    gulp.task('build', function(){
        var bundle = browserify({entries: path.ENTRY_POINT})
        .transform('reactify')
        .bundle()
        .pipe(source(path.MINIFIED_OUT))
        .pipe(gulp.dest(path.DEST));
    });

    gulp.task('default', ['build']);

