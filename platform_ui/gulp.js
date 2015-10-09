var gulp = require('gulp');
var uglify = require('gulp-uglify');
var source = require('vinyl-source-stream');
var browserify = require('browserify');
var watchify = require('watchify');
var reactify = require('reactify');
var streamify = require('gulp-streamify');

var path = {
    HTML: 'src/index.html',
    MINIFIED_OUT: 'build.min.js',
    OUT: 'application.js',
    DEST: 'build',
    DEST_BUILD: 'build',
    DEST_SRC: 'src',
    ENTRY_POINT: './'
}

    gulp.task('copy', function(){
        gulp.src(path.HTML)
            .pipe(gulp.dest(path.DEST));
    });

    gulp.task('watch', function(){
        gulp.watch(path.HTML, ['copy']);

        var watcher = watchify(browserify({
            entries: [path.ENTRY_POINT],
            transform: [reactify],
            debug: true,
            cache: {}, pacakgeCache: {}, fullPath: true
        }));

        return watcher.on('update', function(){
            watcher.bundle()
                .pipe(source(path.OUT))
                .pipe(gulp.dest(path.DEST_SRC))
                console.log('Updated');
        })
        .bundle()
        .pipe(source(path.OUT))
        .pipe(gulp.dest(path.DEST_SRC));
    });

    gulp.task('build', function(){
        browserify({
            entries: [path.ENTRY_POINT],
            transform: [reactify],
        )}
        .bundle()
        .pipe(source(path.MINIFIED_OUT))
        .pipe(streamify(uglify(path.MINIFED_OUT)))
        .pipe(gulp.dest(path.DEST_BUILD));
    });

    gulp.task('production', ['build']);

    gulp.task('default', ['watch']);
}
